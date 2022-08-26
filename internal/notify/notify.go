package notify

import (
	"context"
	"errors"
	"sync"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PollNotifications(wg *sync.WaitGroup) error {

	sugar := logger.GetInstance()
	config := config.GetInstance()
	ticker := time.NewTicker(time.Duration(config.NotificationPollingTime) * time.Second)

	sugar.Debugln("PollNotifications called")

	// Initial run because for loop only starts after ticker duration
	err := getDbNotifications()
	if err != nil {
		sugar.Error(err)
		return err
	}

	for range ticker.C {
		getDbNotifications()
	}

	wg.Done()

	return nil
}

func getDbNotifications() error {

	sugar := logger.GetInstance()
	collAppts := db.Datastore.Database.Collection("appointments")
	collNotifs := db.Datastore.Database.Collection("notifications")

	// Get all appointments
	cursorAppts, err := collAppts.Find(context.TODO(), bson.M{})
	if err != nil {
		sugar.Error(err)
		return err
	}

	for cursorAppts.Next(context.TODO()) {

		var appt models.DbAppointment

		if err = cursorAppts.Decode(&appt); err != nil {
			sugar.Error(err)
			return err
		}

		filter := bson.M{"locationIds": appt.LocationId}

		// Get notifications that match LocationId
		cursorNotifs, err := collNotifs.Find(context.TODO(), filter)
		if err != nil {
			sugar.Error(err)
			return err
		}

		for cursorNotifs.Next(context.TODO()) {

			var notif models.Notification

			if err = cursorNotifs.Decode(&notif); err != nil {
				sugar.Error(err)
				return err
			}

			sugar.Infof("LocationId: %v", appt.LocationId)
			sugar.Infof("User: %v", notif.UserId)
			sugar.Infof("Token: %v", notif.Token)

			// If returned appointment date is before target date
			// And the last notification is not equal to LastUpdated
			// if (wsApptDate.Before(notification.TargetDate)) &&
			// 	(dbAppt.LastUpdated != notification.LastNotifiedDate) {
			// }

			result, err := sendNotification(notif.Token)
			if err != nil {
				sugar.Error(err)
				return err
			}

			if result == "Success" {

			}
		}

		cursorNotifs.Close(context.TODO())

	}

	cursorAppts.Close(context.TODO())

	// 	notifyResult, err := notify.SendNotification(notification.Token)
	// 	if err != nil {
	// 		sugar.Error(err)
	// 	}

	// 	sugar.Infof("Notification sent: %v", notifyResult)

	// 	filter := bson.M{
	// 		"token": notification.Token,
	// 	}
	// Update notification record
	// 	update := bson.M{
	// 		"$set": bson.M{
	// 			"DbAppointment.Notification.lastNotifiedDate": updateTime,
	// 		},
	// 	}

	// 	updateResult, err := coll.UpdateOne(context.TODO(), filter, update)
	// 	if err != nil {
	// 		sugar.Error(err)
	// 	}

	// 	sugar.Infof("%v records updated", updateResult.ModifiedCount)

	// }

	// sugar.Infof("Processed %v notifications", j)
	// }

	return nil

}

func sendNotification(token string) (string, error) {

	sugar := logger.GetInstance()
	sugar.Infoln("sendNotification called")

	return "Success", nil
}

func UpdateNotification(notification models.Notification) error {

	sugar := logger.GetInstance()
	sugar.Debugln("UpdateNotification called")

	coll := db.Datastore.Database.Collection("notifications")
	opts := options.Update().SetUpsert(true)

	filter := bson.M{
		"userId": notification.UserId,
	}

	update := bson.M{
		"$set": bson.M{
			"userId":      notification.UserId,
			"locationIds": notification.LocationIds,
			"token":       notification.Token,
		},
	}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		sugar.Error(err)
	} else if result.ModifiedCount > 0 {
		sugar.Infof("Notifications updated for user %v, %v records updated", notification.UserId, result.ModifiedCount)
	} else if result.UpsertedCount > 0 {
		sugar.Infof("Notifications added for user %v, _id: %v", notification.UserId, result.UpsertedID)
	} else {
		sugar.Infoln("No documents added/modified")
		return errors.New("no documents added/modified")
	}
	return nil
}
