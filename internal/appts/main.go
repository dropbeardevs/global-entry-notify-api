package appts

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"text/template"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/locations"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/utils"
)

func PollAppointments(wg *sync.WaitGroup) {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {

		for _, location := range *locations.LocationsList {
			appointmentsList := GetAppointments(location.Id)

			//appointment := GetAppointments(5010)

			// Empty appointment check
			if len(appointmentsList) > 0 {

				for i, appointment := range appointmentsList {
					// Go through appointment list
					apptDateString := appointment.StartDate

					apptDate, err := time.Parse("2006-01-02T15:04", apptDateString)
					if err != nil {
						log.Fatalln(err)
					}

					for j, notifications := range appointment.NotificationsList {
						if (apptDate.Before(user.BeforeDate)) &&
							(utils.Contains(user.LocationIds, location.Id)) {
							log.Printf("Yay! %v", user.UserId)
						}

						log.Printf("Proessed %v notifications", j)
					}

					log.Printf("Proessed %v appointments", i)
				}

			}

			time.Sleep(1 * time.Second)
		}

	}

	wg.Done()
}

func GetAppointments(locationId int) []Appointment {

	var buf bytes.Buffer
	msg := config.Config.Urls["appts"]
	substitute := ApptUrlValues{LocationId: locationId}

	tmpl, err := template.New("msg").Parse(msg)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(&buf, substitute)
	if err != nil {
		log.Fatal(err)
	}

	url := buf.String()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var appts []Appointment

	json.Unmarshal(body, &appts)

	return appts

}
