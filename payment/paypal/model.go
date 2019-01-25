package paypal

type (
	Transaction struct {
		Amount Amount `json:"amount"`
	}

	Amount struct {
		Total    string `json:"total"`
		Currency string `json:"currency"`
	}

	request struct {
		Transactions []Transaction `json:"transactions"`
		PaymentID    string        `json:"paymentID,omitempty"`
		PayerID      string        `json:"payerID,omitempty"`
	}
)
