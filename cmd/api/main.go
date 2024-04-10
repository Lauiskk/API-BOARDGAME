package main

import (
	database "boardgame/internal/api/platform/postgres"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	// Aqui você pode definir suas rotas

	app.Listen(":3000")
}
