package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markelca/toggle.go/flags/storage"
)


func main() {
    r := gin.Default()

    repository := storage.NewMemoryRepository()
    controller := NewFlagController(repository)

    r.GET("/flags", controller.ListFlags)
    r.GET("/flags/:flagid", controller.FindFlag)
    r.PUT("/flags/:flagid", controller.UpdateFlag)
    r.POST("/flags", controller.CreateFlag)
    r.Run() // listen and serve on localhost:8080 
}
