package models

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"mt-scale/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	ctx        context.Context
	checkCount int
)

func init() {
	fmt.Println(">>>>>> :init MongoDB")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	url := utils.GetConfStr("database.url")

	clientOptions := options.Client().ApplyURI(url)
	clientOptions.TLSConfig = &tls.Config{}
	tlsCfg := clientOptions.TLSConfig
	tlsCfg.InsecureSkipVerify = true

	var err error
	client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

// MongoDbCheck Check mongo database connect.
func MongoDbCheck() {
	checkCount++
	fmt.Printf(">>>>>> :Check mongo database %d times.\n", checkCount)

	if client != nil {
		if err := client.Ping(context.TODO(), nil); err != nil {
			fmt.Println(">>>>>> :Database error ", err)
			log.Fatal(err)
		} else {
			fmt.Println(">>>>>> :Database is ok.")
		}
	} else {
		fmt.Println(">>>>>> :Mongo database client is nil .")
	}
}

func database() *mongo.Database {
	dbName := utils.GetConfStr("database.name")
	db := client.Database(dbName)
	return db
}

// Collection Get mongo collection
func Collection(name string) (*mongo.Collection, context.Context) {
	db := database()
	collection := db.Collection(name)
	return collection, ctx
}
