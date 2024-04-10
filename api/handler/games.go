package handler

import (
	"boardgame/api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HandlerGames struct {
	DB *gorm.DB
}

func (h *HandlerGames) GetAllGames(c *fiber.Ctx) error {
	var games []model.Game
	result := h.DB.Find(&games)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar todos os usuarios",
		})
	}
	return c.JSON(games)
}
