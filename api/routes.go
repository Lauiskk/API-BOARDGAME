package routes

import (
	handler "boardgame/api/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handlerGames *handler.HandlerGames, handlerCustomers *handler.HandlerCustomers, handleRentals *handler.HandlerRentals) {
	api := app.Group("/api")
	api.Get("/games", handlerGames.GetAllGames)
	api.Post("/games", handlerGames.CreateGame)

	api.Get("/customers", handlerCustomers.GetAllCustomers)
	api.Get("/customers/:id", handlerCustomers.GetCustomerByID)
	api.Post("/customers", handlerCustomers.CreateCustomer)
	api.Put("/customers/:id", handlerCustomers.UpdateCustomer)

	api.Get("/rentals", handleRentals.GetAllRentals)
	api.Post("/rentals", handleRentals.CreateRental)

}
