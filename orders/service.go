package orders

import (
	"context"
	"os"
	"time"

	"github.com/imadu/e-commerce-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName = os.Getenv("DB_NAME")

var orderCollection = db.Client.Database(dbName).Collection("orders")

var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

//GetOrder returns an order by id
func GetOrder(id string) (Order, error) {
	var result Order
	filter := bson.M{"_id": bson.M{"$eq": id}}
	defer cancel()
	err := orderCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil

}

//GetOrders returns a list of orders
func GetOrders(limit int64, page int64, q Order) ([]*Order, error) {
	var result []*Order
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(page * limit)

	defer cancel()
	cur, err := orderCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var order Order
		err := cur.Decode(&order)
		if err != nil {
			return nil, err
		}

		result = append(result, &order)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//Create makes an order
func Create(o Order) (*mongo.InsertOneResult, error) {
	defer cancel()

	result, err := orderCollection.InsertOne(ctx, o)

	if err != nil {
		return nil, err
	}

	return result, err

}
