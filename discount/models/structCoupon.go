package models

type Coupon struct {
	Id              int       `json:"id,omitempty"`
	DiscountPercent float64   `json:"discountPercent,omitempty"`
	DiscountAmount  float64   `json:"discountAmount,omitempty"`
	ExpirationDate  string `json:"expirationDate,omitempty"`
	UsageLimit      float64   `json:"usageLimit,omitempty"`
	TimesUsed       float64   `json:"timesUsed,omitempty"`
	CreatedAt       string   `json:"createdAt,omitempty"`
}
