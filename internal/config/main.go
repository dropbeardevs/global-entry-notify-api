package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var Config *Configuration

func LoadConfig(filename string) {

	var c Configuration

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}

	Config = &c

}
