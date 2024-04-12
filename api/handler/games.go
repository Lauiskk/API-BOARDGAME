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
			"error": "Erro ao buscar todos os games",
		})
	}
	return c.JSON(games)
}

func (h *HandlerGames) CreateGame(c *fiber.Ctx) error {
	game := new(model.Game)
	if err := c.BodyParser(game); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados do jogo inválidos",
		})
	}

	if game.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O nome do jogo não pode estar vazio",
		})
	}

	if game.Price == 0 && game.Stock == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O preço e o stock não podem ser 0",
		})
	}

	var existingGame model.Game

	result := h.DB.Where("Name = ?", game.Name).First(&existingGame)
	if result.RowsAffected != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Este Jogo já existe",
		})
	}

	h.DB.Create(&game)
	return c.Status(fiber.StatusCreated).JSON(game)
}
