package endpoint

import (
	"fmt"
	"net/http"
)

func HandlerCoffee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone wants to brew coffee, but i'm a teapot!")
	http.Error(w, "I'm a teapot!", http.StatusTeapot)
}
