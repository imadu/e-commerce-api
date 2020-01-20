package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/imadu/cakes-and-cream/db/models"
)

var client *mongo.Client
var collection = client.Database("cakes-and-cream-go").Collection("categories")

// CreateNewCategory function
func CreateNewCategory(c echo.Context) error {
	r := new(models.Category)
	if err := c.Bind(r); err != nil {
		log.Errorf("Could not bind request to struct: %+v", err)
		return sendError(c, "", "", "")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, r)
	return sendSuccess(c, result.InsertedID)
}

// GetCategories Function
func GetCategories(c echo.Context) error {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var categories []*models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return sendError(c, "", "Categories do not exist", "")
	}

	for result.Next(context.TODO()) {
		var category models.Category
		err := result.Decode(&category)
		if err != nil {
			log.Errorf("something went wrong", err)
		}
		categories = append(categories, &category)
	}

	return sendData(c, categories)
}

// GetCategory function
func GetCategory(c echo.Context) error {
	id := c.Param("id")
	var result models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, id).Decode(&result)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return sendError(c, "", "Category do not exist", "")
	}
	return sendData(c, result)
}
