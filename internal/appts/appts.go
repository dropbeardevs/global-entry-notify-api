package appts

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
	"text/template"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Refreshes list of appointments based on locations
func PollAppointmentList(wg *sync.WaitGroup) error {

	var err error
	sugar := logger.GetInstance()
	config := config.GetInstance()
	ticker := time.NewTicker(time.Duration(config.LocationsPollingTime) * time.Second)

	sugar.Debugln("PollAppointmentList called")

	locationList, err := locations.GetLocations()
	if err != nil {
		sugar.Error(err)
	}

	// Initial run because for loop only starts after ticker duration
	populateAppointmentsDb(locationList)

	for range ticker.C {
		populateAppointmentsDb(locationList)
	}

	wg.Done()

	return nil
}

func populateAppointmentsDb(locationList *[]models.Location) {

	sugar := logger.GetInstance()
	coll := db.GetInstance().Database.Collection("appointments")

	sugar.Debugln("populateAppointmentsDb called")

	for _, location := range *locationList {

		filter := bson.M{
			"locationId": location.LocationId,
		}

		count, err := coll.CountDocuments(context.TODO(), filter)
		if err != nil {
			sugar.Error(err)
		}

		if count == 0 {

			var appointmentEntry models.DbAppointment

			appointmentEntry.LocationId = location.LocationId
			appointmentEntry.LastUpdated = time.Now()

			result, err := coll.InsertOne(context.TODO(), appointmentEntry)
			if err != nil {
				sugar.Error(err)
			} else {
				sugar.Infof("Appointment LocationId %v inserted. Insert ID: %v", appointmentEntry.LocationId, result.InsertedID)
			}
		}
	}
}

// Get individual appointments from each location
func PollAppointments(wg *sync.WaitGroup) error {

	sugar := logger.GetInstance()
	config := config.GetInstance()

	ticker := time.NewTicker(time.Duration(config.AppointmentPollingTime) * time.Second)

	sugar.Debugln("PollAppointments called")

	for range ticker.C {
		getDbAppointments()
	}

	wg.Done()

	return nil
}

func getDbAppointments() error {

	sugar := logger.GetInstance()
	coll := db.GetInstance().Database.Collection("appointments")

	cursor, err := coll.Find(context.TODO(), bson.M{})
	if err != nil {
		sugar.Error(err)
		return err
	}

	for cursor.Next(context.TODO()) {

		var dbAppt models.DbAppointment
		var wsApptDate time.Time

		if err = cursor.Decode(&dbAppt); err != nil {
			sugar.Error(err)
			return err
		}

		wsAppt := getWsAppointment(dbAppt.LocationId)

		if wsAppt.StartTimestamp == "" {
			// If blank, save as time zero
			wsApptDate = time.Time{}
		} else {
			wsApptDate, err = dateStringToDate(wsAppt.StartTimestamp)
			if err != nil {
				sugar.Error(err)
			}
		}

		updateTime := time.Now()

		// If there are changes to the appointment date, update document
		if wsApptDate != dbAppt.Date {

			filter := bson.M{
				"locationId": dbAppt.LocationId,
			}

			update := bson.M{
				"$set": bson.M{
					"date":        wsApptDate,
					"lastUpdated": updateTime,
				},
			}

			result, err := coll.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				sugar.Error(err)
			}

			sugar.Infof("LocationId: %v - %v records updated", dbAppt.LocationId, result.ModifiedCount)
		}

		time.Sleep(2 * time.Second)
	}

	cursor.Close(context.TODO())

	return nil
}

func getWsAppointment(locationId int) models.WsAppointment {

	sugar := logger.GetInstance()
	config := config.GetInstance()
	var buf bytes.Buffer
	msg := config.Urls["appts"]
	substitute := models.ApptUrlValue{LocationId: locationId}

	sugar.Debugln("getWsAppointment called")

	tmpl, err := template.New("msg").Parse(msg)
	if err != nil {
		sugar.Error(err)
	}

	// Insert LocationId into URL
	err = tmpl.Execute(&buf, substitute)
	if err != nil {
		sugar.Error(err)
	}

	url := buf.String()

	// Get appointment from WS
	resp, err := http.Get(url)
	if err != nil {
		sugar.Error(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Error(err)
	}

	var appts []models.WsAppointment

	json.Unmarshal(body, &appts)

	if len(appts) > 0 {
		return appts[0]
	}

	return models.WsAppointment{}

}

func dateStringToDate(dateString string) (time.Time, error) {

	sugar := logger.GetInstance()

	date, err := time.Parse("2006-01-02T15:04", dateString)
	if err != nil {
		sugar.Error(err)
		return time.Time{}, errors.New("unable to convert string to date")
	}

	return date, nil
}
