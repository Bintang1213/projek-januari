package routes

import (
    "food-gofiber/controllers"
    "food-gofiber/middleware"

    jwtware "github.com/gofiber/contrib/jwt"
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    api := app.Group("/api")

    // --- 1. PUBLIC ROUTES ---
    api.Post("/register", controllers.Register)
    api.Post("/login", controllers.Login)
    api.Get("/menu", controllers.GetAllMenus)
    
    // Webhook harus Public (Tanpa JWT)
    api.Post("/midtrans-webhook", controllers.HandleMidtransNotification)

    // --- 2. MIDDLEWARE JWT ---
    api.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte("SECRET_KEY_KAMU")}, 
    }))

    // --- 3. PROTECTED ROUTES ---
    api.Get("/cart", controllers.GetCart)
    api.Post("/cart", controllers.AddToCart)
    api.Post("/cart/decrement", controllers.DecrementCart)
    
    // --- FITUR ORDER ---
    api.Post("/order", controllers.CreateOrder)
    api.Get("/my-orders", controllers.GetMyOrders)
    
    api.Get("/order/:id", controllers.GetOrderDetail) 

    // --- 4. ADMIN ROUTES ---
    admin := api.Group("/admin", middleware.IsAdmin)
    admin.Post("/menu", controllers.CreateMenu)
    admin.Put("/menu/:id", controllers.UpdateMenu)
    admin.Delete("/menu/:id", controllers.DeleteMenu)
}