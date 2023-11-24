package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markelca/toggle.go/flags"
)

func ListFlags(c *gin.Context) {
    c.JSON(http.StatusOK, flags.List())
}

func FindFlag(c *gin.Context) {
    for _,flag := range flags.List() {
        if flag.Name == c.Params.ByName("flagid") {
            c.JSON(http.StatusOK, flag.Value)
            return
        }
    }
    c.Status(http.StatusNotFound)
}
func UpdateFlag(c *gin.Context) {
    var body struct{value bool}
    // var found bool
    c.Bind(&body)
    name := c.Params.ByName("flagid")
    if !flags.Exists(name) {
        c.Status(http.StatusNotFound)
    }
}

func CreateFlag(c *gin.Context) {
    var flag flags.Flag
    jsonErr := c.BindJSON(&flag)
    if jsonErr != nil {
        log.Println("Error!", jsonErr)
        return
    }
    flagErr := flags.Create(flag)
    if flagErr != nil {
        msg := fmt.Sprintf("Error - %s (%s)", flagErr.Error(), flag.Name)
        c.String(http.StatusConflict, msg)
        return
    }
    c.Status(http.StatusCreated)
}
