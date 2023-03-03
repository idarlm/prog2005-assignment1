package main

import (
	endpoints "assignment1/internal"
	"fmt"
	"net/http"
)

func main() {
	//set all endpoints
	endpoints.HandleDiag()
	endpoints.HandleCoffee()

	//start server
	fmt.Println("Starting server on port:", endpoints.DEFAULT_PORT)

	err := http.ListenAndServe(":"+endpoints.DEFAULT_PORT, nil)
	if err != nil {
		fmt.Println("Error when starting server:", err.Error())
	}
}
