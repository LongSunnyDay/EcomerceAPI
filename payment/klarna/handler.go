package klarna

import (
	"encoding/json"
	"github.com/Flaconi/go-klarna"
	"go-api-ws/helpers"
	"net/http"
	"net/url"
	"time"
)

const (
	KlarnaTestingApiEndpoint = "https://api.playground.klarna.com"
	user                     = "PK05992_f8487a01c7bd"
	pass                     = "GbLTYA8kmHMH2kty"
)

func CreateSession(w http.ResponseWriter, r *http.Request) {
	var request go_klarna.PaymentOrder
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.PanicErr(err)

	uri, err := url.Parse(KlarnaTestingApiEndpoint)
	helpers.PanicErr(err)

	conf := go_klarna.Config{
		BaseURL:     uri,
		Timeout:     time.Second * 10,
		APIUsername: user,
		APIPassword: pass}

	client := go_klarna.NewClient(conf)

	paymentSrv := go_klarna.NewPaymentSrv(client)

	ps, err := paymentSrv.CreateNewSession(&request)
	helpers.PanicErr(err)

	//fmt.Println(&ps.PaymentMethodCategories)

	helpers.WriteResultWithStatusCode(w, ps, http.StatusOK)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var request go_klarna.PaymentOrder
	err := json.NewDecoder(r.Body).Decode(&request)
	helpers.PanicErr(err)

	token := r.URL.Query().Get("token")

	uri, err := url.Parse(KlarnaTestingApiEndpoint)
	helpers.PanicErr(err)

	conf := go_klarna.Config{
		BaseURL:     uri,
		Timeout:     time.Second * 10,
		APIUsername: user,
		APIPassword: pass}

	client := go_klarna.NewClient(conf)

	paymentSrv := go_klarna.NewPaymentSrv(client)

	ps, err := paymentSrv.CreateNewOrder(token, &request)
	helpers.PanicErr(err)

	helpers.WriteResultWithStatusCode(w, ps, http.StatusOK)
}
