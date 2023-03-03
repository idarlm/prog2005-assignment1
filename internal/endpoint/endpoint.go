package endpoint

import (
	"net/http"
)

// Filter out all METHOD types other than GET.
func HandleGET(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Only GET method is allowed.", http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}
