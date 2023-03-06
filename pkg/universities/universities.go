package universities

import (
	"assignment1/pkg/restclient"
)

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

// prod universities api
func (uc *UniClient) Prod() string {
	rc := restclient.NewRestClient(API_URL)
	status, err := rc.Prod()
	if err != nil {
		status = "Service unavailable"
	}

	return status
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
func (uc *UniClient) Search() error {
	err := uc.client.GetContent(&uc.content)
	if err != nil {
		uc.content = nil
	}
	return err
}

// set name query and perform request
func (uc *UniClient) SearchByName(value string) error {
	uc.client.SetQuery("name", value)
	err := uc.Search()
	return err
}

func (uc *UniClient) SearchByCountry(value string) error {
	uc.client.SetQuery("country", value)
	err := uc.Search()
	return err
}

func (uc *UniClient) SearchNameCountry(name, country string) error {
	uc.client.SetQuery("name", name)
	err := uc.SearchByCountry(country)
	return err
}
