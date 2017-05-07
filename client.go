package openapi

import "fmt"

const BaseUri = "https://api.openapi.ro/api/"

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
