package db

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

//Client is the exported client that returns the database
var Client *mongo.Client

//InitDB exposes the database so we can pass it around in the other handlers
func InitDB() *mongo.Client {
	once.Do(func() {
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://imadu:default@12345@ds063789.mlab.com:63789/cakes-and-cream-go"))
		if err != nil {
			log.Fatalf("could not connect to db")
		}

	})

	return Client
}
