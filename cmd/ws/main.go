package main

import (
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

    host := fmt.Sprintf(":%v", appPort)
    controller.Init(host)
}
