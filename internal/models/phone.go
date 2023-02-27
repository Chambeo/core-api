package models

type Phone struct {
	PhoneCode   string `json:"phone_code,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
