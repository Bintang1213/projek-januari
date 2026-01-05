package controllers

import (
	"fmt"
	"food-gofiber/config"
	"food-gofiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetAllMenus - Tetap sama
func GetAllMenus(c *fiber.Ctx) error {
	var menus []models.Menu
	config.DB.Find(&menus)
	return c.JSON(menus)
}

// CreateMenu - Tambahkan pengecekan error parse
func CreateMenu(c *fiber.Ctx) error {
	// Ambil price dari form dan konversi
	priceStr := c.FormValue("price")
	price, _ := strconv.ParseFloat(priceStr, 64)

	menu := models.Menu{
		Name:        c.FormValue("name"),
		Description: c.FormValue("description"),
		Price:       price,
		Category:    c.FormValue("category"),
	}

	// Upload Gambar
	file, err := c.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", filename))
		menu.Image = filename
	}

	if err := config.DB.Create(&menu).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal simpan menu"})
	}
	return c.JSON(fiber.Map{"message": "Menu berhasil ditambah", "data": menu})
}

// UpdateMenu - Pastikan ID diambil dengan benar
func UpdateMenu(c *fiber.Ctx) error {
	id := c.Params("id")
	var menu models.Menu

	// Cari data lama
	if err := config.DB.First(&menu, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Menu tidak ditemukan"})
	}

	// Update field satu per satu hanya jika nilainya ada di form
	if name := c.FormValue("name"); name != "" {
		menu.Name = name
	}
	if category := c.FormValue("category"); category != "" {
		menu.Category = category
	}
	if desc := c.FormValue("description"); desc != "" {
		menu.Description = desc
	}
	
	if priceStr := c.FormValue("price"); priceStr != "" {
		price, _ := strconv.ParseFloat(priceStr, 64)
		menu.Price = price
	}

	// Update gambar jika ada file baru
	file, err := c.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", filename))
		menu.Image = filename
	}

	// Gunakan Save untuk memperbarui semua field model menu
	if err := config.DB.Save(&menu).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal update database"})
	}

	return c.JSON(fiber.Map{"message": "Menu berhasil diupdate", "data": menu})
}

// DeleteMenu - Tetap sama
func DeleteMenu(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Menu{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal hapus menu"})
	}
	return c.JSON(fiber.Map{"message": "Menu berhasil dihapus"})
}