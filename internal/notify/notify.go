package notify

import (
	"context"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func SendNotification(token string) (string, error) {

	result := "Success"

	return result, nil
}

func AddNotification(notificationDetails models.NotificationDetails) error {

	sugar := logger.GetInstance()
	coll := db.Datastore.Database.Collection("appointments")
	var dbAppt models.DbAppointment

	sugar.Debugln("AddNotification called")

	for _, locationId := range notificationDetails.LocationIds {

		notification := models.Notification{
			Token:            notificationDetails.Token,
			TargetDate:       notificationDetails.TargetDate,
			LastNotifiedDate: time.Time{},
		}

		filter := bson.M{
			"locationId": locationId,
		}

		err := coll.FindOne(context.TODO(), filter).Decode(&dbAppt)
		if err != nil {
			sugar.Error(err)
		}

		dbAppt.NotificationList = append(dbAppt.NotificationList, notification)

		update := bson.M{
			"$set": bson.M{
				"notificationList": dbAppt.NotificationList,
			},
		}

		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			sugar.Error(err)
		} else {
			sugar.Infof("Notifications updated on LocationId %v, %v records updated", locationId, result.ModifiedCount)
		}
	}

	return nil
}
