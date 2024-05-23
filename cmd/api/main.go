package main

import (
	"fmt"
	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
)

func main() {
	params, paramErr := GetConnectionParams()
	if len(paramErr) > 0 {
		errMsg := "Param errors have been found:\n"
		for _, err := range paramErr {
			errMsg += fmt.Sprintf("  - %v\n", err.Error())
		}
		panic(errMsg)
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
