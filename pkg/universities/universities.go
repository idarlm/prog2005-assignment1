package universities

import "assignment1/pkg/restclient"

const API_URL = "http://universities.hipolabs.com"
const API_SEARCH_URL = API_URL + "/search"

type University struct {
	StateProvince string   `json:"state-province"`
	Domains       []string `json:"domains"`
	Country       string   `json:"country"`
	WebPages      []string `json:"web_pages"`
	Name          string   `json:"name"`
	AlphaTwoCode  string   `json:"alpha_two_code"`
}

func SearchByName(name string) (result []University) {
	client := restclient.NewRestClient(API_SEARCH_URL)
	client.AddQuery("name", name)
	client.GetContent(&result)
	return
}
