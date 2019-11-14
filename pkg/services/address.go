package services

import (
	"encoding/json"
	"fmt"
	"github.com/jonatascabral/zipcodes-api/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	url = "https://viacep.com.br/ws/%s/json/"
)

type Response struct {
	Code string `json:"cep"`
	Address string `json:"logradouro"`
	Complement string `json:"complemento"`
	District string `json:"bairro"`
	City string `json:"localidade"`
	State string `json:"uf"`
}

func GetAddress (address *models.Address) (*models.Address, error) {
	url := parseUrl(address.Code)
	log.Println(fmt.Sprintf("Calling url %s", url))
	response, err := http.Get(url)

	if err != nil {
		log.Println("Error getting address: http.Get failed", err)
		return nil, err
	}
	defer response.Body.Close()
	jsonResponse, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println("Error getting address: read response body", err)
		return nil, err
	}

	var wsResponse *Response
	err = json.Unmarshal([]byte(jsonResponse), &wsResponse)

	if err != nil {
		log.Println("Error getting address: json parse", err)
		return nil, err
	}

	address.Address = wsResponse.Address
	address.City = wsResponse.City
	address.State = wsResponse.State
	address.District = wsResponse.District
	address.Country = "Brazil"

	return address, nil
}

func parseUrl(postalCode string) string {
	return fmt.Sprintf(url, postalCode)
}