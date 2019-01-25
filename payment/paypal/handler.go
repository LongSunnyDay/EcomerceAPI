package paypal

import (
	"encoding/json"
	"github.com/netlify/PayPal-Go-SDK"
	"go-api-ws/helpers"
	"net/http"
	"os"
)

const (
	user        = "AShDJc2Z2Jg0XVjH2V8M3v0d68QdB8B7xXceNdNVxmnLpxAU1P32L2tVPk52953w2KXK9Hmcog_wQzkN"
	pass        = "ELjl13pl77EvUNckgmMVazobPbWO42zIAqblRHxPBG75LWcD_yw76rkPGqJlUC480a81IC4xzlookgWZ"
	paypalApi   = "https://api.sandbox.paypal.com"
	contentType = "application/json"
	redurectURL = "http://www.duckduckgo.com"
	cancelURL   = "https://www.google.lt"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var requestFromClient request
	err := json.NewDecoder(r.Body).Decode(&requestFromClient)
	helpers.PanicErr(err)

	client, err := paypalsdk.NewClient(user, pass, paypalsdk.APIBaseSandBox)
	helpers.PanicErr(err)
	err = client.SetLog(os.Stdout)
	helpers.CheckErr(err)

	_, err = client.GetAccessToken()
	helpers.PanicErr(err)

	amount := paypalsdk.Amount{
		Total:    requestFromClient.Transactions[0].Amount.Total,
		Currency: requestFromClient.Transactions[0].Amount.Currency}

	paymentResult, err := client.CreateDirectPaypalPayment(amount, redurectURL, cancelURL, "")
	helpers.PanicErr(err)
	helpers.WriteJsonResult(w, paymentResult)
}

func Execute(w http.ResponseWriter, r *http.Request) {
	var requestFromClient request
	err := json.NewDecoder(r.Body).Decode(&requestFromClient)
	helpers.PanicErr(err)

	client, err := paypalsdk.NewClient(user, pass, paypalsdk.APIBaseSandBox)
	helpers.PanicErr(err)
	err = client.SetLog(os.Stdout)
	helpers.CheckErr(err)
	_, err = client.GetAccessToken()
	helpers.PanicErr(err)

	executeResult, err := client.ExecuteApprovedPayment(requestFromClient.PaymentID, requestFromClient.PayerID)
	helpers.PanicErr(err)

	helpers.WriteJsonResult(w, executeResult)
}
