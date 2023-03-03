package endpoints

import (
	endpoint "assignment1/internal/endpoint"
	"fmt"
	"net/http"
)

func ListenCoffee() {
	path := COFFEE_PATH
	http.HandleFunc(path, endpoint.HandlerCoffee)
	printListener(path)
}

func ListenDiag() {
	path := ROOT_PATH_V1 + DIAG_PATH
	http.HandleFunc(path, endpoint.HandlerDiag)
	printListener(path)
}

func printListener(path string) {
	fmt.Println("Set endpoint:", path)
}
