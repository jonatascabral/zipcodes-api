package controllers

import (
	"fmt"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/jonatascabral/zipcodes-api/pkg/services"
	"github.com/labstack/echo"
	"net/http"
)

func ImportCsv(c echo.Context) error {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
		return err
	}

	file, err := formFile.Open()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
		return err
	}

	Csv, err := services.NewCsv(file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
		return err
	}

	for _, line := range Csv.Data {
		_, zipcode := line[0], line[1]
		_, err = rabbitmq.Publish("zipcodes", fmt.Sprintf("{\"Code\": \"%s\"}", zipcode))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error parsing CSV formFile")
			return err
		}
	}

	return c.String(http.StatusNoContent, "")
}

func Ping(c echo.Context) error {
	return c.String(http.StatusOK, "It works")
}
