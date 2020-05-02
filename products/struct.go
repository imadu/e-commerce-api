package products

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Attribute struct
type Attribute struct {
	Size        int
	Flavor      string
	Description string
}

//Product struct
type Product struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Price     float64            `json:"price" bson:"price"`
	Attribute []Attribute        `json:"attribute" bson:"attribute"`
	Category  string             `json:"category" bson:"category"`
}
