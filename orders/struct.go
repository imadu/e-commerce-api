package orders

import (
	"time"

	"github.com/imadu/e-commerce-api/products"
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

//Order struct
type Order struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	CustomerName    string             `json:"customer_name" bson:"customer_name"`
	CustomerEmail   string             `json:"customer_email" bson:"customer_email"`
	CustomerPhone   string             `json:"customer_phone" bson:"customer_phone"`
	DeliveryAddress string             `json:"address" bson:"address"`
	Product         []products.Product `json:"product" bson:"product"`
	Reference       string             `json:"reference" bson:"reference"`
	PaymentStatus   Payment            `json:"payment_status" bson:"payment_status"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
}
