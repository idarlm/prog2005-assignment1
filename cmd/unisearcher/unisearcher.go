package main

import (
	endpoints "assignment1/internal"
	"fmt"
	"net/http"
	"os"
)

func main() {
	endpoints.SetApiEndpoints()
	endpoints.SetDebugEndpoints()

	port := os.Getenv("$PORT")
	if port == "" {
		port = endpoints.DEFAULT_PORT
	}

	fmt.Println("Listening on port:", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error when starting server:", err.Error())
	}
}
