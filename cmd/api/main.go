package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/markelca/toggle.go/flags/storage"
    "strconv"
)


func main() {
    redisHost := os.Getenv("REDIS_HOST")
    redisPort := os.Getenv("REDIS_PORT")

     port, err := strconv.Atoi(redisPort)
     if err != nil {
         panic(err)
     }

    r := gin.Default()

    // repository := storage.NewMemoryRepository()
    repository := storage.NewRedisRepository(redisHost, port)
    controller := NewFlagController(repository)

    r.GET("/flags", controller.ListFlags)
    r.GET("/flags/:flagid", controller.GetFlag)
    r.PUT("/flags/:flagid", controller.UpdateFlag)
    r.POST("/flags", controller.CreateFlag)
    r.Run() // listen and serve on localhost:8080 
}
