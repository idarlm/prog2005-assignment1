package restclient

import (
	"encoding/json"
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

func (client RestClient) AddQuery(key string, value string) {
	query := client.request.URL.Query()
	query.Add(key, value)
	client.request.URL.RawQuery = query.Encode()
}

func (client RestClient) GetContent(output any) {
	// instantiate client
	c := &http.Client{}
	defer c.CloseIdleConnections()

	// issue request
	res, err := c.Do(client.request)
	if err != nil {
		log.Fatal("Error in response:", err.Error())
	}

	decoder := json.NewDecoder(res.Body)
	e := decoder.Decode(output)
	if e != nil {
		log.Fatal(e)
	}
}
