package main

import (
	endpoints "assignment1/internal"
	"fmt"
	"net/http"
)

func main() {
	//set all endpoints to listen to
	endpoints.ListenCoffee()
	endpoints.ListenDiag()

	//start server
	err := http.ListenAndServe(":"+endpoints.DEFAULT_PORT, nil)
	if err != nil {
		fmt.Println("Server running and listening on port:", endpoints.DEFAULT_PORT)
	}
}
