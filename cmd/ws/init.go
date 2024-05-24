package main

import (
	"fmt"
	"github.com/markelca/toggles/internal/envs"
	"github.com/markelca/toggles/internal/websocket"
	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
	"log"
	"net/http"
	"time"
)

func Init() error {
	params, paramErr := envs.GetConnectionParams()
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

	controller := websocket.WSController{FlagService: flagService, CacheClient: cacheClient}

	hub := websocket.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(hub, controller, w, r)
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
