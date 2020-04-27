package orders

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
	CustomerName    string             `json:"customer_name" bson:"customer_name"`
	CustomerEmail   string             `json:"customer_email" bson:"customer_email"`
	CustomerPhone   string             `json:"customer_phone" bson:"customer_phone"`
	DeliveryAddress string             `json:"address" bson:"address"`
	Cakes           []Cake             `json:"cakes" bson:"cakes"`
	PaymentStatus   Payment            `json:"payment_status" bson:"payment_status"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
}
