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
	"go.mongodb.org/mongo-driver/mongo"
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

		filter := bson.M{"notificationDetails.locationId": appt.LocationId}

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

			for _, notifDetail := range notif.NotificationDetails {

				// Find notification detail matching location Id
				if notifDetail.LocationId == appt.LocationId {

					// If returned appointment date is before target date
					// And appointment dates are different
					if (appt.Date.Before(notifDetail.TargetDate)) &&
						(appt.Date != notifDetail.AppointmentDate) {

						result, err := sendNotification(notif.Token)
						if err != nil {
							sugar.Error(err)
							return err
						}

						if result == "Success" {

							filter := bson.M{"notificationDetails.locationId": appt.LocationId}

							update := bson.M{
								"$set": bson.M{
									"notificationDetails.$.appointmentDate":  appt.Date,
									"notificationDetails.$.lastNotifiedDate": time.Now(),
								},
							}

							result, err := collNotifs.UpdateOne(context.TODO(), filter, update)
							if err != nil {
								sugar.Error(err)
							} else if result.ModifiedCount > 0 {
								sugar.Infof("NotificationDetail updated for user %v", notif.UserId)
							} else if result.UpsertedCount > 0 {
								sugar.Infof("NotificationDetail added for user %v, _id: %v", notif.UserId, result.UpsertedID)
							} else {
								sugar.Infoln("No documents added/modified")
								return errors.New("no documents added/modified")
							}
						}
					}
				}
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

func UpdateNotificationToken(userId string, token string) error {

	sugar := logger.GetInstance()
	sugar.Debugln("UpdateNotificationToken called")

	coll := db.Datastore.Database.Collection("notifications")
	opts := options.Update().SetUpsert(true)

	filter := bson.M{
		"userId": userId,
	}

	update := bson.M{
		"$set": bson.M{
			"token":       token,
			"lastUpdated": time.Now(),
		},
	}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		sugar.Error(err)
	} else if result.ModifiedCount > 0 {
		sugar.Infof("Token updated for user %v, token: %v", userId, token)
	} else if result.UpsertedCount > 0 {
		sugar.Infof("Token %v added for user %v, _id: %v", token, userId, result.UpsertedID)
	} else {
		sugar.Infoln("No documents added/modified")
		return errors.New("no documents added/modified")
	}

	return nil

}

func UpdateNotificationDetails(userId string, notificationDetails models.NotificationDetails) error {

	sugar := logger.GetInstance()
	sugar.Debugln("UpdateNotificationDetails called")

	coll := db.Datastore.Database.Collection("notifications")
	opts := options.Update().SetUpsert(true)
	var filter bson.M

	filter = bson.M{
		"userId":                         userId,
		"notificationDetails.locationId": notificationDetails.LocationId,
	}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {

			filter = bson.M{
				"userId": userId,
			}

			update := bson.M{
				"$push": bson.M{
					"notificationDetails": notificationDetails,
				},
				"$set": bson.M{
					"lastUpdated": time.Now(),
				},
			}

			updateResult, err := coll.UpdateOne(context.TODO(), filter, update, opts)
			if err != nil {
				sugar.Error(err)
				return err
			} else if updateResult.ModifiedCount > 0 {
				sugar.Infof("Notification Details updated for user %v", userId)
			} else if updateResult.UpsertedCount > 0 {
				sugar.Infof("Notification Details added for user %v, _id: %v", userId, updateResult.UpsertedID)
			} else {
				sugar.Infoln("No documents added/modified")
				return errors.New("no documents added/modified")
			}
		} else {
			return err
		}
	} else {

		update := bson.M{
			"$set": bson.M{
				"notificationDetails.$": notificationDetails,
				"lastUpdated":           time.Now(),
			},
		}

		updateResult, err := coll.UpdateOne(context.TODO(), filter, update, opts)
		if err != nil {
			sugar.Error(err)
			return err
		} else if updateResult.ModifiedCount > 0 {
			sugar.Infof("Notification Details updated for user %v", userId)
		} else if updateResult.UpsertedCount > 0 {
			sugar.Infof("Notification Details added for user %v, _id: %v", userId, updateResult.UpsertedID)
		} else {
			sugar.Infoln("No documents added/modified")
			return errors.New("no documents added/modified")
		}
	}

	return nil

}

func DeleteNotificationDetails(userId string, locationId int) error {

	sugar := logger.GetInstance()
	sugar.Debugln("DeleteNotificationDetails called")

	coll := db.Datastore.Database.Collection("notifications")
	opts := options.Update().SetUpsert(false)

	filter := bson.M{
		"userId":                         userId,
		"notificationDetails.locationId": locationId,
	}

	update := bson.M{
		"$pull": bson.M{
			"notificationDetails": bson.M{
				"locationId": locationId,
			},
		},
		"$set": bson.M{
			"lastUpdated": time.Now(),
		},
	}

	updateResult, err := coll.UpdateMany(context.TODO(), filter, update, opts)
	if err != nil {
		sugar.Error(err)
		return err
	} else if updateResult.ModifiedCount > 0 {
		sugar.Infof("Notification Details updated for user %v", userId)
	} else if updateResult.UpsertedCount > 0 {
		sugar.Infof("Notification Details added for user %v, _id: %v", userId, updateResult.UpsertedID)
	} else {
		sugar.Infoln("No documents added/modified")
		return errors.New("no documents added/modified")
	}

	return nil

}
