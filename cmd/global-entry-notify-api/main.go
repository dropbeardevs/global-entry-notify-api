package main

import (
	"runtime"
	"sync"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/api/grpcsrv"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/appts"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/notify"
)

func init() {

	// Get current running filename
	_, filename, _, _ := runtime.Caller(0)

	sugar := logger.GetInstance()

	sugar.Infof("Running filename: %v", filename)

}

func main() {

	var wg sync.WaitGroup

	wg.Add(5)

	go grpcsrv.InitGrpcSrv(&wg)

	go locations.PollLocations(&wg)

	time.Sleep(10 * time.Second)

	go appts.PollAppointmentList(&wg)

	time.Sleep(10 * time.Second)

	go appts.PollAppointments(&wg)

	time.Sleep(10 * time.Second)

	go notify.PollNotifications(&wg)

	wg.Wait()
}
