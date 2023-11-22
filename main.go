package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

var flags []Flag

type Foo struct {
    Foo string `json:"foo"`
    Fizz string `json:"fizz"`
}

func main() {
    r := gin.Default()

    r.GET("/flags", func(c *gin.Context) {
        c.JSON(http.StatusOK, flags)
    })

    r.GET("/flags/:flagid", func(c *gin.Context) {
        // c.JSON(http.StatusOK, flags)
    })

    r.POST("/flags", func(c *gin.Context) {
        var flag Flag
        err := c.BindJSON(&flag)
        if err != nil {
            log.Println("Error!")
            return
        }
        flags = append(flags,flag)
        c.JSON(http.StatusCreated, flag)
    })

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
