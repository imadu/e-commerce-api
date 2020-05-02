package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname   string             `json:"first_name" bson:"first_name"`
	Lastname    string             `json:"last_name" bson:"last_name"`
	Phonenumber string             `json:"phone_number" bson:"phone_number"`
	Password    string
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"email" bson:"email"`
	Createdat   time.Time `json:"created_at" bson:"created_at"`
}
