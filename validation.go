package openapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Valid bool   `json:"valid"`
	CIF   string `json:"cif,omitempty"`
	CNP   string `json:"cnp,omitempty"`
	IBAN  string `json:"iban,omitempty"`
}

func (c *ApiClient) ValidateCIF(cif string) (*Response, error) {
	resp, err := validate(c, "cif", cif)

	return resp, err
}

func (c *ApiClient) ValidateCNP(cnp string) (*Response, error) {
	resp, err := validate(c, "cnp", cnp)

	return resp, err
}

func (c *ApiClient) ValidateIBAN(iban string) (*Response, error) {
	resp, err := validate(c, "iban", iban)

	return resp, err
}

func validate(c *ApiClient, endpoint string, query string) (*Response, error) {
	resp, err := c.apiCall("validate/"+endpoint, query)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	var response *Response
	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, fmt.Errorf("%v", err)
		}

		return response, nil
	default:
		return nil, fmt.Errorf("Openapi error - status code: %v", resp.StatusCode)
	}
}
