package controllers

import (
	"net/http"
	"time"

	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productRatingInput struct {
	ProductId  uint   `json:"product_id"`
	CustomerId uint   `json:"customer_id"`
	Rating     uint   `json:"rating"`
	Comment    string `json:"comment"`
}

func CreateRating(c *gin.Context) {
	// Validate input
	var input productRatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product ratomg
	rating := models.ProductRating{ProductId: input.ProductId, CustomerId: input.CustomerId, Rating: input.Rating, Comment: input.Comment}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func GetRatingById(c *gin.Context) { // Get model if exist
	var rating models.ProductRating

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func GetRatingByProductId(c *gin.Context) { // Get model if exist
	var rating []models.ProductRating

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("product_id = ?", c.Param("id")).Find(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func UpdateRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var rating models.ProductRating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productRatingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.ProductRating
	updatedInput.ProductId = uint(input.ProductId)
	updatedInput.CustomerId = uint(input.CustomerId)
	updatedInput.Rating = uint(input.Rating)
	updatedInput.Comment = input.Comment
	updatedInput.UpdatedAt = time.Now()

	db.Model(&rating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

func DeleteRating(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var rating models.ProductRating
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&rating)

	c.JSON(http.StatusOK, gin.H{"data": "Success Delete"})
}
