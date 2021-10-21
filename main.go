package main

import (
	"go-mall/config"
	"go-mall/docs"
	"go-mall/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @termsOfService http://swagger.io/terms/

	// for load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Go-Mall Swagger API Documentation"
	docs.SwaggerInfo.Description = "This is a sample server Go-Mall."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)
	r.Run()
}
