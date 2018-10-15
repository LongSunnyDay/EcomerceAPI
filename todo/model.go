package todo

type Todo struct {
	Id           int    `json:"Id,sting"`
	Title        string `json:"Title"`
	Category     string `json:"Category"`
	//Created   string `json:"Created"`
	//Completed string `json:"Completed"`
	State        string `json:"State"`
}
