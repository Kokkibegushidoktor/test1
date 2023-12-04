package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticker struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Symbol string             `bson:"symbol"`
}

type Rate struct {
	SymbolID  primitive.ObjectID  `bson:"symbolID"`
	Timestamp primitive.Timestamp `bson:"timestamp"`
	Price     float32             `bson:"price"`
}
