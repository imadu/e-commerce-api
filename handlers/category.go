package handlers

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/imadu/e-commerce-api/db/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var categoryCollection = client.Database("cakes-and-cream-go").Collection("categories")

// CreateNewCategory function
func CreateNewCategory(c echo.Context) error {
	r := new(models.Category)
	if err := c.Bind(r); err != nil {
		log.Errorf("Could not bind request to struct: %+v", err)
		return sendError(c, "", "", "")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := categoryCollection.InsertOne(ctx, r)
	return sendSuccess(c, result.InsertedID)
}

// GetCategories Function
func GetCategories(c echo.Context) error {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var categories []*models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := categoryCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return sendError(c, "", "Categories do not exist", "")
	}

	for result.Next(ctx) {
		var category models.Category
		err := result.Decode(&category)
		if err != nil {
			log.Errorf("something went wrong", err)
		}
		categories = append(categories, &category)
	}

	if err := result.Err(); err != nil {
		log.Fatal(err)
	}
	result.Close(ctx)

	return sendData(c, categories)
}

// GetCategory function
func GetCategory(c echo.Context) error {
	id := c.Param("id")
	var result models.Category
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := categoryCollection.FindOne(ctx, id).Decode(&result)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return sendError(c, "404", "Category do not exist", "failed")
	}
	return sendData(c, result)
}

//UpdateCategory function
func UpdateCategory(c echo.Context) error {
	id := c.Param("id")
	// body is supposed to be of type models.Category to that we can pass it to the update variable
	body := new(models.Category)
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{primitive.E{Key: "name", Value: body.Name}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := categoryCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return sendError(c, "400", "could not update category", "failed")
	}
	return sendSuccess(c, result.MatchedCount)
}

// DeleteCategory function
func DeleteCategory(c echo.Context) error {
	id := c.Param("id")
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := categoryCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return sendError(c, "500", "could not delete", "failed")
	}
	return sendSuccess(c, result.DeletedCount)
}
