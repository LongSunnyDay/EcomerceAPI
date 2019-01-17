package total

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/auth"
	"go-api-ws/helpers"
	"go-api-ws/payment_methods"
	"net/http"
	"time"
)

func GetTotals(w http.ResponseWriter, r *http.Request) {
	urlToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	var groupId int64
	if len(urlToken) > 0 {
		token, err := auth.ParseToken(urlToken)
		helpers.PanicErr(err)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				groupIdFloat := claims["groupId"].(float64)
				groupId = int64(groupIdFloat)
			}
		}
	} else {
		groupId = 1
	}

	var totals Totals
	var addressInfo AddressData
	_ = json.NewDecoder(r.Body).Decode(&addressInfo)

	totals.CalculateTotals(urlCartId, addressInfo, groupId)

	totalsResp := TotalsResp{
		Totals: totals}
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: totalsResp}
	response.SendResponse(w)
}

func GetTotalsWithPaymentMethods(w http.ResponseWriter, r *http.Request) {
	urlToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	var groupId int64
	if len(urlToken) > 0 {
		token, err := auth.ParseToken(urlToken)
		helpers.PanicErr(err)
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				groupIdFloat := claims["groupId"].(float64)
				groupId = int64(groupIdFloat)
			}
		}
	} else {
		groupId = 1
	}

	var totals Totals
	var addressInfo AddressData
	_ = json.NewDecoder(r.Body).Decode(&addressInfo)

	totals.CalculateTotals(urlCartId, addressInfo, groupId)
	paymentMethods := payment_methods.GetActualPaymentMethodsFromDb()

	totalsResp := TotalsResp{
		Totals:         totals,
		PaymentMethods: paymentMethods}
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: totalsResp}
	response.SendResponse(w)
}
