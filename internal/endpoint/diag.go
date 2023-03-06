package endpoint

import (
	"assignment1/pkg/countries"
	"assignment1/pkg/universities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var Timestamp time.Time

type DiagResp struct {
	UniversitiesApi string `json:"universitiesapi"`
	CountriesApi    string `json:"countriesapi"`
	Version         string `json:"version"`
	Uptime          string `json:"uptime"`
}

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	res := DiagResp{}

	// prod universities api
	uc := universities.NewUniClient()
	res.UniversitiesApi = uc.Prod()

	// prod countries api
	cc := countries.NewClient()
	res.CountriesApi = cc.Prod()

	res.Version = "v1" // hardcoded lmao
	res.Uptime = fmt.Sprintf("%ds", int(time.Since(Timestamp).Seconds()))

	//encode result
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	if err != nil {
		log.Println("DiagHandler: error on encode json:", err.Error())
		http.Error(w, "Internal error.", http.StatusInternalServerError)
	}

	w.Header().Add("content-type", "application/json")
}
