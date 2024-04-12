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
	result := h.DB.Preload("Customer").Preload("Game").Find(&rentals)
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
	h.DB.Save(&game)

	if rental.DaysRented == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Não é possível alugar por 0 dias",
		})
	}

	rental.Customer = customer
	rental.Game = game

	rental.OriginalPrice = game.Price * rental.DaysRented
	rental.RentDate = time.Now().Format("2006-01-02")

	h.DB.Create(&rental)
	return c.Status(fiber.StatusCreated).JSON(rental)
}

func (h *HandlerRentals) FinalizeRental(c *fiber.Ctx) error {
	var rental model.Rental
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(err.Error())
	}

	if err := h.DB.Preload("Game").Where("id = ?", id).First(&rental).Error; err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Erro ao buscar aluguel inexistente",
		})
	}

	if rental.ReturnDate != "" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "error",
		})
	}

	rental.ReturnDate = time.Now().Format("2006-01-02")

	rentDate, err := time.Parse("2006-01-02", rental.RentDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao converter a data de aluguel",
		})
	}
	expectedReturnDate := rentDate.AddDate(0, 0, int(rental.DaysRented))
	daysLate := int(time.Since(expectedReturnDate).Hours() / 24)
	if daysLate > 0 {
		rental.DelayFee = uint(daysLate) * rental.Game.Price
	}

	h.DB.Save(&rental)
	return c.Status(fiber.StatusOK).JSON(&rental)
}

func (h *HandlerRentals) DeleteRental(c *fiber.Ctx) error {
	id := c.Params("id")
	var rental model.Rental
	result := h.DB.First(&rental, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).SendString("Sem aluguel com esse ID " + id)
	}

	if rental.ReturnDate == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Erro, aluguel não fechado",
		})
	}
	h.DB.Delete(&rental)

	return c.SendString("Aluguel deletado com sucesso")
}
