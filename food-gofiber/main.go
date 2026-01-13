package main

import (
	"food-gofiber/config"
	"food-gofiber/models"
	"food-gofiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Koneksi Database
	config.ConnectDB()

	
	err := config.DB.AutoMigrate(
		&models.User{}, 
		&models.Menu{}, 
		&models.Order{}, 
		&models.OrderItem{},
		&models.Cart{},
	)
	
	if err != nil {
		panic("Gagal melakukan migrasi database: " + err.Error())
	}

	app := fiber.New()

	// Middleware
	app.Use(logger.New()) 
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", 
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Static("/uploads", "./public/uploads")

	routes.Setup(app)

	app.Listen(":3001")
}