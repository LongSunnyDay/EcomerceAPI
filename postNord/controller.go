package postNord

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"net/http"
	"os"
	"time"
)

func (userAddress UserAddressData) NewTransitTimeForm() (form TransitTimeForm) {

	form = TransitTimeForm{
		DateOfDeparture:         time.Now().Add(time.Hour * 48).Format("2006-01-02"),
		ServiceCode:             serviceCode,
		ServiceGroupCode:        serviceGroupCode,
		ResponseContent:         responseContent,
		FromAddressStreetName:   fromAddressStreetName,
		FromAddressStreetNumber: fromAddressStreetNumber,
		FromAddressPostalCode:   fromAddressPostalCode,
		FromAddressCountryCode:  fromAddressCountryCode,
		ToAddressStreetName:     userAddress.StreetName,
		ToAddressStreetNumber:   userAddress.StreetNumber,
		ToAddressPostalCode:     userAddress.PostalCode,
		ToAddressCountryCode:    userAddress.CountryCode}
	return
}

func (transitData TransitTimeForm) PostTransitData() *http.Request {
	req, err := http.NewRequest("GET", postNordUrl+transitTimeInfoUrl, nil)
	helpers.PanicErr(err)

	q := req.URL.Query()

	q.Add("apikey", apiKey)
	q.Add("dateOfDeparture", transitData.DateOfDeparture)
	q.Add("serviceCode", transitData.ServiceCode)
	q.Add("serviceGroupCode", transitData.ServiceGroupCode)
	q.Add("responseContent", transitData.ResponseContent)
	q.Add("fromAddressStreetName", transitData.FromAddressStreetName)
	q.Add("fromAddressStreetNumber", transitData.FromAddressStreetNumber)
	q.Add("fromAddressPostalCode", transitData.FromAddressPostalCode)
	q.Add("fromAddressCountryCode", transitData.FromAddressCountryCode)
	q.Add("toAddressStreetName", transitData.ToAddressStreetName)
	q.Add("toAddressStreetNumber", transitData.ToAddressStreetNumber)
	q.Add("toAddressPostalCode", transitData.ToAddressPostalCode)
	q.Add("toAddressCountryCode", transitData.ToAddressCountryCode)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req
}

//noinspection GoRedundantTypeDeclInCompositeLit
func (form *OrderPickupForm) MakeOrderPickup() {

	f, err := os.OpenFile("./logs/Order_"+form.Order.OrderReference+"_log.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	helpers.CheckErr(err)
	defer f.Close()

	log.SetOutput(f)

	form.Shipment.Service = OrderService{
		BasicServiceCode: serviceCode}
	sgc := serviceGroupCode
	form.Shipment.Service.AdditionalServiceCode = append(form.Shipment.Service.AdditionalServiceCode, &sgc)

	form.Location = Location{
		Place:        place,
		StreetName:   fromAddressStreetName,
		StreetNumber: fromAddressStreetNumber,
		PostalCode:   fromAddressPostalCode,
		City:         city,
		CountryCode:  fromAddressCountryCode}

	form.Order.CustomerNumber = customerNumber
	form.Order.ContactName = contactName
	form.Order.ContactEmail = contactEmail
	form.Order.PhoneNumber = phoneNumber
	form.Order.SmsNumber = smsNumber
	form.Order.EntryCode = entryCode

	form.DateAndTimes = DateAndTimes{
		ReadyPickupDate:    time.Now().Add(time.Hour * 48).Format("2006-01-02T15:04:05Z"),
		LatestPickupDate:   time.Now().AddDate(0, 0, 7).Format("2006-01-02T15:04:05Z"),
		EarliestPickupDate: time.Now().Add(time.Hour * 48).Format("2006-01-02T15:04:05Z")}

	form.Pickup = []*Pickup{
		&Pickup{TypeOfItem: parcel,
			NoUnits: 1}}

	bodyBites := new(bytes.Buffer)
	err = json.NewEncoder(bodyBites).Encode(form)
	helpers.PanicErr(err)

	req, err := http.NewRequest("POST", postNordUrl+orderPickupUrl, bodyBites)
	helpers.PanicErr(err)

	q := req.URL.Query()
	q.Add("apikey", apiKey)

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	log.Printf("%+v\n", req)

	hc := &http.Client{}

	resp, err := hc.Do(req)
	helpers.PanicErr(err)

	log.Printf("%+v\n", resp)

	var pickupResp OrderPickupResponse

	err = json.NewDecoder(resp.Body).Decode(&pickupResp)
	helpers.PanicErr(err)

	log.Printf("%+v\n", pickupResp)

	SavePickupData(pickupResp)

}

func SavePickupData(response OrderPickupResponse) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO pickup_order("+
		"order_id, "+
		"pickup_time) "+
		"VALUES(?, ?)",
		response.OrderId,
		response.PickupTime)
	helpers.PanicErr(err)
}
