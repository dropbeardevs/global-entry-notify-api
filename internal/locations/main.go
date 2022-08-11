package locations

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	LocationsList = &locations

}

func SaveLocationsInDb(locations []Locations) {

	ctx := context.Background()

	for _, location := range locations {

		var dbLocation Locations // This is location from database
		var err error

		doc, err := db.DbClient.Collection.Doc(strconv.Itoa(location.Id)).Get(ctx)
		if err != nil {
			// Check if doc exists, otherwise, insert item
			if status.Code(err) == codes.NotFound {

				_, err = db.DbClient.Collection.Doc(strconv.Itoa(location.Id)).Set(ctx, location)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}

		// If doc exists, unmarshal into Locations struct
		err = doc.DataTo(&dbLocation)
		if err != nil {
			log.Fatalln(err)
		} else if dbLocation.LastUpdatedDate != location.LastUpdatedDate { // Compare LastUpdatedDate between response and database
			_, err = db.DbClient.Collection.Doc(strconv.Itoa(location.Id)).Set(ctx, location)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}
