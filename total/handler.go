package total

import (
	"encoding/json"
	"go-api-ws/auth"
	"go-api-ws/helpers"
	"go-api-ws/payment_methods"
	"net/http"
)

var (
	totals      Totals
	addressInfo AddressData
	groupId     int64
)

func GetTotals(w http.ResponseWriter, r *http.Request) {

	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	urlCartId, err := helpers.GetCartIdFromUrl(r)
	helpers.PanicErr(err)
	if len(urlToken) > 0 {
		token := auth.ParseToken(urlToken)
		claims, err := auth.GetTokenClaims(token)
		helpers.CheckErr(err)
		if err != nil {
			helpers.WriteResultWithStatusCode(w, err, http.StatusForbidden)
		} else {
			if auth.CheckIfTokenIsNotExpired(claims) {
				groupIdFloat := claims["groupId"].(float64)
				groupId = int64(groupIdFloat)
			}
		}
	} else {
		groupId = 1
	}

	err = json.NewDecoder(r.Body).Decode(&addressInfo)
	helpers.PanicErr(err)

	totals.CalculateTotals(urlCartId, addressInfo, groupId)

	totalsResp := TotalsResp{
		Totals: totals}
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: totalsResp}
	response.SendResponse(w)
}

func GetTotalsWithPaymentMethods(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	urlCartId, err := helpers.GetCartIdFromUrl(r)
	helpers.PanicErr(err)
	if len(urlToken) > 0 {
		token := auth.ParseToken(urlToken)
		claims, err := auth.GetTokenClaims(token)
		helpers.CheckErr(err)
		if err != nil {
			helpers.WriteResultWithStatusCode(w, err, http.StatusForbidden)
		} else {
			if auth.CheckIfTokenIsNotExpired(claims) {
				groupIdFloat := claims["groupId"].(float64)
				groupId = int64(groupIdFloat)
			}
		}
	} else {
		groupId = 1
	}

	err = json.NewDecoder(r.Body).Decode(&addressInfo)
	helpers.PanicErr(err)

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
