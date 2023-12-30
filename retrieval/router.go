package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var myValidator *validator.Validate

type AuthRouter struct {
	userSrv *UserServiceImpl
}

func AddHealthCheckRouter(app *fiber.App) {
	app.Get("/health_check", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})
}

func AddAuthRouter(app *fiber.App, AuthService *UserServiceImpl) {
	myValidator = validator.New()
	r := &AuthRouter{
		AuthService,
	}
	v1Group := app.Group("/v1")
	v1Group.Post("/auth/signin", r.SignIn)
	v1Group.Post("/auth/signup", r.SignUp)
}

func (r *AuthRouter) SignIn(ctx *fiber.Ctx) error {
	req := &User{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	user, ok := r.userSrv.ValidateUserCredential(req)
	if !ok {
		return ctx.Status(http.StatusNotFound).SendString("Username not found")
	}

	token, err := GenerateTokenForUser(user)
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	out, _ := json.Marshal(map[string]string{"jwt": token})
	return ctx.Status(http.StatusOK).Send(out)
}

func (r *AuthRouter) SignUp(ctx *fiber.Ctx) error {
	req := &User{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	err := myValidator.Struct(req)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	user, err := r.userSrv.Register(req)
	if err != nil {
		return err
	}

	token, err := GenerateTokenForUser(user)
	if err != nil {
		return err
	}

	out, _ := json.Marshal(map[string]string{"jwt": token})
	return ctx.Status(http.StatusOK).Send(out)
}
