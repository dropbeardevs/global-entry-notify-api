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

	var err error

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := homeDir + "/Developer/secrets/global-entry-notify-api/config.yml"
	credsPath := homeDir + "/Developer/secrets/global-entry-notify-api/global-entry-c8373-fe72e1ae9c11.json"

	var wg sync.WaitGroup

	config.LoadConfig(&configPath)

	fb.InitFirebaseApp(&credsPath)

	db.InitDatastore()

	//defer db.DbClient.Client.Close()

	users.GetUsers()

	locations, err := locations.GetLocations()
	if err != nil {
		log.Fatal(err)
	}

	// locations.SetLocations()
	//wg.Add(1)
	appts.PollAppointments(&wg, locations)
	//wg.Wait()
}
