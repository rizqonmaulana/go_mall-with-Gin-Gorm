package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-mall/controllers"
	"go-mall/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Auth
	r.POST("/customer/register", controllers.RegisterCustomer)
	r.POST("/customer/login", controllers.LoginCustomer)

	r.POST("/seller/register", controllers.RegisterSeller)
	r.POST("/seller/login", controllers.LoginSeller)

	// Categories Routes
	r.GET("/categories", controllers.GetAllCategory)
	r.GET("/categories/:id", controllers.GetCategoryById)
	categoriesMiddlewareRoute := r.Group("/categories")
	categoriesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	categoriesMiddlewareRoute.POST("/", controllers.CreateCategory)
	categoriesMiddlewareRoute.PATCH("/:id", controllers.UpdateCategory)
	categoriesMiddlewareRoute.DELETE("/:id", controllers.DeleteCategory)

	// Product Routes
	r.GET("/products", controllers.GetAllProduct)
	r.GET("/products/search", controllers.SearchProduct)
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
