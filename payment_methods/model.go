package payment_methods

const (
	collectionName = "payment_methods"
)

type (
	Methods []Method

	Method struct {
		Id             int    `json:"id,omitempty"`
		Code           string `json:"code"`
		Title          string `json:"title"`
		IsServerMethod bool   `json:"is_server_method,omitempty"`
	}
)

