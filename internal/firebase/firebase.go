package fb

import (
	"context"
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

			// Use a service account
			ctx := context.Background()

			app, err := firebase.NewApp(ctx, nil)
			if err != nil {
				sugar.Fatalf("error initializing app: %v\n", err)
			}

			firebaseApp = app
		},
	)

	return firebaseApp
}
