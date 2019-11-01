package routes

import (
	"github.com/jonatascabral/zipcodes-api/pkg/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Ping)
	http.HandleFunc("/import/csv", controllers.ImportCsv)
}