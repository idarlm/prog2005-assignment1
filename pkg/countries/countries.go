package countries

import (
	"assignment1/pkg/restclient"
)

type CountryClient struct {
	client restclient.RestClient
	basic  []BasicInfo
}

func NewClient() CountryClient {
	return CountryClient{restclient.NewRestClient(API_BASE_V3_1), nil}
}

func (cc *CountryClient) Basic() []BasicInfo {
	return cc.basic
}

func (cc *CountryClient) Client() restclient.RestClient {
	return cc.client
}

func (cc *CountryClient) SearchBasic() error {
	cc.client.AddQuery("fields", "name")
	cc.client.AddQuery("fields", "capital")
	cc.client.AddQuery("fields", "borders")
	cc.client.AddQuery("fields", "maps")
	cc.client.AddQuery("fields", "languages")
	cc.client.AddQuery("fields", "cca2")

	err := cc.client.GetContent(&cc.basic)
	if err != nil {
		cc.basic = nil
	}
	return err
}

func (cc *CountryClient) SearchByName(value string) error {
	cc.client.SetPath(ENDPOINT_NAME + value)
	err := cc.SearchBasic()
	return err
}

func (cc *CountryClient) SearchByAlpha(value []string) error {
	cc.client.SetPath(ENDPOINT_ALPHA)
	for _, v := range value {
		cc.client.AddQuery("codes", v)
	}
	err := cc.SearchBasic()
	return err
}
