package main

import (
	endpoints "assignment1/internal"
	"fmt"
	"net/http"
)

func main() {
	endpoints.SetEndpoints()

	//start server
	fmt.Println("Starting server on port:", endpoints.DEFAULT_PORT)

	err := http.ListenAndServe(":"+endpoints.DEFAULT_PORT, nil)
	if err != nil {
		fmt.Println("Error when starting server:", err.Error())
	}
}
