package models

import (
	"time"
)

type (
	Product struct {
		ID            uint            `gorm:"primary_key" json:"id"`
		Name          string          `json:"name"`
		Description   string          `json:"description"`
		Price         uint            `json:"price"`
		Stock         uint            `json:"stock"`
		SellerId      uint            `json:"seller_id"`
		CategoryId    uint            `json:"category_id"`
		CreatedAt     time.Time       `json:"created_at"`
		UpdatedAt     time.Time       `json:"updated_at"`
		ProductRating []ProductRating `json:"-"`
		Seller        Seller          `json:"-"`
		Category      Category        `json:"-"`
	}
)
