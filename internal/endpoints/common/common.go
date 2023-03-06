// common.go
// contains functions which may be useful across several endpoints

package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Writes bad request error
func ErrorBadRequest(w *http.ResponseWriter) {
	status := http.StatusBadRequest
	http.Error(*w, fmt.Sprint(status, " Bad request"), status)
}

func ErrorInternalError(w *http.ResponseWriter) {
	status := http.StatusInternalServerError
	http.Error(*w, fmt.Sprint(status, " Internal server error"), status)
}

func ErrorNotFound(w *http.ResponseWriter) {
	status := http.StatusNotFound
	http.Error(*w, fmt.Sprint(status, " Not found"), status)
}

func ContentTypeJson(w *http.ResponseWriter) {
	(*w).Header().Add("content-type", "application/json")
}

func FormatString(str string) string {
	result := strings.ReplaceAll(str, "%20", " ")
	result = strings.ReplaceAll(result, "+", " ")

	return result
}

func EncodeJson(w *http.ResponseWriter, v any) error {
	// encode data to json format
	encoder := json.NewEncoder(*w)
	err := encoder.Encode(v)
	if err != nil {
		fmt.Println("Error in UniinfoHandler:", err.Error())
		ErrorInternalError(w)
	}

	return err
}
