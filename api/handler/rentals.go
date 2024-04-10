package handler

import (
	"boardgame/api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HandlerRentals struct {
	DB *gorm.DB
}

func (h *HandlerRentals) GetAllRentals(c *fiber.Ctx) error {
	var rentals []model.Rental
	result := h.DB.Find(&rentals)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar todos os usuarios",
		})
	}
	return c.JSON(rentals)
}
