package helpers

const (
	HeaderContentType   = "Content-Type"
	MIMEApplicationJSON = "application/json"
)

type (
	Response struct {
		Code   int         `json:"code,omitempty"`
		Result interface{} `json:"result,omitempty"`
		Meta   interface{} `json:"meta,omitempty"`
	}
	Closer interface {
		Close() error
	}
)
