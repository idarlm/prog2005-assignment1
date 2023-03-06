// common.go
// contains functions which may be useful across several endpoints

package common

import (
	"fmt"
	"net/http"
)

// Writes bad request error
func ErrorBadRequest(w *http.ResponseWriter) {
	status := http.StatusBadRequest
	http.Error(*w, fmt.Sprint(status, "Bad request"), status)
}

func ErrorInternalError(w *http.ResponseWriter) {
	status := http.StatusInternalServerError
	http.Error(*w, fmt.Sprint(status, "Bad request"), status)
}

func ErrorNotFound(w *http.ResponseWriter) {
	status := http.StatusNotFound
	http.Error(*w, fmt.Sprint(status, "Not found"), status)
}
