package main

import (
	"log"
)

func main() {

	app := InitServer()

	arango := InitArangoDB()

	storageService := NewUserService(arango)

	AddHealthCheckRouter(app)
	AddAuthRouter(app, storageService)

	AddStorageGateway(app)

	log.Fatal(app.Listen(":" + Port))
}
