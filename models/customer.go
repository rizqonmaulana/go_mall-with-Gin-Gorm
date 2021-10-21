package models

import (
	"time"
)

type (
	Customer struct {
		ID            uint            `gorm:"primary_key" json:"id"`
		Username      string          `json:"username"`
		Email         string          `json:"email"`
		Password      string          `json:"password"`
		CreatedAt     time.Time       `json:"created_at"`
		UpdatedAt     time.Time       `json:"updated_at"`
		ProductRating []ProductRating `json:"-"`
		Order         []Order         `json:"-"`
	}
)
