package models

type Address struct {
	Street   string `json:"address_street,omitempty"`
	Number   string `json:"address_number,omitempty"`
	Details  string `json:"address_details"`
	Province string `json:"province"`
	City     string `json:"city,omitempty"`
	ZipCode  string `json:"zip_code,omitempty"`
}
