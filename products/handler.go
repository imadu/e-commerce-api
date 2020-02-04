package products

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/imadu/e-commerce-api/db"
	"github.com/imadu/e-commerce-api/util"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Query struct
type Query struct {
	Name     string
	Category string
	Limit    int64
	Page     int64
}

var dbName = os.Getenv("DB_NAME")

var productCollection = db.Client.Database(dbName).Collection("products")

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

//GetProducts returns a list of products based on the query parameters
func GetProducts(c echo.Context) error {
	var q Query
	q.Limit, _ = strconv.ParseInt(c.QueryParam("limit"), 10, 64)
	q.Name = c.QueryParam("name")
	q.Page, _ = strconv.ParseInt(c.QueryParam("page"), 10, 64)
	q.Category = c.QueryParam("category")

	findOptions := options.Find()
	findOptions.SetLimit(q.Limit)
	findOptions.SetSkip(q.Page * q.Limit)

	var results []*Product
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := productCollection.Find(ctx, bson.D{{q}}, findOptions)

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
