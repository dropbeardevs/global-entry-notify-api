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
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/notify"
	"go.mongodb.org/mongo-driver/bson"
)

func PopulateAppointmentsDb(locationList *[]models.Location) error {

	sugar := logger.GetInstance()
	coll := db.Datastore.Database.Collection("appointments")

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
				return err
			} else {
				sugar.Infof("DB Result: %v\n", result)
			}
		}

	}

	return nil

}

func PollAppointments(wg *sync.WaitGroup) error {

	sugar := logger.GetInstance()
	coll := db.Datastore.Database.Collection("appointments")

	ticker := time.NewTicker(60 * time.Second)

	for range ticker.C {

		cursor, err := coll.Find(context.TODO(), bson.M{})
		if err != nil {
			sugar.Error(err)
			return err
		}

		for cursor.Next(context.TODO()) {

			var dbAppt models.DbAppointment

			if err = cursor.Decode(&dbAppt); err != nil {
				sugar.Error(err)
				return err
			}

			wsAppt := GetWsAppointment(dbAppt.LocationId)

			if wsAppt.StartTimestamp != "" {
				updateTime := time.Now()

				wsApptDate, err := dateStringToDate(wsAppt.StartTimestamp)
				if err != nil {
					sugar.Error(err)
				}

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

					sugar.Infof("LocationId: %v - %v records updated\n", wsAppt.LocationId, result.ModifiedCount)
				}

				// Go through all notifications
				for j, notification := range dbAppt.NotificationList {

					// If returned appointment date is before target date
					// And the last notification is not equal to LastUpdated
					if (wsApptDate.Before(notification.TargetDate)) &&
						(dbAppt.LastUpdated != notification.LastNotifiedDate) {

						notifyResult, err := notify.SendNotification(notification.Token)
						if err != nil {
							sugar.Error(err)
						}

						sugar.Infof("Notification sent: %v\n", notifyResult)

						filter := bson.M{
							"token": notification.Token,
						}
						// Update notification record
						update := bson.M{
							"$set": bson.M{
								"DbAppointment.Notification.lastNotifiedDate": updateTime,
							},
						}

						updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
						if err != nil {
							sugar.Error(err)
						}

						sugar.Infof("%v records updated\n", updateResult.ModifiedCount)

					}

					sugar.Infof("Processed %v notifications\n", j)
				}
				time.Sleep(1 * time.Second)
			}
		}

		cursor.Close(context.TODO())
	}

	wg.Done()

	return nil
}

func GetWsAppointment(locationId int) models.WsAppointment {

	sugar := logger.GetInstance()
	config := config.GetInstance()
	var buf bytes.Buffer
	msg := config.Urls["appts"]
	substitute := models.ApptUrlValue{LocationId: locationId}

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
