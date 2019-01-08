package models

type Discount struct {
	Id     int    `json:"id,omitempty"`
	Sku     string    `json:"sku,omitempty"`
	DiscountPercent     int    `json:"discountPercent,omitempty"`
	DiscountAmount     int    `json:"discountAmount,omitempty"`
}
