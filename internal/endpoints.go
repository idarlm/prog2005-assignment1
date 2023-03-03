package endpoints

import (
	endpoint "assignment1/internal/endpoint"
	"fmt"
	"net/http"
)

func ListenCoffee() {
	path := COFFEE_PATH
	http.HandleFunc(path, endpoint.HandlerCoffee)
	printEndpoint(path)
}

func ListenDiag() {
	path := ROOT_PATH_V1 + DIAG_PATH
	http.HandleFunc(path, endpoint.HandlerDiag)
	printEndpoint(path)
}

func printEndpoint(path string) {
	fmt.Println("Active endpoint:", path)
}
