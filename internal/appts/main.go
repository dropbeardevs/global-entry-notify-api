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
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/users"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/utils"
)

func PollAppointments(wg *sync.WaitGroup) {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {

		for _, location := range *locations.LocationsList {
			appointment := GetAppointments(location.Id)

			//appointment := GetAppointments(5010)

			// Empty appointment check
			if len(appointment) > 0 {
				// Go through user list
				apptDateString := appointment[0].StartDate

				apptDate, err := time.Parse("2006-01-02T15:04", apptDateString)
				if err != nil {
					log.Fatalln(err)
				}

				for _, user := range *users.UserList {
					if (apptDate.Before(user.BeforeDate)) &&
						(utils.Contains(user.LocationIds, location.Id)) {
						log.Printf("Yay! %v", user.UserId)
					}
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
