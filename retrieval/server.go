package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitServer() *fiber.App {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New())

	return app
}
