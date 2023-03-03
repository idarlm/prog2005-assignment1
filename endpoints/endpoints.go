package endpoints

import (
	"assignment1/endpoints/endpoint"
	"fmt"
	"net/http"
)

func ListenCoffee() {
	http.HandleFunc(ROOT_PATH_V1+COFFEE_PATH, endpoint.HandlerCoffee)
	printListener(ROOT_PATH_V1 + COFFEE_PATH)
}

func printListener(path string) {
	fmt.Println("Listening on endpoint:", path)
}
