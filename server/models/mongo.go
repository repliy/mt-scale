package models

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"mt-scale/common"
	"mt-scale/syslog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	ctx        context.Context
	checkCount int
)

func init() {
	syslog.Debug("Init mongoDB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := common.GetConfStr("database.url")
	clientOptions := options.Client().ApplyURI(url)
	clientOptions.TLSConfig = &tls.Config{}
	tlsCfg := clientOptions.TLSConfig
	tlsCfg.InsecureSkipVerify = true

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
		if err := client.Ping(context.TODO(), nil); err != nil {
			syslog.Error("Database error.", err)
		}
	} else {
		syslog.Error("Mongo database client is nil .")
	}
}

func database() *mongo.Database {
	dbName := common.GetConfStr("database.name")
	db := client.Database(dbName)
	return db
}

// Collection Get mongo collection
func Collection(name string) (*mongo.Collection, context.Context) {
	db := database()
	collection := db.Collection(name)
	return collection, ctx
}
