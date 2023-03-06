package neighbour

import (
	"assignment1/pkg/countries"
	"assignment1/pkg/universities"
	"fmt"
	"net/http"
)

func NeighbourHandler(w http.ResponseWriter, r *http.Request) {
	// check if request brokey

	//getNeighbours()

	//getCountries()

	//getUniversities()

	// compile data into response struct
	// (return specified amount of responses)
}

func getNeighbours(w *http.ResponseWriter, countryName string) ([]string, error) {
	// search for country name

	// return neighbourlist
	return nil, fmt.Errorf("neighbour: not implemented")
}

func getCountries(w *http.ResponseWriter, codes []string) ([]countries.BasicInfo, error) {
	// search for all neighbour countries by code

	// return countries
	return nil, fmt.Errorf("neighbour: not implemented")
}

func getUniversities(w *http.ResponseWriter, names []string) ([]universities.UniversityInfo, error) {
	// search for unis with uniname component in given countries
	// (seperate request for each country name)

	//return universities
	return nil, fmt.Errorf("neighbour: not implemented")
}
