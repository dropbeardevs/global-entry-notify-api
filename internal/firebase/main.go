package fb

import (
	"context"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/auth"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebaseApp(credsPath *string) {

	sugar := logger.GetInstance()

	// Use a service account
	ctx := context.Background()

	json := auth.GetFirebaseCreds(credsPath)

	sa := option.WithCredentialsJSON(json)

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		sugar.Error(err)
	}

	FirebaseApp = app

}
