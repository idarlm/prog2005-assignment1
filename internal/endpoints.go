package endpoints

import (
	"assignment1/internal/endpoint"
	"assignment1/internal/stub"
	"fmt"
	"net/http"
)

// Assign handler func to path
func HandlePath(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, handler)
	fmt.Println("Set endpoint:", path)
}

// Only allow GET METHOD for handler
func HandleGET(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func SetAllEndpoints() {
	HandlePath(COFFEE_PATH, endpoint.CoffeeHandler)
	HandlePath(ROOT_PATH_V1+DIAG_PATH, HandleGET(endpoint.DiagHandler))
}

func SetStubEndpoints() {
	HandlePath("/stub/hipo", HandleGET(stub.HipoHandler))
}
