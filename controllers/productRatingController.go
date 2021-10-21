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

// CreateRating godoc
// @Summary Create New Rating.
// @Description Creating a new Rating.
// @Tags Rating
// @Param Body body productRatingInput true "the body to create a new Rating"
// @Produce json
// @Success 200 {object} models.ProductRating
// @Router /products/rating [post]
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

// GetRatingById godoc
// @Summary Get Rating by Id.
// @Description Get a Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Success 200 {object} models.ProductRating
// @Router /products/rating/{id} [get]
func GetRatingById(c *gin.Context) { // Get model if exist
	var rating models.ProductRating

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetRatingByProductId godoc
// @Summary Get Rating by Product Id.
// @Description Get a Rating by Product id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Success 200 {object} models.ProductRating
// @Router /products/{id}/rating [get]
func GetRatingByProductId(c *gin.Context) { // Get model if exist
	var rating []models.ProductRating

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("product_id = ?", c.Param("id")).Find(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// UpdateRating godoc
// @Summary Update Rating.
// @Description Update Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Param Body body productRatingInput true "the body to update Rating"
// @Success 200 {object} models.ProductRating
// @Router /produts/ratings/{id} [patch]
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

// DeleteRating godoc
// @Summary Delete one Rating.
// @Description Delete a Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Success 200 {object} map[string]boolean
// @Router /products/rating/{id} [delete]
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
