package auth

import (
	"io/ioutil"
	"log"
	"os"
)

func GetFirebaseCreds(credsPath *string) []byte {
	json, err := os.Open(*credsPath)
	if err != nil {
		log.Fatalln(err)
	}

	byteVal, err := ioutil.ReadAll(json)
	if err != nil {
		log.Fatalln(err)
	}

	return byteVal
}
