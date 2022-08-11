package db

import (
	"context"
	"log"

	fb "bitbucket.org/dropbeardevs/global-entry-notify-api/internal/firebase"
	"cloud.google.com/go/firestore"
)

// TODO: Implement interface
type FirestoreClient struct {
	Client     *firestore.Client
	Collection *firestore.CollectionRef
}

var DbClient *FirestoreClient

func InitDataStore() {
	firestoreClient := FirestoreClient{}
	var err error

	ctx := context.Background()

	firestoreClient.Client, err = fb.FirebaseApp.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	firestoreClient.Collection = firestoreClient.Client.Collection("global-entry")

	DbClient = &firestoreClient
}
