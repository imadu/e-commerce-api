package users

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/imadu/e-commerce-api/db"
	"github.com/imadu/e-commerce-api/util"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var dbName = os.Getenv("DB_NAME")

var userCollection = db.Client.Database(dbName).Collection("users")

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUser(username string) User {
	var result User
	filter := bson.D{primitive.E{Key: "username", Value: username}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

//GetUser returns the user by id
func GetUser(c echo.Context) error {
	var result User
	id := c.Param("id")
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return util.SendError(c, "404", "User does not exist", "failed")
	}

	return util.SendData(c, result)
}

// GetUsers returns all the users
func GetUsers(c echo.Context) error {
	p := c.QueryParam("page")
	l := c.QueryParam("limit")

	page, _ := strconv.ParseInt(p, 10, 64)

	limit, _ := strconv.ParseInt(l, 10, 64)

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(page * limit)

	var results []*User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := userCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		c.Response().WriteHeader(http.StatusNotFound)
		return util.SendError(c, "500", "could not find users", "failed")
	}

	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			log.Errorf("something went wrong", err)
		}
		results = append(results, &user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(ctx)
	return util.SendData(c, results)

}

// CreateUser creates a new user
func CreateUser(c echo.Context) error {
	password := c.FormValue("password")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	mod := mongo.IndexModel{
		Keys: bson.M{
			"username": -1,
		}, Options: options.Index().SetUnique(true),
	}

	ind, err := userCollection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Fatalf("Indexes().CreateOne() ERROR:", err)
	} else {
		// API call returns string of the index name
		fmt.Println("CreateOne() index:", ind)
		fmt.Println("CreateOne() type:", reflect.TypeOf(ind))
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Fatalf("could not hash password", err)
	}

	u := new(User)
	u.Password = hashedPassword
	if err := c.Bind(u); err != nil {
		log.Errorf("Could not bind request to struct: %+v", err)
		defer cancel()
		return util.SendError(c, "500", "something went wrong", "failed")
	}
	defer cancel()
	result, _ := userCollection.InsertOne(ctx, u)
	return util.SendSuccess(c, result.InsertedID)

}

// UpdateUser updates the user
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	body := new(User)
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{primitive.E{Key: "firstName", Value: body.Firstname}, primitive.E{Key: "lastName", Value: body.Lastname}, primitive.E{Key: "email", Value: body.Email}, primitive.E{Key: "phoneNumber", Value: body.Phonenumber}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		return util.SendError(c, "400", "could not update category", "failed")
	}

	return util.SendSuccess(c, result.MatchedCount)
}

// DeleteUser deletes a user
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		return util.SendError(c, "500", "could not delete %+v", "failed")
	}

	return util.SendSuccess(c, result.DeletedCount)
}
