package appts

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/notify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PopulateAppointmentsDb(locationList *[]models.Location) error {

	coll := db.Datastore.Database.Collection("appointments")

	for _, location := range *locationList {

		filter := bson.M{
			"locationId": location.LocationId,
		}

		count, err := coll.CountDocuments(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 {

			var appointmentEntry models.Appointment

			appointmentEntry.LocationId = location.LocationId
			appointmentEntry.LastUpdated = time.Now()

			result, err := coll.InsertOne(context.TODO(), appointmentEntry)
			if err != nil {
				log.Println(err)
				return err
			} else {
				log.Printf("DB Result: %v\n", result)
			}
		}

	}

	return nil

}

func PollAppointments(wg *sync.WaitGroup) error {
	coll := db.Datastore.Database.Collection("appointments")

	ticker := time.NewTicker(5 * time.Second)

	for range ticker.C {

		cursor, err := coll.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Println(err)
			return err
		}

		defer cursor.Close(context.TODO())

		for cursor.Next(context.TODO()) {

			var dbAppt models.DbAppointment

			if err = cursor.Decode(&dbAppt); err != nil {
				log.Println(err)
				return err
			}

			wsAppt := GetWsAppointment(dbAppt.LocationId)

			if wsAppt.StartTimestamp != "" {
				updateTime := time.Now()

				wsApptDate, err := dateStringToDate(wsAppt.StartTimestamp)
				if err != nil {
					log.Fatal(err)
				}

				filter := bson.M{
					"locationId": wsAppt.LocationId,
				}
				// opts := options.Replace().SetUpsert(true)

				// Get appointment from DB
				err = coll.FindOne(context.TODO(), filter).Decode(&dbAppt)
				// If document doesn't exist, insert document
				if err == mongo.ErrNoDocuments {

					dbAppt.LocationId = wsAppt.LocationId
					dbAppt.Date = wsApptDate
					dbAppt.LastUpdated = updateTime

					result, err := coll.InsertOne(context.TODO(), dbAppt)
					if err != nil {
						log.Fatal(err)
					} else {
						log.Printf("DB Result: %v\n", result)
					}
				} else if err != nil {
					log.Fatal(err)
					return errors.New("Error getting appointments")
				}

				if wsApptDate != dbAppt.Date {

					update := bson.M{
						"$set": bson.M{
							"date":        wsApptDate,
							"lastUpdated": updateTime,
						},
					}

					result, err := coll.UpdateOne(context.TODO(), filter, update)
					if err != nil {
						log.Fatal(err)
					}

					log.Printf("%v records updated\n", result.ModifiedCount)
				}

				// Go through all notifications
				for j, notification := range dbAppt.NotificationsList {

					// If returned appointment date is before target date
					// And the last notification is not equal to LastUpdated
					if (wsApptDate.Before(notification.TargetDate)) &&
						(dbAppt.LastUpdated != notification.LastNotifiedDate) {
						notifyResult, err := notify.SendNotification(notification.Token)
						if err != nil {
							log.Fatal(err)
						}

						log.Printf("Notification sent: %v\n", notifyResult)

						filter := bson.M{
							"token": notification.Token,
						}
						// Update notification record
						update := bson.M{
							"$set": bson.M{
								"lastNotifiedDate": updateTime,
							},
						}

						updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
						if err != nil {
							log.Fatal(err)
						}

						log.Printf("%v records updated\n", updateResult.ModifiedCount)

					}

					log.Printf("Processed %v notifications\n", j)
				}
				time.Sleep(1 * time.Second)
			}
		}
	}

	wg.Done()

	return nil
}

func GetWsAppointment(locationId int) models.WsAppointment {

	var buf bytes.Buffer
	msg := config.Config.Urls["appts"]
	substitute := models.ApptUrlValue{LocationId: locationId}

	tmpl, err := template.New("msg").Parse(msg)
	if err != nil {
		log.Fatal(err)
	}

	// Insert LocationId into URL
	err = tmpl.Execute(&buf, substitute)
	if err != nil {
		log.Fatal(err)
	}

	url := buf.String()

	// Get appointment from WS
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var appts []models.WsAppointment

	json.Unmarshal(body, &appts)

	if len(appts) > 0 {
		return appts[0]
	}

	return models.WsAppointment{}

}

func dateStringToDate(dateString string) (time.Time, error) {

	date, err := time.Parse("2006-01-02T15:04", dateString)
	if err != nil {
		log.Fatalln(err)
		return time.Time{}, errors.New("Unable to convert string to date")
	}

	return date, nil
}
