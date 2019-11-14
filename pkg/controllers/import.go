package controllers

import (
	"fmt"
	"github.com/jonatascabral/zipcodes-api/pkg/rabbitmq"
	"github.com/jonatascabral/zipcodes-api/pkg/services"
	"log"
	"net/http"
)

func ImportCsv(writer http.ResponseWriter, request *http.Request) {
	Csv, err := services.NewCsv(request.Body)
	defer request.Body.Close()
	if err != nil {
		log.Println("Error parsing CSV file", err)
		http.Error(writer, "Error parsing CSV file", http.StatusInternalServerError)
		return
	}

	for _, line := range Csv.Data {
		_, zipcode := line[0], line[1]
		_, err = rabbitmq.Publish("zipcodes", fmt.Sprintf("{\"Code\": \"%s\"}", zipcode))
		if err != nil {
			log.Println("Error parsing CSV file", err)
			http.Error(writer, "Error parsing CSV file", http.StatusInternalServerError)
			return
		}
	}
}

func Ping(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "It works")
}
