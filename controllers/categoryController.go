package controllers

import (
	"net/http"
	"time"

	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryInput struct {
	Category string `json:"category"`
}

// GetAllCategory godoc
// @Summary Get all Category.
// @Description Get a list of Category.
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /categories [get]
func GetAllCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var category []models.Category
	db.Find(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// CreateCategory godoc
// @Summary Create New Category.
// @Description Creating a new Category.
// @Tags Category
// @Param Body body categoryInput true "the body to create a new Category"
// @securityDefinitions.apikey ApiKeyAuth
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Category
	category := models.Category{Category: input.Category}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetCategoryById godoc
// @Summary Get Category by Id.
// @Description Get an Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategoryById(c *gin.Context) { // Get model if exist
	var category models.Category

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Param Body body categoryInput true "the body to update category"
// @Success 200 {object} models.Category
// @Router /categories/{id} [patch]
func UpdateCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.Category = input.Category
	updatedInput.UpdatedAt = time.Now()

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory godoc
// @Summary Delete one Category.
// @Description Delete a Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": "Success Delete"})
}
