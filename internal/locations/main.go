package locations

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAndSaveWsLocations() (*[]models.Location, error) {

	sugar := logger.GetInstance()

	var err error
	config := config.GetInstance()

	resp, err := http.Get(config.Urls["locations"])
	if err != nil {
		sugar.Error(err)
		return nil, errors.New("error getting location")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Error(err)
		return nil, errors.New("error getting location")
	}

	var locations []models.Location

	err = json.Unmarshal(body, &locations)
	if err != nil {
		sugar.Error(err)
		return nil, errors.New("error unmarshalling location json")
	}

	if locations != nil {
		SaveLocationsDb(&locations)

		return &locations, nil
	} else {
		return nil, errors.New("no locations returned")
	}

}

func SaveLocationsDb(wsLocations *[]models.Location) {

	sugar := logger.GetInstance()

	coll := db.Datastore.Database.Collection("locations")

	// ctx := context.Background()

	for _, wsLocation := range *wsLocations {

		var dbLocation models.Location
		updateTime := time.Now()

		filter := bson.M{"locationId": wsLocation.LocationId}
		//opts := options.Replace().SetUpsert(true)

		err := coll.FindOne(context.TODO(), filter).Decode(&dbLocation)

		// If document doesn't exist, insert document
		if err == mongo.ErrNoDocuments {

			wsLocation.LastUpdated = updateTime

			result, err := coll.InsertOne(context.TODO(), wsLocation)
			if err != nil {
				sugar.Error(err)
			} else {
				sugar.Infof("Insert ID: %v\n", result.InsertedID)
			}
		} else if err != nil {
			sugar.Error(err)
		} else { // Continue on if record exists

			// Convert times
			dbUpdatedDate, err := time.Parse("2006-01-02T15:04:05", dbLocation.WsLocationLastUpdated)
			if err != nil {
				sugar.Error(err)
			}

			wsUpdatedDate, err := time.Parse("2006-01-02T15:04:05", wsLocation.WsLocationLastUpdated)
			if err != nil {
				sugar.Error(err)
			}

			if wsUpdatedDate.After(dbUpdatedDate) {

				wsLocation.LastUpdated = updateTime

				result, err := coll.ReplaceOne(context.TODO(), filter, wsLocation)
				if err != nil {
					sugar.Error(err)
				}

				sugar.Infof("%v documents modified\n", result.ModifiedCount)
			}
		}

	}

	// for _, location := range locations {

	// 	var dbLocation Locations // This is location from database
	// 	var err error

	// 	err := coll.FindOne(context.TODO(), bson.M{"orderNo": orderNo}).Decode(&order)

	// 	doc, err := db.DbClient.Collection.Doc(strconv.Itoa(location.Id)).Get(ctx)
	// 	if err != nil {
	// 		// Check if doc exists, otherwise, insert item
	// 		if status.Code(err) == codes.NotFound {

	// 			result, err := coll.InsertOne(context.TODO(), location)
	// 			if err != nil {
	// 				log.Fatalln(err)
	// 			}
	// 		}
	// 	}

	// 	// If doc exists, unmarshal into Locations struct
	// 	err = doc.DataTo(&dbLocation)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	} else if dbLocation.LastUpdatedDate != location.LastUpdatedDate { // Compare LastUpdatedDate between response and database
	// 		_, err = db.DbClient.Collection.Doc(strconv.Itoa(location.Id)).Set(ctx, location)
	// 		if err != nil {
	// 			log.Fatalln(err)
	// 		}
	// 	}
	// }
}
