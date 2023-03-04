package restclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Test() {
	url := "http://localhost:8080/stub/hipo"

	// create request
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error when creating request:", err.Error())
		return
	}

	// set content type
	r.Header.Add("content-type", "application/json")

	// instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// issue request
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("Error in response:", err.Error())
		return
	}

	fmt.Println("Status:", res.Status)
}

type RestClient struct {
	URL string
}

func (client RestClient) GetContent(output any) {
	// create request
	r, err := http.NewRequest(http.MethodGet, client.URL, nil)
	if err != nil {
		log.Fatal("Error in request:", err.Error())
	}

	// set content type
	r.Header.Add("content-type", "application/json")

	// instantiate client
	c := &http.Client{}
	defer c.CloseIdleConnections()

	// issue request
	res, err := c.Do(r)
	if err != nil {
		log.Fatal("Error in response:", err.Error())
	}

	decoder := json.NewDecoder(res.Body)
	e := decoder.Decode(output)
	if e != nil {
		log.Fatal(e)
	}
}
