package openapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CompanyResponse struct {
	UltimaPrelucrare string `json:"ultima_prelucrare"`
	UltimaDeclaratie string `json:"ultima_declaratie"`
	TvaLaIncasare    []struct {
		TipDescriere string      `json:"tip_descriere"`
		Tip          string      `json:"tip"`
		Publicare    string      `json:"publicare"`
		PanaLa       interface{} `json:"pana_la"`
		DeLa         string      `json:"de_la"`
		Actualizare  string      `json:"actualizare"`
	} `json:"tva_la_incasare"`
	Tva         string `json:"tva"`
	Telefon     string `json:"telefon"`
	Stare       string `json:"stare"`
	Radiata     bool   `json:"radiata"`
	NumarRegCom string `json:"numar_reg_com"`
	Meta        struct {
		UpdatedAt     string `json:"updated_at"`
		LastChangedAt string `json:"last_changed_at"`
	} `json:"meta"`
	Judet         string      `json:"judet"`
	ImpozitProfit interface{} `json:"impozit_profit"`
	ImpozitMicro  string      `json:"impozit_micro"`
	Fax           interface{} `json:"fax"`
	Denumire      string      `json:"denumire"`
	CodPostal     string      `json:"cod_postal"`
	Cif           string      `json:"cif"`
	Adresa        string      `json:"adresa"`
	ActAutorizare interface{} `json:"act_autorizare"`
	Accize        interface{} `json:"accize"`
}

func (c *ApiClient) Companies(cif string) (*CompanyResponse, error) {
	resp, err := c.apiCall("companies", cif)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	var response *CompanyResponse
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
