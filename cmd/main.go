package main

import (
	"boardgame/api/handler"
	database "boardgame/internal/api/platform/postgres"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"

	routes "boardgame/api"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

}

func main() {

	db := database.ConnectDatabase()
	if db != nil {
		fmt.Println("Conexão com o banco de dados estabelecida com sucesso.")
	} else {
		fmt.Println("Falha ao estabelecer conexão com o banco de dados.")
	}

	app := fiber.New()

	app.Use(cors.New())

	handlerGames := &handler.HandlerGames{
		DB: db,
	}

	handlerRentals := &handler.HandlerRentals{
		DB: db,
	}

	handlerCustomers := &handler.HandlerCustomers{
		DB: db,
	}

	routes.RegisterRoutes(app, handlerGames, handlerCustomers, handlerRentals)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
