package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Ticker struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Symbol string             `bson:"symbol"`
}

type Rate struct {
	Symbol string    `json:"symbol" bson:"symbol"`
	Time   time.Time `bson:"time"`
	Price  string    `json:"price" bson:"price"`
}
