package paypal

type Transaction struct {
	Amount Amount `json:"amount"`
}

type Amount struct {
	Total    string `json:"total"`
	Currency string `json:"currency"`
}

type request struct {
	Transactions []Transaction `json:"transactions"`
	PaymentID    string        `json:"paymentID,omitempty"`
	PayerID      string        `json:"payerID,omitempty"`
}

//type Payer struct {
//	PaymentMethod string `json:"payment_method"`
//}

//type RedirectUrls struct {
//	ReturnUrl string `json:"return_url"`
//	CancelUrl string `json:"cancel_url"`
//}

//type RequestToPaypalBody struct {
//	Intent       string       `json:"intent"`
//	Payer        Payer        `json:"payer"`
//	Transactions []Transaction `json:"transactions"`
//	RedirectUrls RedirectUrls `json:"redirect_urls"`
//}
