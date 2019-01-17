package postNord

import (
	"encoding/json"
	"go-api-ws/helpers"
	"net/http"
)

func GetTransitTimeInformation(w http.ResponseWriter, r *http.Request) {
	var userAddressData UserAddressData
	err := json.NewDecoder(r.Body).Decode(&userAddressData)
	helpers.PanicErr(err)

	form := userAddressData.NewTransitTimeForm()
	req := form.PostTransitData()

	hc := &http.Client{}

	resp, err := hc.Do(req)
	helpers.PanicErr(err)

	var transitTimeResp ComplexFieldName

	err = json.NewDecoder(resp.Body).Decode(&transitTimeResp)
	helpers.CheckErr(err)

	helpers.WriteResultWithStatusCode(w, transitTimeResp, http.StatusOK)
}
