package stub

import (
	"fmt"
	"net/http"
	"os"
)

func ParseFile(name string) []byte {
	file, e := os.ReadFile(name)
	if e != nil {
		fmt.Println("File error:", e)
		os.Exit(1)
	}
	return file
}

func HipoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	content := ParseFile("./internal/stub/hipostub.json")
	fmt.Fprint(w, string(content))
}

func CountryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	content := ParseFile("./internal/stub/countrystub.json")
	fmt.Fprint(w, string(content))
}
