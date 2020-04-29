package users

import (
	"context"
	"os"
	"time"

	"github.com/imadu/e-commerce-api/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var dbName = os.Getenv("DB_NAME")

var userCollection = db.Client.Database(dbName).Collection("users")

var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//CreateUser creates a user
func CreateUser(user User) (*mongo.InsertOneResult, error) {
	mod := mongo.IndexModel{
		Keys: bson.M{
			"username": -1,
		}, Options: options.Index().SetUnique(true),
	}

	_, err := userCollection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		return nil, err
	}

	password, _ := hashPassword(user.Password)

	user.Password = password

	defer cancel()
	result, _ := userCollection.InsertOne(ctx, user)

	return result, nil

}

//GetUser returns a user by id
func GetUser(id string) (User, error) {
	var result User
	filter := bson.M{"_id": bson.M{"$eq": id}}

	defer cancel()

	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//GetUsername returns a user search by username
func GetUsername(username string) (User, error) {
	var result User
	filter := bson.M{"username": bson.M{"$eq": username}}

	defer cancel()
	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

//GetUsers returns the list of users
func GetUsers(limit int64, page int64) ([]*User, error) {
	if limit == 0 {
		limit = 10.0
	}

	if page == 0 {
		page = 1.0
	}
	var results []*User
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(page * limit)

	defer cancel()
	cur, err := userCollection.Find(ctx, bson.M{}, findOptions)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		results = append(results, &user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil

}

//UpdateUser updates a user by id
func UpdateUser(id string, user User) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{
		"first_name":   user.Firstname,
		"last_name":    user.Lastname,
		"Phone_number": user.Phonenumber,
		"email":        user.Email,
	}}

	defer cancel()
	result, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

//DeleteUser deletes a user by id
func DeleteUser(id string) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}
