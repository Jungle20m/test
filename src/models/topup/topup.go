package topup

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TopupModel struct {
	ID                   primitive.ObjectID `bson:"_id"`
	CustomerID           string             `bson:"customer_id"`
	PhoneNumber          string             `bson:"phone_number"`
	RecipientPhoneNumber string             `bson:"recipient_phone_number"`
	Amount               int                `bson:"amount"`
	Brand                string             `bson:"brand"`
	PaymentMethod        string             `bson:"payment_method"`
	Status               string             `bson:"status"`
	Description          string             `bson:"description"`
	TopupTime            time.Time          `bson:"topup_time"`
	TopupYear            int                `bson:"topup_year"`
	TopupMonth           int                `bson:"topup_month"`
	TopupDay             int                `bson:"topup_day"`
	CreateTime           time.Time          `bson:"create_time"`
	UpdateTime           time.Time          `bson:"update_time"`
}

type TopupGmvModel struct {
	ID  string `bson:"_id"`
	GMV int    `bson:"amount"`
}
