package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-mall/controllers"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Categories Routes
	r.GET("/categories", controllers.GetAllCategory)
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories/:id", controllers.GetCategoryById)
	// r.GET("/categories/:id/movies", controllers.GetMoviesByRatingId)
	r.PATCH("/categories/:id", controllers.UpdateCategory)
	r.DELETE("categories/:id", controllers.DeleteCategory)

	// Product Routes
	r.GET("/products", controllers.GetAllProduct)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProductById)
	r.GET("/products/:id/category", controllers.GetProductByCategoryId)
	r.PATCH("/products/:id", controllers.UpdateProduct)
	r.DELETE("products/:id", controllers.DeleteProduct)

	// Product Comment Routes
	r.POST("/products/rating", controllers.CreateRating)
	r.GET("/products/rating/:id", controllers.GetRatingById)
	r.GET("/products/:id/rating", controllers.GetRatingByProductId)
	r.PATCH("/products/rating/:id", controllers.UpdateRating)
	r.DELETE("products/rating/:id", controllers.DeleteRating)

	// Order Routes
	r.GET("/order", controllers.GetAllOrder)
	r.GET("/order/:id", controllers.GetOrderById)
	r.POST("/order", controllers.CreateOrder)
	r.PATCH("/order/:id", controllers.UpdateOrderStatus)
	r.DELETE("/order/:id", controllers.DeleteOrder)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
