package products

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

var productCollection = db.Client.Database(dbName).Collection("products")

var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

//GetProduct returns a product via id
func GetProduct(id string) (Product, error) {
	var result Product

	filter := bson.M{"_id": bson.M{"$eq": id}}
	defer cancel()
	err := productCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return result, err
	}
	return result, nil

}

//GetProducts returns a array of the products
func GetProducts(limit int64, page int64) ([]*Product, error) {
	var result []*Product
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(page * limit)

	defer cancel()

	cur, err := productCollection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var product Product
		err := cur.Decode(&product)
		if err != nil {
			return nil, err
		}
		result = append(result, &product)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil

}

//CreateProduct creates a product
func CreateProduct(p Product) (*mongo.InsertOneResult, error) {
	mod := mongo.IndexModel{
		Keys: bson.M{
			"name": -1,
		}, Options: options.Index().SetUnique(true),
	}

	_, err := productCollection.Indexes().CreateOne(ctx, mod)

	if err != nil {
		defer cancel()
		return nil, err
	}

	defer cancel()

	result, _ := productCollection.InsertOne(ctx, p)

	return result, nil

}

//UpdateProduct updates a product
func UpdateProduct(id string, p Product) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{
		"name":      p.Name,
		"price":     p.Price,
		"category":  p.Category,
		"attribute": p.Attribute,
	}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := productCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

//DeleteProduct removes a product
func DeleteProduct(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	defer cancel()
	result, err := productCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
