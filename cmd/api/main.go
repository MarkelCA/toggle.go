package main

import (
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/markelca/toggles/storage"
	"github.com/markelca/toggles/flags"
)


func main() {
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
    db,err := storage.NewMongoClient(mongoHost,mongoPort)
    if err != nil {
        panic("Couldn't connect to mongo!")
    }
    db.Get("aoeu")
    repository := storage.NewRedisClient(redisHost, redisPort)
    service := flags.NewFlagService(repository,db)
    controller := NewFlagController(service)

    r := gin.Default()
    r.GET("/flags", controller.ListFlags)
    r.GET("/flags/:flagid", controller.GetFlag)
    r.PUT("/flags/:flagid", controller.UpdateFlag)
    r.POST("/flags", controller.CreateFlag)
    r.Run() // listen and serve on localhost:8080 
}
