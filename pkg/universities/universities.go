package universities

import "assignment1/pkg/restclient"

const API_URL = "http://universities.hipolabs.com"
const API_SEARCH_URL = API_URL + "/search"
const STUB_URL = "http://localhost:8080/stub/hipo"

// default data structure for api response
type UniversityInfo struct {
	AlphaTwoCode  string   `json:"alpha_two_code"`
	StateProvince string   `json:"state-province"`
	Domains       []string `json:"domains"`
	Country       string   `json:"country"`
	WebPages      []string `json:"web_pages"`
	Name          string   `json:"name"`
}

type UniClient struct {
	client  restclient.RestClient
	content []UniversityInfo
}

func NewUniClient() UniClient {
	return UniClient{restclient.NewRestClient(API_SEARCH_URL), nil}
}

// return a copy of content
func (uc *UniClient) Content() []UniversityInfo {
	return uc.content
}

// append key/value pair to query
func (uc *UniClient) AddQuery(key, value string) {
	uc.client.AddQuery(key, value)
}

// set query to key/value pair
func (uc *UniClient) SetQuery(key, value string) {
	uc.client.SetQuery(key, value)
}

// perform request and parse response data
func (uc *UniClient) Search() {
	uc.client.GetContent(&uc.content)
}

// set name query and perform request
func (uc *UniClient) SearchByName(value string) {
	uc.client.SetQuery("name", value)
	uc.Search()
}

func (uc *UniClient) SearchByCountry(value string) {
	uc.client.SetQuery("country", value)
	uc.Search()
}
