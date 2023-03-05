package countries

import "assignment1/pkg/restclient"

type CountryClient struct {
	client   restclient.RestClient
	response BasicInfo
}

func (cc *CountryClient) Response() BasicInfo {
	return cc.response
}

func (cc *CountryClient) Client() restclient.RestClient {
	return cc.client
}

func (cc *CountryClient) SearchByName(name string) {

}
