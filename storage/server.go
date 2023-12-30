package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitServer() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: MaxBodyLength,
		//ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		//	return ctx.SendStatus(http.StatusInternalServerError)
		//},
	})
	app.Use(requestid.New())
	app.Use(logger.New())

	return app
}
