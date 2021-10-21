package models

import (
	"time"
)

type (
	Seller struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Order     []Order   `json:"-"`
		Product   []Product `json:"-"`
	}
)
