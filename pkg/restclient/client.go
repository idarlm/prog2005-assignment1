package restclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RestClient struct {
	request *http.Request
}

func NewRestClient(url string) RestClient {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Error in request:", err.Error())
	}

	return RestClient{r}
}

func (client *RestClient) AddQuery(key string, value string) {
	query := client.request.URL.Query()
	query.Add(key, value)
	client.request.URL.RawQuery = query.Encode()
}

func (client *RestClient) SetQuery(key string, value string) {
	query := client.request.URL.Query()
	query.Set(key, value)
	client.request.URL.RawQuery = query.Encode()
}

func (client *RestClient) ClearQuery() {
	query := client.request.URL.Query()
	for k := range query {
		delete(query, k)
	}
}

func (client *RestClient) GetContent(output any) error {
	// instantiate client
	c := &http.Client{}
	defer c.CloseIdleConnections()

	fmt.Println("Performing query: ", client.request.URL.String())

	// issue request
	res, err := c.Do(client.request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// fail if not ok
	if res.StatusCode != http.StatusOK {
		output = nil
		return fmt.Errorf("restclient: did not recieve status code 200. output is set to nil")
	}

	// decode json data
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(output)

	return err
}
