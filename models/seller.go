package models

import (
	"go-mall/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Seller struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Order     []Order   `json:"-"`
		Product   []Product `json:"-"`
	}
)

func LoginCheckSeller(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Seller{}

	err = db.Model(Seller{}).Where("username = ?", username).Take(&u).Error

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

func (s *Seller) SaveUserSeller(db *gorm.DB) (*Seller, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Seller{}, errPassword
	}
	s.Password = string(hashedPassword)
	//remove spaces in username
	s.Username = html.EscapeString(strings.TrimSpace(s.Username))

	var err error = db.Create(&s).Error
	if err != nil {
		return &Seller{}, err
	}
	return s, nil
}
