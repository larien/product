package entity

import (
	"time"
)

// Product represents the Product object and contains the fields that can be manipulated by the service
type Product struct {
	ID           string     `json:"id" db:"id"`
	PriceInCents int64      `json:"price_in_cents" db:"price_in_cents"`
	Title        string     `json:"title" db:"title"`
	Description  string     `json:"description" db:"description"`
	CreatedAt    time.Time  `json:"-" db:"created_at"`
	UpdatedAt    time.Time  `json:"-" db:"updated_at"`
	DeletedAt    *time.Time `json:"-" db:"deleted_at"`
	Discount     Discount   `json:"discount"`
}
