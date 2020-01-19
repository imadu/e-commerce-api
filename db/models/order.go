package models

import (
	"time"
)

type Order struct {
	customer_name    string
	customer_email   string
	customer_phone   string
	delivery_address string
	created_at       time.Time
}
