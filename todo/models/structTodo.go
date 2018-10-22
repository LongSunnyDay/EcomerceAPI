package models

type Todo struct {
	ID       int    `json:"ID,omitempty"`
	Title    string `json:"Title,omitempty"`
	Category string `json:"Category,omitempty"`
	Content  string `json:"Content,omitempty"`
	Created  string `json:"Creation_time,omitempty"`
	Modified string `json:"Modification_time,omitempty"`
	State    string `json:"State,omitempty"`
}
