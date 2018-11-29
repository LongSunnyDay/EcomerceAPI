package total

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/auth"
	"go-api-ws/helpers"
	"net/http"
	"time"
)

func GetTotals(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("ME CALLED")
	urlToken := r.URL.Query()["token"][0]
	urlCartId := r.URL.Query()["cartId"][0]
	token, err := auth.ParseToken(urlToken)
	helpers.PanicErr(err)
	var groupId float64
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			groupId = claims["groupId"].(float64)
		}
	}

	var totals Totals
	totals.getItems(urlCartId)
	totals.getSubtotal()

	var addressInfo AddressData
	_ = json.NewDecoder(r.Body).Decode(&addressInfo)
	totals.getShipping(addressInfo)

	rates := totals.getTaxRates(groupId)
	totals.calculateTax(rates)
	totals.calculateGrandtotal(rates)
	//fmt.Printf("%+v/n", totals)
	totalsResp := TotalsResp{
		Totals:totals}
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: totalsResp}
	response.SendResponse(w)

}
