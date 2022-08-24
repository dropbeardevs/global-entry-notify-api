package main

import (
	"log"
	"os"
	"sync"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/appts"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	fb "bitbucket.org/dropbeardevs/global-entry-notify-api/internal/firebase"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
)

func main() {
	var err error

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := homeDir + "/Developer/secrets/global-entry-notify-api/config.yml"
	credsPath := homeDir + "/Developer/secrets/global-entry-notify-api/global-entry-c8373-fe72e1ae9c11.json"

	config.LoadConfig(&configPath)

	var wg sync.WaitGroup

	fb.InitFirebaseApp(&credsPath)

	db.InitDatastore()

	//defer db.DbClient.Client.Close()

	// locationList, err := locations.GetAndSaveWsLocations()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// appts.PopulateAppointmentsDb(locationList)

	// locations.SetLocations()
	wg.Add(3)
	go locations.PollLocations(&wg)

	time.Sleep(10 * time.Second)

	go appts.PollAppointmentList(&wg)

	time.Sleep(10 * time.Second)

	go appts.PollAppointments(&wg)
	wg.Wait()
}
