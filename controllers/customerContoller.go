package controllers

import (
	"fmt"
	"go-mall/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type ChangePasswordInput struct {
	NewPassword        string `json:"new_password" binding:"required"`
	NewPasswordConfirm string `json:"new_password_confirm" binding:"required"`
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

// UpdateCustomerPassword godoc
// @Summary Update Customer Password.
// @Description Update customer password by id.
// @Tags Auth Customer
// @Produce json
// @Param id path string true "Customer id"
// @Param Body body ChangePasswordInput true "the body to update customer password"
// @Success 200 {object} map[string]interface{}
// @Router /customer/{id} [patch]
func UpdateCustomerPassword(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var customer models.Customer
	if err := db.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.NewPassword != input.NewPasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "new password and confirm password doesn't match"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)

	passwordStr := string(hashedPassword)

	var updatedInput models.Customer
	updatedInput.Password = passwordStr

	db.Model(&customer).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}
