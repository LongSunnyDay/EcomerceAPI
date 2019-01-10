package models

type Discount struct {
	Id     int    `json:"id,omitempty"`
	Sku     string    `json:"sku,omitempty"`
	DiscountPercent     float64    `json:"discountPercent,omitempty"`
	DiscountAmount     float64    `json:"discountAmount,omitempty"`
}
