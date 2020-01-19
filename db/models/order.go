package models

import (
	"time"
)
type Payment int
const (
	paid = iota
	pending
	failed
)

func ReturnStatus(p Payment) String() string {
	return [...]string{"paid", "pending", "failed"}[p]
}
//Cake struct
type Cake struct{
	name string
	price float64
	quanitiy int
	category string
}
//Order struct
type Order struct {
	CustomerName    string
	CustomerEmail   string
	CustomerPhone   string
	DeliveryAddress string
	Cakes []Cake
	PaymentStatus ReturnStatus()
	Createdat       time.Time
}
