package handlers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/imadu/e-commerce-api/db/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// CreateNewCategory function
func CreateNewCategory(c echo.Context, client *mongo.Client) {
	r := new(models.Category)
	if err := c.Bind(r); err != nil {
		log.Errorf("Could not bind request to struct: %+v", err)
		return sendError(c, "", "", "")
	}
	collection := client.Database("cakes-and-cream-go").Collection("category")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, r)
	return sendSuccess(c, nil)
}

// GetCategories Function
func GetCategories(c echo.Context, client *mongo.Client) []models.Category {
	var categories []models.Category
	collection := client.Database("cakes-and-cream-go").Collection("category")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.Find(&categories, ctx)
	return sendData(c, categories)

}
