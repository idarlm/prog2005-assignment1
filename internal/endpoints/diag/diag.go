package diag

import (
	"assignment1/internal/endpoints/common"
	"assignment1/pkg/countries"
	"assignment1/pkg/universities"
	"fmt"
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
	err := common.EncodeJson(&w, res)
	if err != nil {
		fmt.Println("diag: error when encoding json:", err.Error())
		return
	}

	common.ContentTypeJson(&w)
}
