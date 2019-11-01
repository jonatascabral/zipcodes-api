package services

import "net/http"

func StartServer(port string) error {
	return http.ListenAndServe(port, nil)
}
