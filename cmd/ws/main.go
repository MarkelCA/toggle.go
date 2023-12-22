package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/gorilla/websocket"
	"github.com/markelca/toggles/flags"
	"github.com/markelca/toggles/storage"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan interface{})
var upgrader = websocket.Upgrader{
    CheckOrigin: customUpgrader,
}

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

    // repository := storage.NewMemoryRepository()
    db,err := flags.NewFlagMongoRepository(mongoHost,mongoPort)
    if err != nil {
        panic("Couldn't connect to mongo!")
    }

    repository := storage.NewRedisClient(redisHost, redisPort)
    service := flags.NewFlagService(repository,db)

    fmt.Println(repository,service)

    host := fmt.Sprintf(":%v", appPort)
    InitWS(host)
}
