package controllers

import (
	"net/http"
	"time"

	"go-mall/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-mall/helper"
)

type getOrderDetail struct {
	ProductId  int    `json:"product_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Qty        int    `json:"qty"`
	SellerId   int    `json:"seller_id"`
	CustomerId int    `json:"customer_id"`
}

type orderStatusInput struct {
	OrderStatus string `json:"order_status"`
}

type orderInput struct {
	SellerId        int                `json:"seller_id"`
	CustomerId      int                `json:"customer_id"`
	TotalPrice      int                `json:"total_price"`
	ProductDetailId []orderDetailInput `json:"product_detail_id"`
}

type orderDetailInput struct {
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

func GetAllOrder(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var order []models.Order
	db.Find(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrder(c *gin.Context) {
	// Validate input
	var input orderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var invoice = helper.RandSeq(10)

	// Create Order
	order := models.Order{SellerId: uint(input.SellerId), CustomerId: uint(input.CustomerId), TotalPrice: uint(input.TotalPrice), Invoice: invoice, OrderStatus: "Pending"}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&order)

	for i := 0; i < len(input.ProductDetailId); i++ {
		productId := uint(input.ProductDetailId[i].ProductId)
		qty := uint(input.ProductDetailId[i].Qty)

		orderDetail := models.OrderDetail{OrderId: uint(order.ID), ProductId: productId, Qty: qty}
		db := c.MustGet("db").(*gorm.DB)
		db.Create(&orderDetail)

		var product models.Product

		if err := db.Where("id = ?", productId).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		qtyNow := product.Stock
		qtyNow -= qty

		var updatedInput models.Product
		updatedInput.Stock = qtyNow

		db.Model(&product).Updates(updatedInput)

	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func GetOrderById(c *gin.Context) { // Get model if exist
	var getOrderDetail []getOrderDetail

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Raw("SELECT * FROM orders JOIN order_details ON orders.id = order_details.order_id JOIN products ON order_details.product_id = products.id WHERE orders.id = ?", c.Param("id")).First(&getOrderDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// db.Model(&order{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})

	c.JSON(http.StatusOK, gin.H{"data": getOrderDetail})
}

// func GetProductByCategoryId(c *gin.Context) { // Get model if exist
// 	var products []models.Product

// 	db := c.MustGet("db").(*gorm.DB)

// 	if err := db.Where("category_id = ?", c.Param("id")).Find(&products).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": products})
// }

func UpdateOrderStatus(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var order models.Order
	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input orderStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Order
	updatedInput.OrderStatus = input.OrderStatus
	updatedInput.UpdatedAt = time.Now()

	db.Model(&order).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func DeleteOrder(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)

	var order models.Order
	var orderDetail []models.OrderDetail

	if err := db.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Where("order_id = ?", c.Param("id")).Find(&orderDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&orderDetail)
	db.Delete(&order)

	c.JSON(http.StatusOK, gin.H{"data": "Success Delete"})
}
