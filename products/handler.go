package products

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/imadu/cakes-and-cream/db"
	"github.com/imadu/cakes-and-cream/util"
)

var productCollection = db.Client.Database("cakes-and-cream-go").Collection("products")

func getName(name string) Product {
	var result Product
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := productCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result

}

// GetProduct returns a single order
func GetProduct(c echo.Context) error {
	var result Product
	id := c.Param("id")
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := productCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return util.SendError(c, "404", "User does not exist", "failed")
	}

	return util.SendData(c, result)

}

//Create creates a product
func Create(c echo.Context) error {
	name := c.FormValue("name")

	nameExists := getName(name)

	if nameExists.Name == name {
		return util.SendError(c, "400", "duplicate names cannot exist", "failed")
	}

	p := new(Product)
	if err := c.Bind(p); err != nil {
		log.Fatalf("Could not bind request to struct: %+v", err)
		return util.SendError(c, "500", "something went wrong", "failed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := productCollection.InsertOne(ctx, p)

	return util.SendSuccess(c, result.InsertedID)

}
