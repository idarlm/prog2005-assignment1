package endpoints

import (
	endpoint "assignment1/internal/endpoint"
	"fmt"
	"net/http"
)

func ListenCoffee() {
	http.HandleFunc(ROOT_PATH_V1+COFFEE_PATH, endpoint.HandlerCoffee)
	printListener(ROOT_PATH_V1 + COFFEE_PATH)
}

func ListenDiag() {
	http.HandleFunc(ROOT_PATH_V1+DIAG_PATH, endpoint.HandlerDiag)
	printListener(ROOT_PATH_V1 + DIAG_PATH)
}

func printListener(path string) {
	fmt.Println("Set endpoint:", path)
}
