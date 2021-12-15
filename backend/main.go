package main

import (
	"fmt"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../.env")

	fmt.Println("Starting server...")

	app := fiber.New()
	port := os.Getenv("PORT")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Starting server on port: " + port)

	mongoClient := database.ConnectDB()

	mongoClient.Database(os.Getenv("DB_NAME"))
	mongoClient.Collection(os.Getenv("DB_TAX_COLLECTION"))

	app.Listen(":" + port)
}
