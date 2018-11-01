package user

import "net/http"

type User struct {
	ID       string   `json:"id,omitempty"`
	Customer Customer `json:"customer,omitempty"`
	Password string   `json:"password,omitempty"`
	GroupId int `json:"group_id:omitmepty"`
}

type Customer struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

type UpdatePassword struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}
type JwtAuthToken struct {
	Token string `json:"token,omitempty"`
}

type LoginForm struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type Response struct {
	Acknowledged bool          `json:"acknowledged,omitempty"`
	Code         int           `json:"code,omitempty"`
	CreatedAt    string        `json:"created_at,omitempty"`
	Payload      *http.Request `json:"payload,omitempty"`
	Result       interface{}   `json:"result,omitempty"`
	ResultCode   int           `json:"result_code,omitempty"`
	TaskID       string        `json:"task_id,omitempty"`
	Transmited   bool          `json:"transmited,omitempty"`
	TransmitedAt string        `json:"transmited_at,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
	Url          string        `json:"url,omitempty"`
	Meta         interface{}   `json:"meta,omitempty"`
}

type Order struct {
	BaseGrandTotal int      `json:"base_grand_total,omitempty"`
	GrandTotal     int      `json:"grand_total,omitempty"`
	Items          []string `json:"items"`
	TotalCount     int      `json:"total_count,omitempty"`
}

type Item struct {
	SKU string `json:"sku,omitempty"`
}

type Info struct {
	NameSpaced string `json:"name_spaced"`
	State      struct {
		Token string `json:"token"`
	}
}

type MeUser struct {
	Code   int    `json:"code"`
	Result Result `json:"result,omitempty"`
}
type Result struct {
	Addresses              interface{} `json:"addresses"`
	CreatedAt              string      `json:"created_at,omitempty"`
	CreatedIn              string      `json:"created_in,omitempty"`
	DisableAutoGroupChange int         `json:"disable_auto_group_change,omitempty"`
	Email                  string      `json:"email,omitempty"`
	FirstName              string      `json:"firstname,omitempty"`
	GroupID                int         `json:"group_id,omitempty"`
	ID                     int         `json:"id,omitempty"`
	LastName               string      `json:"lastname,omitempty"`
	StoreID                int         `json:"store_id,omitempty"`
	UpdatedAt              string      `json:"updated_at,omitempty"`
	WebsiteID              int         `json:"website_id,omitempty"`
}
