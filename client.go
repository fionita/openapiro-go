package openapi

import (
	"fmt"
	"net/http"
)

const baseUri = "https://api.openapi.ro/api/"

type Config struct {
	Token string
}

type ApiClient struct {
	conf *Config
}

type Client interface {
	Company(cif string) (*CompanyResponse, error)
}

func Init(conf *Config) (*ApiClient, error) {
	if conf.Token == "" {
		return nil, fmt.Errorf("%v", "No apiKey provided!")
	}

	return &ApiClient{conf}, nil
}

func (c *ApiClient) apiCall(endpiont, query string) (*http.Response, error) {
	url := fmt.Sprintf(baseUri + endpiont + "/" + query)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	req.Header.Add("x-api-key", c.conf.Token)

	httpClient := http.Client{}

	resp, err := httpClient.Do(req)

	return resp, err
}
