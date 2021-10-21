package controllers

import (
	"net/http"
	"time"

	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	SellerId    int    `json:"seller_id"`
	CategoryId  int    `json:"category_id"`
}

type productAndCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	SellerId    int    `json:"seller_id"`
	CategoryId  int    `json:"category_id"`
	Category    string `json:"category"`
}

// GetAllProduct godoc
// @Summary Get all Product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /products [get]
func GetAllProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var product []productAndCategory

	if err := db.Raw("SELECT * FROM products JOIN categories ON products.category_id = categories.id").First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// SearchProduct godoc
// @Summary Search Product.
// @Description Search a list of Product.
// @Tags Product
// @Produce json
// @Query name path string true "Product name"
// @Success 200 {object} []models.Product
// @Router /products/search [get]
func SearchProduct(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var product []productAndCategory

	queryParams := c.Query("name")
	queryParamsLike := "%" + queryParams + "%"

	if err := db.Raw(`SELECT * FROM products JOIN categories ON products.category_id = categories.id WHERE products.name LIKE ?`, queryParamsLike).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// CreateProduct godoc
// @Summary Create New Product.
// @Description Creating a new Product.
// @Tags Product
// @Param Body body productInput true "the body to create a new Product"
// @Produce json
// @Success 200 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Category
	product := models.Product{Name: input.Name, Description: input.Description, Price: uint(input.Price), Stock: uint(input.Stock), SellerId: uint(input.SellerId), CategoryId: uint(input.CategoryId)}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductById godoc
// @Summary Get Product by Id.
// @Description Get a Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProductById(c *gin.Context) { // Get model if exist
	var product models.Product

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GetProductByCategoryId godoc
// @Summary Get Product by Category Id.
// @Description Get a Product by category id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} models.Product
// @Router /products/{id}/category [get]
func GetProductByCategoryId(c *gin.Context) { // Get model if exist
	var products []models.Product

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("category_id = ?", c.Param("id")).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Param Body body productInput true "the body to update Product"
// @Success 200 {object} models.Product
// @Router /products/{id} [patch]
func UpdateProduct(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Product
	updatedInput.Name = input.Name
	updatedInput.Description = input.Description
	updatedInput.Price = uint(input.Price)
	updatedInput.Stock = uint(input.Stock)
	updatedInput.SellerId = uint(input.SellerId)
	updatedInput.CategoryId = uint(input.CategoryId)
	updatedInput.UpdatedAt = time.Now()

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DeleteProduct godoc
// @Summary Delete one Product.
// @Description Delete a Product by id.
// @Tags Product
// @Produce json
// @Param id path string true "Product id"
// @Success 200 {object} map[string]boolean
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": "Success Delete"})
}
