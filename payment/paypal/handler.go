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

	//ReturnUrl := "localhost:3000"
	//CancelUrl := "www.google.lt"

	//requestBytes := new(bytes.Buffer)
	//json.NewEncoder(requestBytes).Encode(request)
	//
	//req, err :=http.NewRequest("POST", paypalApi + "/v1/payments/payment", requestBytes)
	//helpers.PanicErr(err)
	//req.SetBasicAuth(user, pass)
	////resp, err := http.Post(paypalApi + "/v1/payments/payment", contentType, requestBytes)
	//client := &http.Client{}
	//
	//resp, err := client.Do(req)
	//helpers.PanicErr(err)
	//fmt.Println(resp)

	var requestFromClient request
	err := json.NewDecoder(r.Body).Decode(&requestFromClient)
	helpers.PanicErr(err)

	client, err := paypalsdk.NewClient(user, pass, paypalsdk.APIBaseSandBox)
	helpers.PanicErr(err)
	client.SetLog(os.Stdout)

	_, err = client.GetAccessToken()
	helpers.PanicErr(err)
	//fmt.Println("ACCESS TOKEN ", accessToken)

	amount := paypalsdk.Amount{
		Total:    requestFromClient.Transactions[0].Amount.Total,
		Currency: requestFromClient.Transactions[0].Amount.Currency}

	paymentResult, err := client.CreateDirectPaypalPayment(amount, redurectURL, cancelURL, "")
	helpers.PanicErr(err)
	//fmt.Printf("%+v", paymentResult)
	helpers.WriteJsonResult(w, paymentResult)
}

func Execute(w http.ResponseWriter, r *http.Request) {
	var requestFromClient request
	err := json.NewDecoder(r.Body).Decode(&requestFromClient)
	helpers.PanicErr(err)

	client, err := paypalsdk.NewClient(user, pass, paypalsdk.APIBaseSandBox)
	helpers.PanicErr(err)
	client.SetLog(os.Stdout)

	_, err = client.GetAccessToken()
	helpers.PanicErr(err)

	executeResult, err := client.ExecuteApprovedPayment(requestFromClient.PaymentID, requestFromClient.PayerID)
	helpers.PanicErr(err)

	helpers.WriteJsonResult(w, executeResult)
}
