package routes

import (
	"food-gofiber/controllers"
	"food-gofiber/middleware"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Get("/menu", controllers.GetAllMenus) 

	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("SECRET_KEY_KAMU")}, 
	}))

	// Group khusus Admin
	admin := api.Group("/admin", middleware.IsAdmin)
	admin.Post("/menu", controllers.CreateMenu)
	admin.Put("/menu/:id", controllers.UpdateMenu)
	admin.Delete("/menu/:id", controllers.DeleteMenu)
}