package uniinfo

import (
	"assignment1/pkg/countries"
	"assignment1/pkg/universities"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Default response type
type UniinfoDefault struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
}

func UniinfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	path := strings.Split(r.URL.EscapedPath(), "/")

	// check if path is formatted correctly
	if len(path) != 5 && !(len(path) == 6 && path[5] == "") {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	name := path[4]

	//check if name is empty
	if name == "" {
		http.Error(w, "Bad request.", http.StatusBadRequest)
		return
	}

	// format name before performing search
	name = strings.ReplaceAll(name, "%20", " ")
	name = strings.ReplaceAll(name, "+", " ")
	fmt.Println("Uniinfo: handling request for", name)

	// search for university name
	info, err := search(name, &w)
	if err != nil {
		fmt.Println("Error in UniinfoHandler:", err.Error())
		return
	}

	// encode data to json format
	encoder := json.NewEncoder(w)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("Error in UniinfoHandler:", err.Error())
		http.Error(w, "Internal error.", http.StatusInternalServerError)
		return
	}
}

// Search for universities and compose response info
func search(name string, w *http.ResponseWriter) (resp []UniinfoDefault, err error) {
	// search for universities
	uc := universities.NewUniClient()
	err = uc.SearchByName(name)
	if err != nil {
		http.Error(*w, "Internal error.", http.StatusInternalServerError)
		return nil, err
	}
	unis := uc.Content()

	// check if content was found
	if len(unis) == 0 {
		http.Error(*w, "Not found.", http.StatusNotFound)
		return nil, fmt.Errorf("uniinfo: No results")
	}

	// store all country codes in array
	var codes []string
	for _, u := range unis {
		codes = append(codes, u.AlphaTwoCode)
	}

	// search for country info
	cc := countries.NewClient()
	err = cc.SearchByAlpha(codes)
	if err != nil {
		http.Error(*w, "Internal error.", http.StatusInternalServerError)
		return nil, err
	}
	counts := cc.Basic()

	// compile response data
	resp = make([]UniinfoDefault, len(unis))
	for i := range resp {
		resp[i].Name = unis[i].Name
		resp[i].Country = unis[i].Country
		resp[i].Isocode = unis[i].AlphaTwoCode
		resp[i].Webpages = unis[i].WebPages

		c, e := findCountry(resp[i].Isocode, counts)
		if e != nil {
			http.Error(*w, "Internal error.", http.StatusInternalServerError)
			return nil, e
		}

		resp[i].Languages = c.Languages
		resp[i].Map = c.Maps.OpenStreetMaps
	}

	return resp, err
}

func findCountry(code string, collection []countries.BasicInfo) (countries.BasicInfo, error) {
	for _, c := range collection {
		found := c.Cca2 == code
		if found {
			return c, nil
		}
	}

	return countries.BasicInfo{}, fmt.Errorf("uniinfo: could not find country with given code")
}
