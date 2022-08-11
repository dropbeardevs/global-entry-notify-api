package locations

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var LocationsList *[]Locations

func InitLocations() {

	resp, err := http.Get(config.Config.Urls["locations"])

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var locations []Locations

	json.Unmarshal(body, &locations)

	SaveLocationsDb(locations)

	LocationsList = &locations

}

func SaveLocationsDb(locations []Locations) {

	coll := db.Datastore.Database.Collection("locations")

	// ctx := context.Background()

	for _, location := range locations {

		var l Locations

		filter := bson.M{"_id": location.Id}
		opts := options.Replace().SetUpsert(true)

		err := coll.FindOne(context.TODO(), filter).Decode(&l)
		if err != nil {
			log.Fatal(err)
		}

		// Convert times
		dbDate, _ := time.Parse("2006-01-02T15:04:05", l.LastUpdatedDate)
		wsDate, _ := time.Parse("2006-01-02T15:04:05", location.LastUpdatedDate)

		if wsDate.After(dbDate) {
			result, err := coll.ReplaceOne(context.TODO(), filter, location, opts)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Inserted document with _id: %v\n", result.UpsertedID)
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
