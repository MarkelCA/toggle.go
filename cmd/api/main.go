package main

import (
	"fmt"
	"github.com/markelca/toggles/flags"
	"github.com/markelca/toggles/storage"
)

func main() {
	params, paramErr := GetConnectionParams()
	if paramErr != nil {
		panic(fmt.Sprintf("Param errors have been found: %v", paramErr))
	}

	// repository := storage.NewMemoryRepository()
	db, err := flags.NewFlagMongoRepository(params.mongoHost, params.mongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}

	repository := storage.NewRedisClient(params.redisHost, params.redisPort)
	service := flags.NewFlagService(repository, db)
	controller := NewFlagController(service)
	host := fmt.Sprintf(":%v", params.appPort)
	controller.Init(host)
}
