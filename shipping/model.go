package shipping

type (
	methods []Method

	Method struct {
		Id           int     `json:"id,omitempty"`
		CarrierCode  string  `json:"carrier_code"`
		MethodCode   string  `json:"method_code"`
		CarrierTitle string  `json:"carrier_title"`
		MethodTitle  string  `json:"method_title"`
		Amount       float64 `json:"amount"`
		BaseAmount   float64 `json:"base_amount"`
		Available    bool    `json:"available"`
		ErrorMessage string  `json:"error_message"`
		PriceExclTax int     `json:"price_excl_tax"`
		PriceInclTax float64 `json:"price_incl_tax"`
	}

	DiscountConfig struct {
		ShippingDiscountIsAvailable bool    `yaml:"shippingDiscountIsAvailable"`
		ShippingDiscountFrom        float64 `yaml:"shippingDiscountFrom"`
		ShippingDiscountTo          float64 `yaml:"shippingDiscountTo"`
	}
)

