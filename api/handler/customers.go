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
			"error": "Erro ao buscar todos os clientes",
		})
	}
	return c.JSON(customers)
}

func (h *HandlerCustomers) GetCustomerByID(c *fiber.Ctx) error {
	var customer model.Customer
	id := c.Params("id")
	result := h.DB.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Erro ao buscar usuario, usuario não encontrado",
		})
	}

	return c.JSON(customer)
}

func (h *HandlerCustomers) CreateCustomer(c *fiber.Ctx) error {
	customer := new(model.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados de cliente inválidos",
		})
	}

	if customer.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O nome do cliente não pode estar vazio",
		})
	}

	if len(customer.CPF) != 11 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O CPF deve ter 11 caracteres numéricos",
		})
	}

	if len(customer.Phone) != 10 && len(customer.Phone) != 11 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O telefone deve ter 10 ou 11 caracteres numéricos",
		})
	}

	if customer.Birthday.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "A data de aniversário deve ser uma data válida",
		})
	}

	var existingClientCpf model.Customer

	result := h.DB.Where("CPF = ?", customer.CPF).First(&existingClientCpf)
	if result.RowsAffected != 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Este CPF já está cadastrado",
		})
	}

	h.DB.Create(&customer)
	return c.Status(fiber.StatusCreated).JSON(customer)
}

func (h *HandlerCustomers) UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer model.Customer
	result := h.DB.First(&customer, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error ": "Sem customers com esse ID",
		})
	}

	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error",
		})
	}

	var existingClientCpf model.Customer

	cpfConsult := h.DB.Where("CPF = ?", customer.CPF).First(&existingClientCpf)
	if cpfConsult.RowsAffected != 0 && existingClientCpf.ID != customer.ID {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Este CPF já está cadastrado em outro cliente",
		})
	}

	if len(customer.CPF) != 11 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O CPF deve ter 11 caracteres numéricos",
		})
	}

	if len(customer.Phone) != 10 && len(customer.Phone) != 11 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "O telefone deve ter 10 ou 11 caracteres numéricos",
		})
	}
	if customer.Birthday.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "A data de aniversário deve ser uma data válida",
		})
	}

	h.DB.Save(&customer)
	return c.JSON(&customer)
}
