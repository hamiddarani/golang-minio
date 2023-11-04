package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hamiddarani/golang-minio/api/controller"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/upload", controller.UploadFile)
}
