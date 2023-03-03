package endpoint

import (
	"fmt"
	"net/http"
)

func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	fmt.Println("diag") // TODO: everything
}
