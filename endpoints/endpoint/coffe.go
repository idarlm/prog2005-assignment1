package endpoint

import (
	"fmt"
	"net/http"
)

func HandlerCoffee(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, "wants coffee, but i'm a teapot!")
	http.Error(w, "I'm a teapot!", http.StatusTeapot)
}
