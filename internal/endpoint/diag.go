package endpoint

import (
	"fmt"
	"net/http"
)

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("diag") // TODO: everything
}
