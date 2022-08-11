package main

import (
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/appts"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/db"
	fb "bitbucket.org/dropbeardevs/global-entry-notify-api/internal/firebase"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/users"
)

func main() {

	configPath := "/Users/homerdulu/Developer/secrets/global-entry-notify-api/config.yml"
	credsPath := "/Users/homerdulu/Developer/secrets/global-entry-notify-api/global-entry-c8373-fe72e1ae9c11.json"

	var wg sync.WaitGroup

	config.LoadConfig(&configPath)

	fb.InitFirebaseApp(&credsPath)

	db.InitDatastore()

	//defer db.DbClient.Client.Close()

	users.InitUsers()

	locations.InitLocations()

	// locations.SetLocations()
	wg.Add(1)
	go appts.PollAppointments(&wg)
	wg.Wait()
}
