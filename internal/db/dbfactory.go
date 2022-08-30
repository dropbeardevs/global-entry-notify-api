package db

import (
	"context"
	"os"
	"sync"
	"time"

	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/config"
	"bitbucket.org/dropbeardevs/global-entry-notify-api/internal/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoDatastore *MongoDatastore // Reuse this variable for connections
var once sync.Once

type MongoDatastore struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func GetInstance() *MongoDatastore {

	once.Do(
		func() {
			sugar := logger.GetInstance()
			sugar.Infoln("DB initializing")

			mongoDatastore = &MongoDatastore{}

			config := config.GetInstance()

			serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
			clientOptions := options.Client().
				ApplyURI(config.ConnectionString).
				SetServerAPIOptions(serverAPIOptions)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			var err error
			mongoDatastore.Client, err = mongo.Connect(ctx, clientOptions)
			if err != nil {
				sugar.Error(err)
				os.Exit(1)
			}

			mongoDatastore.Database = mongoDatastore.Client.Database("globalEntryNotifyDb")

			// Check the connection
			err = mongoDatastore.Client.Ping(ctx, readpref.Primary())

			if err != nil {
				sugar.Error(err)
				os.Exit(1)
			}

			sugar.Infoln("Connected to MongoDB!")

			databases, err := mongoDatastore.Client.ListDatabaseNames(context.TODO(), bson.M{})
			if err != nil {
				sugar.Error(err)
			}
			sugar.Infof("Databases: %v", databases)
		},
	)

	return mongoDatastore
}
