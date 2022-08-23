package locations

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetLocations() (*[]Locations, error) {

	var err error

	resp, err := http.Get(config.Config.Urls["locations"])
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error getting location")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error getting location")
	}

	var locations []Locations

	err = json.Unmarshal(body, &locations)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error unmarshalling location json")
	}

	if locations != nil {
		SaveLocationsDb(&locations)

		return &locations, nil
	} else {
		return nil, errors.New("no locations returned")
	}

}

func SaveLocationsDb(locations *[]Locations) {

	coll := db.Datastore.Database.Collection("locations")

	// ctx := context.Background()

	for _, location := range *locations {

		var l Locations
		updateTime := time.Now()

		filter := bson.M{"locationId": location.LocationId}
		//opts := options.Replace().SetUpsert(true)

		err := coll.FindOne(context.TODO(), filter).Decode(&l)
		// If document doesn't exist, insert document
		if err == mongo.ErrNoDocuments {

			location.LastUpdated = updateTime

			result, err := coll.InsertOne(context.TODO(), location)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Insert ID: %v\n", result.InsertedID)
			}
		} else if err != nil {
			log.Fatal(err)
		}

		// Convert times
		dbDate, _ := time.Parse("2006-01-02T15:04:05", l.LocationLastUpdatedDate)
		wsDate, _ := time.Parse("2006-01-02T15:04:05", location.LocationLastUpdatedDate)

		if (wsDate.After(dbDate)) && (l.LocationLastUpdatedDate != "") {
			location.LastUpdated = updateTime
			result, err := coll.ReplaceOne(context.TODO(), filter, location)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%v documents modified\n", result.ModifiedCount)
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
