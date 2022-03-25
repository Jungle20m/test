package mobilecard

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MobileCardModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	CustomerID    string             `bson:"customer_id"`
	PhoneNumber   string             `bson:"phone_number"`
	Name          string             `bson:"name"`
	Brand         string             `bson:"brand"`
	Price         int                `bson:"price"`
	Point         int                `bson:"point"`
	Quantity      int                `bson:"quantity"`
	ExchangeTime  time.Time          `bson:"exchange_time"`
	ExchangeYear  int                `bson:"exchange_year"`
	ExchangeMonth int                `bson:"exchange_month"`
	ExchangeDay   int                `bson:"exchange_day"`
	CreateTime    time.Time          `bson:"create_time"`
	UpdateTime    time.Time          `bson:"update_time"`
}

type MobileCardGMVModel struct {
	ID  string `bson:"_id"`
	GMV int    `bson:"price"`
}
