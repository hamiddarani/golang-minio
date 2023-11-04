package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/hamiddarani/golang-minio/api/routes"
)

func main() {
	app := fiber.New()

	routes.PublicRoutes(app)

	app.Listen(":8080")

}
