package endpoints

import (
	"assignment1/internal/endpoint"
	"fmt"
	"net/http"
)

func HandlePath(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, handler)
	fmt.Println("Set endpoint:" + path)
}

func HandleCoffee() {
	HandlePath(COFFEE_PATH, endpoint.HandleGET(endpoint.CoffeeHandler))
}

func HandleDiag() {
	HandlePath(ROOT_PATH_V1+DIAG_PATH, endpoint.HandleGET(endpoint.DiagHandler))
}
