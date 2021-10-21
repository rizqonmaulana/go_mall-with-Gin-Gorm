package models

import (
	"time"
)

type (
	OrderDetail struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		OrderId   uint      `json:"order_id"`
		ProductId uint      `json:"product_id"`
		Qty       uint      `json:"qty"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Product   Product   `json:"-"`
		Order     Order     `json:"-"`
	}
)
