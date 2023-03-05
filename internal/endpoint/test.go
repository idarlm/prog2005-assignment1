package endpoint

import (
	"assignment1/pkg/countries"
	"fmt"
	"net/http"
)

func TestUniHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parsing country info...")
	cc := countries.NewClient()
	err := cc.SearchByName("norway")
	if err != nil {
		fmt.Println("test: error when searching by name:", err)
		return
	}

	res := cc.Basic()[0]
	fmt.Printf("\nName: %s\nNative name(s): %v\nCapital: %v\nBorders: %v\nOpenStreetMaps: %s\n",
		res.Name.Common,
		res.Name.NativeNames,
		res.Capital,
		res.Borders,
		res.Maps.OpenStreetMaps)
}
