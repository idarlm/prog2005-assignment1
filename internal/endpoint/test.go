package endpoint

import (
	"assignment1/pkg/universities"
	"fmt"
	"net/http"
)

func TestUniHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Searching for Norwegian universities with 'technology' as name component...")
	uc := universities.NewUniClient()
	uc.AddQuery("name", "technology")
	uc.AddQuery("country", "norway")
	uc.Search()

	content := uc.Content()
	fmt.Println("Number of entries:", len(content))
	fmt.Println("Retrieved data: ", content)
}
