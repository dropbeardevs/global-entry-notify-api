package main

import (
	"runtime"
	"sync"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/appts"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/initapp"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
)

func init() {
	initapp.InitApp()

	sugar := logger.GetInstance()

	// Get current running filename
	_, filename, _, _ := runtime.Caller(0)

	sugar.Infof("Running file: %v", filename)

}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go locations.PollLocations(&wg)

	time.Sleep(10 * time.Second)

	go appts.PollAppointmentList(&wg)

	time.Sleep(10 * time.Second)

	go appts.PollAppointments(&wg)
	wg.Wait()
}
