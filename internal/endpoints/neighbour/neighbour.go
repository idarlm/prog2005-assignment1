package neighbour

import (
	"assignment1/internal/endpoints/common"
	"assignment1/internal/endpoints/uniinfo"
	"assignment1/pkg/countries"
	"assignment1/pkg/universities"
	"fmt"
	"net/http"
)

func NeighbourHandler(w http.ResponseWriter, r *http.Request) {
	// check if request brokey
	country := "Norway"
	nameComponent := "tech"

	// query for name and return array of neighbours
	codes, err := getNeighbours(&w, country)
	if err != nil {
		fmt.Println("neighbour: error on getNeighbours: ", err.Error())
		return
	}

	// get info about all neighbours
	countries, err := getCountries(&w, codes)
	if err != nil {
		fmt.Println("neighbour: error on getCountries: ", err.Error())
		return
	}

	// get all universities with uniComponent
	unis, err := getUniversities(&w, nameComponent, countries)
	if err != nil {
		fmt.Println("neighbour: error on getUniversities: ", err.Error())
	}

	// compile all data into array of UniinfoDefault
	comp, err := uniinfo.CompileData(&w, unis, countries)
	if err != nil {
		fmt.Println("neighbour: error on CompileData: ", err.Error())
		common.ErrorInternalError(&w)
		return
	}

	// array to json
	err = common.EncodeJson(&w, comp)
	if err != nil {
		fmt.Println("neighbour: error on EncodeJson: ", err.Error())
		common.ErrorInternalError(&w)
		return
	}
	common.ContentTypeJson(&w)
}

func getNeighbours(w *http.ResponseWriter, countryName string) ([]string, error) {
	// search for country name
	cc := countries.NewClient()
	c := cc.Client()
	c.AddQuery("fulltext", "true")

	err := cc.SearchByName(countryName)
	if err != nil {
		common.ErrorInternalError(w)
		return nil, err
	}

	// check if result is empty
	countries := cc.Basic()
	if len(countries) == 0 {
		common.ErrorNotFound(w)
		return nil, fmt.Errorf("getNeighbours: no results")
	}

	codes := countries[0].Borders

	// return neighbourlist
	return codes, err
}

func getCountries(w *http.ResponseWriter, codes []string) ([]countries.BasicInfo, error) {
	// search for all neighbour countries by code
	cc := countries.NewClient()
	err := cc.SearchByAlpha(codes)
	if err != nil {
		common.ErrorInternalError(w)
		return nil, err
	}

	// check if results is empty
	result := cc.Basic()
	if len(result) == 0 {
		common.ErrorNotFound(w)
		return nil, fmt.Errorf("getCountries: no results")
	}

	// return countries
	return result, err
}

func getUniversities(w *http.ResponseWriter, nameComponent string, countryList []countries.BasicInfo) ([]universities.UniversityInfo, error) {
	// search for unis with uniname component in given countries
	// (seperate request for each country name)
	var unis []universities.UniversityInfo
	uc := universities.NewUniClient()

	// append universities to unis
	for _, c := range countryList {
		// search for universities with nameComponent in country c
		err := uc.SearchNameCountry(nameComponent, c.Name.Common)
		if err != nil {
			common.ErrorInternalError(w)
			return nil, err
		}

		// append
		unis = append(unis, uc.Content()...)
	}

	//return universities
	return unis, nil
}
