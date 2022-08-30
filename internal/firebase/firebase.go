package fb

import (
	"context"
	"os"
	"sync"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	firebase "firebase.google.com/go"
)

var firebaseApp *firebase.App
var once sync.Once

func GetInstance() *firebase.App {

	once.Do(
		func() {

			sugar := logger.GetInstance()
			sugar.Infoln("Firebase App initializing")

			// Use a service account
			ctx := context.Background()

			sugar.Infoln("GOOGLE_APPLICATION_CREDENTIALS: ", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

			app, err := firebase.NewApp(ctx, nil)
			if err != nil {
				sugar.Fatalf("error initializing app: %v\n", err)
			}

			firebaseApp = app

			sugar.Infoln("Firebase App initialized")
		},
	)

	return firebaseApp
}
