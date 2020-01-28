package categories

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Category struct
type Category struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}
