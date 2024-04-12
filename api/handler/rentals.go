package handler

import (
	"boardgame/api/model"
	"time"

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
			"error": "Erro ao buscar todos os alugueis",
		})
	}
	return c.JSON(rentals)
}

func (h *HandlerRentals) CreateRental(c *fiber.Ctx) error {
	rental := new(model.Rental)
	if err := c.BodyParser(rental); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados do aluguel inválidos",
		})
	}

	var customer model.Customer
	result := h.DB.First(&customer, rental.CustomerID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Erro ao buscar cliente, cliente não encontrado",
		})
	}

	var game model.Game
	result = h.DB.First(&game, rental.GameID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Erro ao buscar jogo, jogo não encontrado",
		})
	}

	if game.Stock <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Desculpe, este jogo está fora de estoque",
		})
	}

	game.Stock -= 1
	h.DB.Save(game.Stock)

	if rental.DaysRented == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Não é possível alugar por 0 dias",
		})
	}

	rental.Customer = customer
	rental.Game = game

	rental.OriginalPrice = game.Price * rental.DaysRented
	rental.RentDate = time.Now().Format("2006-12-17")

	h.DB.Create(&rental)
	return c.Status(fiber.StatusCreated).JSON(rental)
}
