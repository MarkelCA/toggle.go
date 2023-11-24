package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

var flags []Flag = make([]Flag,0)

func main() {
    r := gin.Default()

    r.GET("/flags", func(c *gin.Context) {
        c.JSON(http.StatusOK, flags)
    })

    r.GET("/flags/:flagid", func(c *gin.Context) {
        var result bool
        for _,flag := range flags {
            if flag.Name == c.Params.ByName("flagid") {
                result = flag.Value
            }
        }
        c.JSON(http.StatusOK, result)
    })

    r.PUT("/flags/:flagid", func(c *gin.Context) {
        var body struct{value bool}
        var found bool
        c.Bind(&body)
        for _,flag := range flags {
            if flag.Name == c.Params.ByName("flagid") {
                flag.Value = body.value
                found = true
            }
        }

        fmt.Println(flags)

        if !found {
            c.Status(http.StatusNotFound)
        }
    })

    r.POST("/flags", func(c *gin.Context) {
        var flag Flag
        err := c.BindJSON(&flag)
        if err != nil {
            log.Println("Error!", err)
            return
        }
        for _,currentFlag := range flags {
            if currentFlag.Name == flag.Name {
                c.Status(http.StatusConflict)
                return
            }
        }

        flags = append(flags,flag)
        c.Status(http.StatusCreated)
    })

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
