package locations

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PollLocations(wg *sync.WaitGroup) error {

	sugar := logger.GetInstance()
	config := config.GetInstance()
	ticker := time.NewTicker(time.Duration(config.LocationsPollingTime) * time.Second)

	sugar.Debugln("PollLocations called")

	// Initial run because for loop only starts after ticker duration
	locationList := getWsLocations()
	saveLocationsDb(locationList)

	for range ticker.C {
		locationList := getWsLocations()
		saveLocationsDb(locationList)
	}

	wg.Done()

	return nil
}

func getWsLocations() *[]models.Location {

	var err error
	sugar := logger.GetInstance()
	config := config.GetInstance()

	sugar.Debugln("getWsLocations called")

	resp, err := http.Get(config.Urls["locations"])
	if err != nil {
		sugar.Error(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		sugar.Error(err)
		return nil
	}

	var locations []models.Location

	err = json.Unmarshal(body, &locations)
	if err != nil {
		sugar.Error(err)
		return nil
	}

	if locations != nil {
		sugar.Infof("%v records found", len(locations))
		return &locations
	} else {
		return nil
	}
}

func saveLocationsDb(wsLocations *[]models.Location) {

	sugar := logger.GetInstance()

	coll := db.GetInstance().Database.Collection("locations")

	sugar.Debugln("saveLocationsDb called")

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
				sugar.Infof("LocationId %v added. Insert ID: %v",
					wsLocation.LocationId, result.InsertedID)
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

				sugar.Infof("Location LocationId %v modified - %v documents modified",
					wsLocation.LocationId, result.ModifiedCount)

			}
		}
	}
}

func GetLocations() (*[]models.Location, error) {
	sugar := logger.GetInstance()
	coll := db.GetInstance().Database.Collection("locations")

	var locationList []models.Location

	sugar.Debugln("GetLocations called")

	filter := bson.M{
		"operational": true,
		"inviteOnly":  false,
		"services": bson.M{
			"name": "Global Entry",
		},
	}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		sugar.Error(err)
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {

		var dbLocation models.Location

		if err = cursor.Decode(&dbLocation); err != nil {
			sugar.Error(err)
			return nil, err
		}

		locationList = append(locationList, dbLocation)
	}

	return &locationList, nil
}
