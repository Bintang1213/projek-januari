package controllers

import (
	"food-gofiber/config"
	"food-gofiber/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateOrder(c *fiber.Ctx) error {
	rawUser := c.Locals("user")
	if rawUser == nil {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	userToken := rawUser.(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var req struct {
		Total         float64            `json:"total"`
		PaymentMethod string             `json:"payment"`
		OrderMethod   string             `json:"method"`
		Address       string             `json:"address"`
		Note          string             `json:"note"`
		Items         []models.OrderItem `json:"items"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Data tidak valid"})
	}

	order := models.Order{
		UserID:        userID,
		Total:         req.Total,
		Status:        "pending",
		PaymentMethod: req.PaymentMethod,
		OrderMethod:   req.OrderMethod,
		Address:       req.Address,
		Note:          req.Note,
		Items:         req.Items,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal simpan order"})
	}

	if req.PaymentMethod == "non-tunai" {
	s := config.MidtransSnap()

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "ORDER-" + strconv.Itoa(int(order.ID)),
			GrossAmt: int64(req.Total),
		},
	}

	snapResp, err := s.CreateTransaction(snapReq)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal Midtrans"})
	}

	return c.JSON(fiber.Map{
		"token":    snapResp.Token,
		"order_id": order.ID,
	})
}


	order.Status = "proses"
	config.DB.Save(&order)
	return c.JSON(fiber.Map{"order_id": order.ID, "success": true})
}

func GetOrderDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.Order
	
	// Preload User sekarang akan bekerja karena field User sudah ada di Model
	if err := config.DB.Preload("Items").Preload("User").First(&order, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Order tidak ditemukan"})
	}

	config.DB.Where("user_id = ?", order.UserID).Delete(&models.Cart{})

	return c.JSON(order)
}

func HandleMidtransNotification(c *fiber.Ctx) error {
	var notification map[string]interface{}
	if err := c.BodyParser(&notification); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	orderIDFull := notification["order_id"].(string)
	orderIDStr := strings.Replace(orderIDFull, "ORDER-", "", 1)
	orderID, _ := strconv.Atoi(orderIDStr)
	transactionStatus, _ := notification["transaction_status"].(string)

	var order models.Order
	if err := config.DB.First(&order, orderID).Error; err != nil {
		return c.Status(404).SendString("Not Found")
	}

	if transactionStatus == "settlement" || transactionStatus == "capture" {
		order.Status = "lunas"
	} else if transactionStatus == "expire" || transactionStatus == "cancel" {
		order.Status = "gagal"
	}

	config.DB.Save(&order)
	return c.Status(200).SendString("OK")
}

func GetMyOrders(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var orders []models.Order
	config.DB.Preload("Items").Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&orders)
	return c.JSON(orders)
}