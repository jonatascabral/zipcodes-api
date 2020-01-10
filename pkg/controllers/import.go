package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/jonatascabral/zipcodes-api/pkg/services"
	"net/http"
)

func ImportCsv(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
	}

	file, err := formFile.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
	}

	Csv, err := services.NewCsv(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
	}

	for _, line := range Csv.Data {
		_, zipcode := line[0], line[1]
		_, err = rabbitmq.Publish("zipcodes", fmt.Sprintf("{\"Code\": \"%s\"}", zipcode))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
		}
	}

	c.String(http.StatusNoContent, "")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}
