package main

import (
	"log"
	"os"
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/appts"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	fb "bitbucket.org/dropbeardevs/global-entry-notify-api/internal/firebase"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/users"
)

func main() {

	var wg sync.WaitGroup

	rootdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	config.LoadConfig(rootdir + "/config.yml")

	fb.InitFirebaseApp()

	db.InitDataStore()

	defer db.DbClient.Client.Close()

	users.InitUsers()

	locations.InitLocations()

	// locations.SetLocations()
	wg.Add(1)
	go appts.PollAppointments(&wg)
	wg.Wait()
}
