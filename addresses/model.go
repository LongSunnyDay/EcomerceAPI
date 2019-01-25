package addresses

type (
	Address struct {
		ID              int64    `json:"id,omitempty"`
		CustomerID      int64    `json:"customer_id"`
		Region          Region   `json:"region"`
		RegionID        int64    `json:"region_id"`
		CountryID       string   `json:"country_id"`
		Street          []string `json:"street"`
		Telephone       string   `json:"telephone"`
		Postcode        string   `json:"postcode"`
		City            string   `json:"city"`
		Firstname       string   `json:"firstname"`
		Lastname        string   `json:"lastname"`
		DefaultShipping bool     `json:"default_shipping"`
		Company         string   `json:"company"`
		VatID           string   `json:"vat_id"`
		DefaultBilling  bool     `json:"default_billing"`
		Email           string   `json:"email"`
		StreetLine0     string   `json:"street_line_0,omitempty"`
		StreetLine1     string   `json:"street_line_1,omitempty"`
	}

	Region struct {
		RegionCode string `json:"region_code" bson:"region_code"`
		Region     string `json:"region" bson:"region"`
		RegionID   int64  `json:"region_id" bson:"region_id"`
	}
)
