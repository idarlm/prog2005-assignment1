package endpoints

import (
	endpoint "assignment1/internal/endpoint"
	"fmt"
	"net/http"
)

func HandleCoffee() {
	path := COFFEE_PATH
	http.HandleFunc(path, endpoint.CoffeeHandler)
	printEndpoint(path)
}

func HandleDiag() {
	path := ROOT_PATH_V1 + DIAG_PATH
	http.HandleFunc(path, endpoint.DiagHandler)
	printEndpoint(path)
}

func printEndpoint(path string) {
	fmt.Println("Active endpoint:", path)
}
