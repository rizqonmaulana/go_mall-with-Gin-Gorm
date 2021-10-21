package models

import (
	"time"
)

type (
	ProductRating struct {
		ID         uint      `gorm:"primary_key" json:"id"`
		ProductId  uint      `json:"product_id"`
		CustomerId uint      `json:"customer_id"`
		Rating     uint      `json:"rating"`
		Comment    string    `json:"comment"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		Product    Product   `json:"-"`
		Customer   Customer  `json:"-"`
	}
)
