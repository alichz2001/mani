package main

import (
	"log"
)

func main() {

	app := InitServer()

	arango := InitArangoDB()

	storageService := NewStorageService(arango)

	AddHealthCheckRouter(app)
	AddStorageRouter(app, storageService)

	log.Fatal(app.Listen(":" + Port))
}
