package main

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type StorageRouter struct {
	srv *StorageServiceImpl
}

func AddHealthCheckRouter(app *fiber.App) {
	app.Get("/health_check", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})
}

func AddStorageRouter(app *fiber.App, storageService *StorageServiceImpl) {
	r := &StorageRouter{
		storageService,
	}
	v1Group := app.Group("/v1")
	v1Group.Post("/storage/file/upload", r.UploadFile)
	v1Group.Get("/storage/file/fetch", r.GetFile)

	app.Static("/v1/storage/files", StorageBasePath, fiber.Static{
		ModifyResponse: r.ServeFile,
	})
}

type GetFileReq struct {
	Name string
	Tags []string
}

func (r *StorageRouter) GetFile(ctx *fiber.Ctx) error {
	req := &GetFileReq{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}
	files, err := r.srv.Fetch(req)
	if err != nil {
		return err
	}

	resp, _ := json.Marshal(files)
	return ctx.Status(http.StatusOK).Send(resp)
}

func (r *StorageRouter) UploadFile(ctx *fiber.Ctx) error {

	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	uploadReq, mapErr := MapAndValidate(form)
	if mapErr != nil {
		return ctx.Status(http.StatusBadRequest).SendString(mapErr.Error())
	}

	c, cancelFunc := context.WithTimeout(ctx.Context(), time.Second*10)
	defer cancelFunc()
	file, uploadErr := r.srv.Upload(c, uploadReq)
	if uploadErr != nil {
		return ctx.Status(http.StatusBadRequest).SendString(uploadErr.Error())
	}

	resp, err := json.Marshal(file)

	return ctx.Status(http.StatusCreated).Send(resp)
}

func (r *StorageRouter) ServeFile(ctx *fiber.Ctx) error {
	//TODO

	return nil
}
