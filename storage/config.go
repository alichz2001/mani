package main

import (
	"log"
	"os"
	"strconv"
)

var (
	Port            string
	StorageBasePath string
	MaxBodyLength   int

	ArangoUsername string
	ArangoPassword string
	ArangoURI      string
	ArangoDB       string
	MaxStorageSize int64
	MaxFileSize    int64
)

func init() {
	MaxBodyLength = 1024 * 1024 * 100

	Port = os.Getenv("PORT")

	ArangoUsername = os.Getenv("ARANGO_USERNAME")
	ArangoPassword = os.Getenv("ARANGO_PASSWORD")
	ArangoURI = os.Getenv("ARANGO_URI")
	ArangoDB = os.Getenv("ARANGO_DATABASE")

	StorageBasePath = os.Getenv("FS_BASE_PATH")

	var err error
	tmp, err := strconv.Atoi(os.Getenv("MAX_STORAGE_SIZE"))
	MaxStorageSize = int64(tmp)
	tmp, err = strconv.Atoi(os.Getenv("MAX_FILE_SIZE"))
	MaxFileSize = int64(tmp)
	if err != nil {
		log.Fatalf("bad config, error: %s", err.Error())
	}
}
