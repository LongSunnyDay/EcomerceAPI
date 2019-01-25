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

const (
	postNordUrl        = "https://atapi2.postnord.com"
	transitTimeInfoUrl = "/rest/transport/v1/transittime/getTransitTimeInformation.json"
	// Todo only Sweden endpoint exist
	orderPickupUrl          = "/rest/order/v1/pickup/SE"
	apiKey                  = "fb4a28c6fa3760efcd5e2a997a473a7c"
	serviceCode             = "19"
	serviceGroupCode        = "NO"
	fromAddressStreetName   = "Nedre Skøyen vei"
	fromAddressStreetNumber = "2"
	fromAddressPostalCode   = "0276"
	fromAddressCountryCode  = "NO"
	responseContent         = "full"
	// Todo pickup order data
	// Order
	customerNumber = "1234567891"
	orderReference = "Ref-1212122A"
	contactName    = "Vardenis Pavardenis"
	contactEmail   = "vardenis@pavardenis.com"
	phoneNumber    = "+4670788888"
	smsNumber      = "+4670788888"
	entryCode      = "8216"
	// Location
	place = "garage"
	city  = "Oslo"

	parcel = "parcel"
)

type (
	TransitTimeForm struct {
		DateOfDeparture         string
		ServiceCode             string
		ServiceGroupCode        string
		FromAddressStreetName   string
		FromAddressStreetNumber string
		FromAddressPostalCode   string
		FromAddressCountryCode  string
		ToAddressStreetName     string
		ToAddressStreetNumber   string
		ToAddressPostalCode     string
		ToAddressCountryCode    string
		ResponseContent         string
	}
	UserAddressData struct {
		StreetName   string `json:"street_name"`
		StreetNumber string `json:"street_number"`
		PostalCode   string `json:"postal_code"`
		CountryCode  string `json:"country_code"`
	}

	ComplexFieldName struct {
		TransitTimeResponse TransitTimeResponse `json:"se.posten.loab.lisp.notis.publicapi.serviceapi.TransitTimeResponse"`
	}
	TransitTimeResponse struct {
		CompositeFault CompositeFault `json:"compositeFault,omitempty"`
		TransitTimes   *[]TransitTime `json:"transitTimes"`
	}
	TransitTime struct {
		DateOfDeparture     string    `json:"dateOfDeparture"`
		LatestTimeOfBooking string    `json:"latestTimeOfBooking"`
		DeliveryDate        string    `json:"deliveryDate"`
		TransitTimeInDays   int64     `json:"transitTimeInDays"`
		PossibleDeviation   bool      `json:"possibleDeviation"`
		Service             Service   `json:"service"`
		DaysPickup          *[]string `json:"daysPickup"`
	}
	Service struct {
		Code         string `json:"code"`
		GroupCode    string `json:"groupCode"`
		Name         string `json:"name"`
		Pickup       bool   `json:"pickup"`
		Distribution bool   `json:"distribution"`
	}
	CompositeFault struct {
		Faults *[]Fault `json:"faults,omitempty"`
	}
	Fault struct {
		FaultCode       string        `json:"faultCode"`
		ExplanationText string        `json:"explanationText"`
		ParamValues     *[]ParamValue `json:"paramValues"`
	}
	ParamValue struct {
		Param string `json:"param"`
		Value string `json:"value"`
	}

	OrderPickupForm struct {
		Shipment     Shipment     `json:"shipment"`
		Location     Location     `json:"location"`
		Order        Order        `json:"order"`
		Pickup       []*Pickup    `json:"pickup"`
		DateAndTimes DateAndTimes `json:"dateAndTimes"`
	}
	Shipment struct {
		Service OrderService `json:"service"`
		Items   []*string    `json:"items"`
	}
	OrderService struct {
		BasicServiceCode      string    `json:"basicServiceCode"`
		AdditionalServiceCode []*string `json:"additionalServiceCode"`
	}
	Location struct {
		Place        string `json:"place"`
		StreetName   string `json:"streetName"`
		StreetNumber string `json:"streetNumber"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		CountryCode  string `json:"countryCode"`
	}
	Order struct {
		CustomerNumber string `json:"customerNumber"`
		OrderReference string `json:"orderReference"`
		ContactName    string `json:"contactName"`
		ContactEmail   string `json:"contactEmail"`
		PhoneNumber    string `json:"phoneNumber"`
		SmsNumber      string `json:"smsNumber"`
		EntryCode      string `json:"entryCode"`
	}
	Pickup struct {
		TypeOfItem string `json:"typeOfItem"`
		NoUnits    int64  `json:"noUnits"`
	}
	DateAndTimes struct {
		ReadyPickupDate    string `json:"readyPickupDate"`
		LatestPickupDate   string `json:"latestPickupDate"`
		EarliestPickupDate string `json:"earliestPickupDate"`
	}

	OrderPickupResponse struct {
		OrderId    string `json:"orderId"`
		PickupTime string `json:"pickupTime"`
	}
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

	f, err := os.OpenFile("./logs/Order_"+ form.Order.OrderReference + "_log.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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

func SavePickupData(response OrderPickupResponse)  {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO pickup_order(" +
		"order_id, " +
		"pickup_time) " +
		"VALUES(?, ?)",
		response.OrderId,
		response.PickupTime)
	helpers.PanicErr(err)
}
