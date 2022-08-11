package auth

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFirebaseCreds() []byte {
	json, err := os.Open("/Users/homerdulu/Developer/secrets/global-entry-notify-api/global-entry-c8373-fe72e1ae9c11.json")
	if err != nil {
		log.Fatalln(err)
	}

	byteVal, err := ioutil.ReadAll(json)
	if err != nil {
		log.Fatalln(err)
	}

	return byteVal
}
