package models

import (
	"time"
)

type (
	Order struct {
		ID          uint          `gorm:"primary_key" json:"id"`
		SellerId    uint          `json:"seller_id"`
		CustomerId  uint          `json:"customer_id"`
		Invoice     string        `json:"invoice"`
		TotalPrice  uint          `json:"total_price"`
		OrderStatus string        `json:"order_status"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
		OrderDetail []OrderDetail `json:"-"`
		Customer    Customer      `json:"-"`
		Seller      Seller        `json:"-"`
	}
)
