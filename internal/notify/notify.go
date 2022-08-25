package notify

import (
	"context"
	"errors"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PollNotifications() {}

func getDbNotifications() {

	sugar := logger.GetInstance()
	collAppts := db.Datastore.Database.Collection("appointments")
	collNotifs := db.Datastore.Database.Collection("notifications")

	// Get all appointments
	cursorAppts, err := collAppts.Find(context.TODO(), bson.M{})
	if err != nil {
		sugar.Error(err)
	}

	for cursor.Next(context.TODO()) {

		var appt models.DbAppointment
		var notification models.Notification

		if err = cursorAppt.Decode(&notification); err != nil {
			sugar.Error(err)
			continue
		}

		// Get all
		cursorNotifs, err := collNotifs.Find(context.TODO(), bson.M{})
		if err != nil {
			sugar.Error(err)
		}

	}

	cursor.Close(context.TODO())

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

			sugar.Infof("Notification sent: %v", notifyResult)

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

			sugar.Infof("%v records updated", updateResult.ModifiedCount)

		}

		sugar.Infof("Processed %v notifications", j)
	}

}

func sendNotification(token string) (string, error) {

	result := "Success"

	return result, nil
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
