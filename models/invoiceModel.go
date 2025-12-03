package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID               primitive.ObjectID `bson:"id"`
	Order_id         string             `json:"order_id"`
	Invoice_id       string             `json:"invoice_id"`
	Payment_method   *string            `json:"payment_method" validate:"eq=CASH|eq=CARD|eq="`
	Payment_status   *string            `json:"payment_status" validate:"required,eq=PANDING|eq=PAID|eq=CANCELLED"`
	Payment_due_date time.Time          `json:"payment_due_date"`
	Created_at       time.Time          `json:"created_at"`
	Updated_at       time.Time          `json:"updated_at"`
}
