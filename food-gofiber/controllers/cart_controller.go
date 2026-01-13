package controllers

import (
	"food-gofiber/config"
	"food-gofiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GetCart(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var cart []models.Cart
	config.DB.Preload("Menu").Where("user_id = ?", userID).Find(&cart)
	return c.JSON(cart)
}

func AddToCart(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var input struct {
		MenuID uint `json:"menu_id"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Input tidak valid"})
	}

	var cart models.Cart
	err := config.DB.Where("user_id = ? AND menu_id = ?", userID, input.MenuID).First(&cart).Error

	if err == nil {
		cart.Quantity++
		config.DB.Save(&cart)
	} else {
		cart = models.Cart{UserID: userID, MenuID: input.MenuID, Quantity: 1}
		config.DB.Create(&cart)
	}
	return c.JSON(fiber.Map{"message": "Berhasil ditambah"})
}

func DecrementCart(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var input struct {
		MenuID uint `json:"menu_id"`
	}
	c.BodyParser(&input)

	var cart models.Cart
	if err := config.DB.Where("user_id = ? AND menu_id = ?", userID, input.MenuID).First(&cart).Error; err == nil {
		if cart.Quantity > 1 {
			cart.Quantity--
			config.DB.Save(&cart)
		} else {
			config.DB.Delete(&cart)
		}
	}
	return c.JSON(fiber.Map{"message": "Berhasil dikurangi"})
}