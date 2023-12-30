package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"golang.org/x/time/rate"
	"net/http"
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
		ModifyRequest: upstreamServiceRateLimiter(modifyRequest),
	})

	app.Group("/v1/storage", jwtMiddleware, proxyHandler)

}

func modifyRequest(ctx *fiber.Ctx) error {
	//TODO maybe some security checks
	return nil
}

func upstreamServiceRateLimiter(next func(ctx *fiber.Ctx) error) func(*fiber.Ctx) error {
	rl := rate.NewLimiter(5, 10)
	return func(ctx *fiber.Ctx) error {
		if !rl.Allow() {

			return ctx.SendStatus(http.StatusTooManyRequests)
		} else {
			return next(ctx)
		}
	}
}
