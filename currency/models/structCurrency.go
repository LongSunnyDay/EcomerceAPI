package models

type Currency struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
	Sign string `json:"sign,omitempty"`
	DefaultCurrency bool `json:"defaultCurrency,omitempty"`
}