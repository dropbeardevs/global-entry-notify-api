package fb

import (
	"context"
	"log"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/auth"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebaseApp(credsPath *string) {

	// Use a service account
	ctx := context.Background()

	json := auth.GetFirebaseCreds(credsPath)

	sa := option.WithCredentialsJSON(json)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	FirebaseApp = app

}
