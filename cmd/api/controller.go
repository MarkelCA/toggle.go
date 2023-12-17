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
    result,_ := fc.repository.List()
    c.JSON(http.StatusOK, result)
}

func (fc FlagController) FindFlag(c *gin.Context) {
    key := c.Params.ByName("flagid")
    value,err := fc.repository.Get(key)
    if err == flags.FlagNotFoundError {
        c.JSON(http.StatusNotFound, nil)
        return
    }

    c.JSON(http.StatusOK, value)
}

func (fc FlagController) UpdateFlag(c *gin.Context) {
    name := c.Params.ByName("flagid")
    var body struct{Value bool `json:"value"`}
    c.BindJSON(&body)

    log.Println(body,name)

    result,_ := fc.repository.Exists(name)
    if !result {
        c.Status(http.StatusNotFound)
        c.JSON(http.StatusNotFound, "Error - Flag not found.")
        return
    }

    fc.repository.Update(name, body.Value)
    c.Status(http.StatusCreated)
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
