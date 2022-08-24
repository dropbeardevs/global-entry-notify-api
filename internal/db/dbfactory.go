package db

import (
	"context"
	"os"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Datastore *MongoDatastore // Reuse this variable for connections

type MongoDatastore struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func InitDatastore() {

	sugar := logger.GetInstance()

	datastore := MongoDatastore{}

	config := config.GetInstance()

	Datastore = &datastore

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(config.ConnectionString).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Datastore.Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		sugar.Error(err)
		os.Exit(1)
	}

	datastore.Database = datastore.Client.Database("globalEntryNotifyDb")

	// Check the connection
	err = datastore.Client.Ping(ctx, readpref.Primary())

	if err != nil {
		sugar.Error(err)
		os.Exit(1)
	}

	sugar.Infoln("Connected to MongoDB!")

	databases, err := datastore.Client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		sugar.Error(err)
	}
	sugar.Infof("Databases: %v", databases)
}
