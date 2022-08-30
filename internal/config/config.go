package config

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"gopkg.in/yaml.v3"
)

var config *models.Configuration
var once sync.Once

func GetInstance() *models.Configuration {
	once.Do(
		func() {

			log.Println("Config initializing")

			var c models.Configuration

			log.Println("GLOBAL_ENTRY_NOTIFY_API_CONFIG: ", os.Getenv("GLOBAL_ENTRY_NOTIFY_API_CONFIG"))

			configFile := os.Getenv("GLOBAL_ENTRY_NOTIFY_API_CONFIG")

			bytes, err := ioutil.ReadFile(configFile)
			if err != nil {
				log.Fatal(err)
			}

			err = yaml.Unmarshal(bytes, &c)
			if err != nil {
				log.Fatal(err)
			}

			config = &c

			log.Println("Config initialized")
		},
	)
	return config
}
