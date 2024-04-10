package routes

import (
	handler "boardgame/api/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, handlerGames *handler.HandlerGames, handlerCustomers *handler.HandlerCustomers, handleRentals *handler.HandlerRentals) {
	api := app.Group("/api")
	api.Get("/games", handlerGames.GetAllGames)
	api.Get("/customers", handlerCustomers.GetAllCustomers)
	api.Get("/rentals", handleRentals.GetAllRentals)

}
