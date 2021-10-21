package controllers

import (
	"fmt"
	"go-mall/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginSeller godoc
// @Summary Login as as seller.
// @Description Logging in to get jwt token to access seller api.
// @Tags Auth Seller
// @Param Body body LoginInput true "the body to login a seller"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /seller/login [post]
func LoginSeller(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Seller{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheckSeller(u.Username, u.Password, db)

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

// RegisterSeller godoc
// @Summary Register a seller.
// @Description registering a seller from public access.
// @Tags Auth Seller
// @Param Body body RegisterInput true "the body to register a seller"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /seller/register [post]
func RegisterSeller(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Seller{}

	u.Username = input.Username
	u.Email = input.Email
	u.Password = input.Password

	_, err := u.SaveUserSeller(db)

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

// UpdateSellerPassword godoc
// @Summary Update Seller Password.
// @Description Update seller password by id.
// @Tags Auth Seller
// @Produce json
// @Param id path string true "Seller id"
// @Param Body body ChangePasswordInput true "the body to update seller password"
// @Success 200 {object} map[string]interface{}
// @Router /seller/{id} [patch]
func UpdateSellerPassword(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var seller models.Seller
	if err := db.Where("id = ?", c.Param("id")).First(&seller).Error; err != nil {
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

	var updatedInput models.Seller
	updatedInput.Password = passwordStr

	db.Model(&seller).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": seller})
}
