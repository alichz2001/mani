package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"time"
)

func AddStorageGateway(app *fiber.App) {

	jwtMiddleware := NewAuthMiddleware(Secret)

	proxyHandler := proxy.New(proxy.Config{
		Timeout: time.Second * 10,
		Servers: []string{StorageServiceURI},
		ModifyResponse: func(ctx *fiber.Ctx) error {
			return nil
		},
	})

	app.Group("/v1/storage", jwtMiddleware, proxyHandler)

}
