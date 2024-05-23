package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/markelca/toggles/pkg/flags"
	"log"
	"net/http"
)

type FlagController struct {
	service flags.FlagService
}

func NewFlagController(r flags.FlagService) FlagController {
	return FlagController{r}
}

func (fc FlagController) Init(host string) {
	r := gin.Default()
	r.GET("/flags", fc.ListFlags)
	r.GET("/flags/:flagid", fc.GetFlag)
	r.PUT("/flags/:flagid", fc.UpdateFlag)
	r.POST("/flags", fc.CreateFlag)
	r.DELETE("/flags/:flagid", fc.DeleteFlag)
	r.Run(host)
}

func (fc FlagController) ListFlags(c *gin.Context) {
	result, err := fc.service.List()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (fc FlagController) GetFlag(c *gin.Context) {
	key := c.Params.ByName("flagid")
	value, err := fc.service.Get(key)

	if err == flags.ErrFlagNotFound {
		c.JSON(http.StatusNotFound, nil)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, value)
}

func (fc FlagController) DeleteFlag(c *gin.Context) {
	name := c.Params.ByName("flagid")
	if err := fc.service.Delete(name); err != nil {
		if err == flags.ErrFlagNotFound {
			c.Status(http.StatusNotFound)
			return
		} else {
			c.Status(http.StatusInternalServerError)
			return
		}
	}
}

func (fc FlagController) UpdateFlag(c *gin.Context) {
	name := c.Params.ByName("flagid")
	var body struct {
		Value bool `json:"value"`
	}
	c.BindJSON(&body)

	result, err := fc.service.Exists(name)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if !result {
		c.Status(http.StatusNotFound)
		c.JSON(http.StatusNotFound, "Error - Flag not found.")
		return
	}

	fc.service.Update(name, body.Value)
	c.Status(http.StatusCreated)
}

func (fc FlagController) CreateFlag(c *gin.Context) {
	var flag flags.Flag
	jsonErr := c.BindJSON(&flag)
	if jsonErr != nil {
		log.Println("Error!", jsonErr)
		return
	}
	flagErr := fc.service.Create(flag)

	if flagErr != nil {
		msg := fmt.Sprintf("Error - %s (%s)", flagErr.Error(), flag.Name)
		c.String(http.StatusConflict, msg)
		return
	}
	c.Status(http.StatusCreated)
}
