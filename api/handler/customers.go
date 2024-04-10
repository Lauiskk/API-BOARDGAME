package handler

import (
	"boardgame/api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HandlerCustomers struct {
	DB *gorm.DB
}

func (h *HandlerCustomers) GetAllCustomers(c *fiber.Ctx) error {
	var customers []model.Customer
	result := h.DB.Find(&customers)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar todos os usuarios",
		})
	}
	return c.JSON(customers)
}
