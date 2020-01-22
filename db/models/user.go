package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User struct
type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname   string             `json:"firstName" bson:"firstName"`
	Lastname    string             `json:"lastName" bson:"lastName"`
	Phonenumber string             `json:"phoneNumber" bson:"phoneNumber"`
	Password    string
	Username    string    `json:"username" bson:"username"`
	Email       string    `json:"email" bson:"email"`
	Createdat   time.Time `json:"createdAt" bson:"createdAt"`
}
