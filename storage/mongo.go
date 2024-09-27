package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func InitMongo(uri string) error {
	clientOptions := options.Client().ApplyURI(uri)

	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")
	return nil
}
