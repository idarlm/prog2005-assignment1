package main

import (
	"assignment1/endpoints"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")
	endpoints.ListenCoffee()

	http.ListenAndServe(":"+endpoints.DEFAULT_PORT, nil)
}
