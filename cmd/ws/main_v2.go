// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"time"
    "fmt"
	"os"
	"strconv"
	"github.com/markelca/toggles/flags"
	"github.com/markelca/toggles/storage"
)

func main() {
    appPort      := os.Getenv("APP_PORT")
    redisHost    := os.Getenv("REDIS_HOST")
    redisPortStr := os.Getenv("REDIS_PORT")
    mongoHost    := os.Getenv("MONGO_HOST")
    mongoPortStr := os.Getenv("MONGO_PORT")

     redisPort, err  := strconv.Atoi(redisPortStr)
     mongoPort, err2 := strconv.Atoi(mongoPortStr)
     if err != nil || err2 != nil{
         panic(err)
     }

    database,err := flags.NewFlagMongoRepository(mongoHost,mongoPort)
    if err != nil {
        panic("Couldn't connect to mongo!")
    }

    cacheClient := storage.NewRedisClient(redisHost, redisPort)
    flagService := flags.NewFlagService(cacheClient,database)

    controller := WSController{flagService,cacheClient}

	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, controller, w, r)
	})

    host := fmt.Sprintf(":%v", appPort)
	server := &http.Server{
		Addr:              host,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
