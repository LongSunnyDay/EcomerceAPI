package postNord

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

