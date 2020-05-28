package models

import (
	"context"
	"log"
	"time"

	"mt-scale/common"
	"mt-scale/syslog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client *mongo.Client
	// ctx        context.Context
	checkCount int
)

func init() {
	syslog.Debug("Init mongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := common.GetConfStr("database.url")
	clientOptions := options.Client().ApplyURI(url)
	// clientOptions.TLSConfig = &tls.Config{}
	// tlsCfg := clientOptions.TLSConfig
	// tlsCfg.InsecureSkipVerify = true

	var err error
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	syslog.Debug("Connected to MongoDB!")
}

// MongoDbCheck Check mongo database connect.
func MongoDbCheck() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
			syslog.Debug("Ping context cancel")
		}()
		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			syslog.Error(err, "MongoDbCheck")
		}
	}
}

// Collection Get mongo collection
func Collection(name string) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	dbName := common.GetConfStr("database.name")
	db := client.Database(dbName)

	collection := db.Collection(name)
	return collection, ctx, cancel
}
