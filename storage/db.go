package main

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type ManiDB struct {
	Database          driver.Database
	StorageCollection driver.Collection
}

func InitArangoDB() *ManiDB {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{ArangoURI},
	})
	if err != nil {
		panic("arangodb connection error")
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(ArangoUsername, ArangoPassword),
	})
	if err != nil {
		panic("arangodb init error")
	}

	ctx := context.Background()
	maniDB, err := c.Database(ctx, ArangoDB)
	if err != nil {
		panic(err)
	}
	storageCollection, err := maniDB.Collection(ctx, "files")
	if err != nil {
		panic(err)
	}
	return &ManiDB{maniDB, storageCollection}
}
