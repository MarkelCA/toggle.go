package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
    r := gin.Default()

    r.GET("/flags", ListFlags)
    r.GET("/flags/:flagid", FindFlag)
    r.PUT("/flags/:flagid", UpdateFlag)
    r.POST("/flags", CreateFlag)
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
