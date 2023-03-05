package endpoint

import (
	"assignment1/pkg/universities"
	"fmt"
	"net/http"
)

func TestUniHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Searching for universities with 'norwegian' as name component...")
	data := universities.SearchByName("norwegian")
	fmt.Println("Retrieved data: ", data)
}
