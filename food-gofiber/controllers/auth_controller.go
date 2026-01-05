package controllers

import (
	"food-gofiber/config"
	"food-gofiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
	}

	// Enkripsi password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
		Role:     data["role"], 
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal membuat user/Email sudah ada"})
	}

	return c.JSON(fiber.Map{"message": "Berhasil daftar!", "user": user})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	c.BodyParser(&data)

	var user models.User
	config.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])) != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Email atau Password salah"})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("SECRET_KEY_KAMU")) 

	return c.JSON(fiber.Map{
		"token": t,
		"role":  user.Role,
		"name":  user.Name,
	})
}