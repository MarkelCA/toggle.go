package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markelca/toggle.go/flags"
)


type FlagController struct {
    repository flags.FlagRepository
}

func NewFlagController(r flags.FlagRepository) FlagController {
    return FlagController{r}
}

func (fc FlagController) ListFlags(c *gin.Context) {
    c.JSON(http.StatusOK, fc.repository.List())
}

func (fc FlagController) FindFlag(c *gin.Context) {
    for _,flag := range fc.repository.List() {
        if flag.Name == c.Params.ByName("flagid") {
            c.JSON(http.StatusOK, flag.Value)
            return
        }
    }
    c.Status(http.StatusNotFound)
}
func (fc FlagController) UpdateFlag(c *gin.Context) {
    var body struct{value bool}
    // var found bool
    c.Bind(&body)
    name := c.Params.ByName("flagid")
    if !fc.repository.Exists(name) {
        c.Status(http.StatusNotFound)
    }
}

func (fc FlagController) CreateFlag(c *gin.Context) {
    var flag flags.Flag
    jsonErr := c.BindJSON(&flag)
    if jsonErr != nil {
        log.Println("Error!", jsonErr)
        return
    }
    flagErr := fc.repository.Create(flag)
    if flagErr != nil {
        msg := fmt.Sprintf("Error - %s (%s)", flagErr.Error(), flag.Name)
        c.String(http.StatusConflict, msg)
        return
    }
    c.Status(http.StatusCreated)
}
