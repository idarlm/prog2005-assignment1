package uniinfo

import (
	"assignment1/internal/endpoints/common"
	"assignment1/pkg/countries"
	"assignment1/pkg/universities"
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
	common.ContentTypeJson(&w)

	path := strings.Split(r.URL.EscapedPath(), "/")

	// check if path is formatted correctly
	if len(path) != 5 && !(len(path) == 6 && path[5] == "") {
		common.ErrorBadRequest(&w)
		return
	}

	name := path[4]

	//check if name is empty
	if name == "" {
		common.ErrorBadRequest(&w)
		return
	}

	// format name before performing search
	name = common.FormatString(name)
	fmt.Println("uniinfo: handling request for", name)

	// search for university name
	info, err := search(name, &w)
	if err != nil {
		fmt.Println("Error in UniinfoHandler:", err.Error())
		return
	}

	// encode data to json format
	err = common.EncodeJson(&w, info)
	if err != nil {
		fmt.Println("error in uniinfo:", err.Error())
	}
}

// Search for universities and compose response info
func search(name string, w *http.ResponseWriter) (resp []UniinfoDefault, err error) {
	// search for universities
	uc := universities.NewUniClient()
	err = uc.SearchByName(name)
	if err != nil {
		common.ErrorInternalError(w)
		return nil, err
	}
	unis := uc.Content()

	// check if content was found
	if len(unis) == 0 {
		common.ErrorNotFound(w)
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
		common.ErrorInternalError(w)
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
			common.ErrorInternalError(w)
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
