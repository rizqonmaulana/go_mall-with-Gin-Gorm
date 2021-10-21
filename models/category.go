package models

import (
	"time"
)

type (
	Category struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Category  string    `json:"category"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Product   []Product `json:"-"`
	}
)
