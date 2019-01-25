package language

type Language struct {
	Id     int    `json:"id,omitempty"`
	Code   string `json:"code,omitempty"`
	Code3  string `json:"code3,omitempty"`
	Name   string `json:"name,omitempty"`
	NameEn string `json:"nameEn,omitempty"`
}
