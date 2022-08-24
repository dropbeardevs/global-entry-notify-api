package config

import (
	"io/ioutil"
	"log"
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"gopkg.in/yaml.v3"
)

var config *models.Configuration
var once sync.Once

func LoadConfig(filename *string) {
	once.Do(
		func() {
			var c models.Configuration

			bytes, err := ioutil.ReadFile(*filename)
			if err != nil {
				log.Fatal(err)
			}

			err = yaml.Unmarshal(bytes, &c)
			if err != nil {
				log.Fatal(err)
			}

			config = &c
		},
	)
}

func GetInstance() *models.Configuration {
	return config
}
