package initapp

import (
	"log"
	"os"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
)

func InitApp() {
	var err error

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configPath := homeDir + "/Developer/secrets/global-entry-notify-api/config.yml"

	config.LoadConfig(&configPath)

}
