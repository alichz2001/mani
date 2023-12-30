package main

import "os"

var (
	StorageServiceURI string
	Secret            string
	Port              string

	ArangoUsername string
	ArangoPassword string
	ArangoURI      string
	ArangoDB       string
)

func init() {
	StorageServiceURI = os.Getenv("STORAGE_SERVICE_URI")
	Secret = os.Getenv("SECRET_TOKEN")
	Port = os.Getenv("PORT")

	ArangoUsername = os.Getenv("ARANGO_USERNAME")
	ArangoPassword = os.Getenv("ARANGO_PASSWORD")
	ArangoURI = os.Getenv("ARANGO_URI")
	ArangoDB = os.Getenv("ARANGO_DATABASE")
}
