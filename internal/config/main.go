package config

import (
	"io/ioutil"
	"log"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/models"
	"gopkg.in/yaml.v3"
)

var Config *models.Configuration

func LoadConfig(filename *string) {

	var c models.Configuration

	bytes, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}

	Config = &c

}
