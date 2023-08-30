package netxdcustomerdalmodels

import "time"

//"time"

//"go.starlark.net/lib/time"

type Customer struct {
	CustomerId int32     `json:"customer_id" bson:"customer_id"`
	FirstName  string    `json:"first_name" bson:"first_name"`
	LastName   string    `json:"last_name" bson:"last_name"`
	BankId     string    `json:"bank_id" bson:"bank_id"`
	Balance    float64   `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" bson:"updated_at"`
	IsActive   bool      `json:"active" bson:"active"`
}
type CustomerResponse struct {
	CustomerId int32  `json:"customer_id" bson:"customer_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}
