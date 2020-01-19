package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() (*mongo.Client, *context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://imadu:default@12345@ds063789.mlab.com:63789/cakes-and-cream-go"))
	if err != nil {
		log.Fatalf("could not connect to db")
	}

	return client, &ctx
}
