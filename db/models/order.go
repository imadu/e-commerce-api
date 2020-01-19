package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Payment struct
type Payment int

/* Paid const
Denied const
Failed const
*/
const (
	Paid Payment = iota
	Denied
	Failed
)

func (p Payment) String() string {
	return [...]string{"Paid", "Denied", "Failed"}[p]
}

//Cake struct
type Cake struct {
	name     string
	price    float64
	quanitiy int
	category string
}

//Order struct
type Order struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	CustomerName    string             `json:"customerName" bson:"customerName"`
	CustomerEmail   string             `json:"customerEmail" bson:"customerEmail"`
	CustomerPhone   string             `json:"customerPhone" bson:"customerPhone"`
	DeliveryAddress string             `json:"address" bson:"address"`
	Cakes           []Cake             `json:"cakes" bson:"cakes"`
	Createdat       time.Time          `json:"created_at" bson:"created_at"`
}
