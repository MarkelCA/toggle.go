package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/markelca/toggles/internal/envs"
	"github.com/markelca/toggles/internal/websocket"
	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
)

func Init() error {
	params, paramErr := envs.GetConnectionParams(envs.EnvNames{
		AppPort:   "APP_PORT",
		RedisHost: "REDIS_HOST",
		RedisPort: "REDIS_PORT",
		MongoHost: "MONGO_HOST",
		MongoPort: "MONGO_PORT",
	})
	if len(paramErr) > 0 {
		errMsg := "Param errors have been found:\n"
		for _, err := range paramErr {
			errMsg += fmt.Sprintf("  - %v\n", err.Error())
		}
		log.Fatal(errMsg)
	}

	database, err := flags.NewFlagMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		return err
	}

	cacheClient := storage.NewRedisClient(params.RedisHost, params.RedisPort)
	flagService := flags.NewFlagService(cacheClient, database)

	hub := websocket.NewHub()
	go hub.Run()

	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		controller := websocket.ControllerV2{FlagService: flagService, CacheClient: cacheClient}
		actionMarshaller := websocket.JsonMarshaller{}
		websocket.ServeWs(hub, controller, w, r, actionMarshaller)
	})

	host := fmt.Sprintf(":%v", params.AppPort)
	server := &http.Server{
		Addr:              host,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
