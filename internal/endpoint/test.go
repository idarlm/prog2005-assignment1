package endpoint

import (
	"assignment1/pkg/countries"
	"fmt"
	"net/http"
)

func TestUniHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Parsing country info...")
	cc := countries.NewClient()
	err := cc.SearchByName("norw")
	if err != nil {
		fmt.Println("test: error when searching by name:", err)
		return
	}

	for _, res := range cc.Basic() {
		fmt.Printf("\nName: %s\nNative name(s): %v\nLanguages: %v\nCapital: %v\nBorders: %v\nOpenStreetMaps: %s\ncca2: %s\n",
			res.Name.Common,
			res.Name.NativeNames,
			res.Languages,
			res.Capital,
			res.Borders,
			res.Maps.OpenStreetMaps,
			res.Cca2)
	}
}
