package models

import (
	"html"
	"strings"
	"time"

	"go-mall/utils/token"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheckCustomer(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Customer{}

	err = db.Model(Customer{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID, u.Email, u.Username)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (c *Customer) SaveUserCustomer(db *gorm.DB) (*Customer, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Customer{}, errPassword
	}
	c.Password = string(hashedPassword)
	//remove spaces in username
	c.Username = html.EscapeString(strings.TrimSpace(c.Username))

	var err error = db.Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}
