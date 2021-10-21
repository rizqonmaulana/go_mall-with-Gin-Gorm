package controllers

import (
	"fmt"
	"go-mall/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// LoginCustomer godoc
// @Summary Login as as customer.
// @Description Logging in to get jwt token to access customer api.
// @Tags Auth Customer
// @Param Body body LoginInput true "the body to login a customer"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /customer/login [post]
func LoginCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Customer{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheckCustomer(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	user := map[string]string{
		"username": u.Username,
		"email":    u.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "user": user, "token": token})

}

// RegisterCustomer godoc
// @Summary Register a Customer.
// @Description registering a customer from public access.
// @Tags Auth Customer
// @Param Body body RegisterInput true "the body to register a customer"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /customer/register [post]
func RegisterCustomer(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Customer{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password
	_, err := u.SaveUserCustomer(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
		"email":    input.Email,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}
