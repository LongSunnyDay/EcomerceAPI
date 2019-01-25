package total

import (
	"go-api-ws/payment_methods"
)

type (
	TotalsResp struct {
		Totals         Totals                   `json:"totals"`
		PaymentMethods []payment_methods.Method `json:"payment_methods"`
	}

	Totals struct {
		GrandTotal                 float64   `json:"grand_total"`
		BaseGrandTotal             float64   `json:"base_grand_total"`
		Subtotal                   float64   `json:"subtotal"`
		BaseSubtotal               float64   `json:"base_subtotal"`
		DiscountAmount             float64   `json:"discount_amount"`
		BaseDiscountAmount         float64   `json:"base_discount_amount"`
		SubtotalWithDiscount       float64   `json:"subtotal_with_discount"`
		BaseSubtotalWithDiscount   float64   `json:"base_subtotal_with_discount"`
		ShippingAmount             float64   `json:"shipping_amount"`
		BaseShippingAmount         float64   `json:"base_shipping_amount"`
		ShippingDiscountAmount     float64   `json:"shipping_discount_amount"`
		BaseShippingDiscountAmount float64   `json:"base_shipping_discount_amount"`
		TaxAmount                  float64   `json:"tax_amount"`
		BaseTaxAmount              float64   `json:"base_tax_amount"`
		WeeTaxAppliedAmount        float64   `json:"wee_tax_applied_amount"`
		ShippingTaxAmount          float64   `json:"shipping_tax_amount"`
		BaseShippingTaxAmount      float64   `json:"base_shipping_tax_amount"`
		SubtotalInclTax            float64   `json:"subtotal_incl_tax"`
		ShippingInclTax            float64   `json:"shipping_incl_tax"`
		BaseShippingInclTax        float64   `json:"base_shipping_incl_tax"`
		BaseCurrencyCode           string    `json:"base_currency_code"`
		QuoteCurrencyCode          string    `json:"quote_currency_code"`
		ItemsQty                   float64   `json:"items_qty"`
		Items                      []*Item   `json:"items"`
		TotalSegments              []Segment `json:"total_segments"`
	}

	Item struct {
		SKU                  string    `json:"sku"`
		ItemId               int       `json:"item_id"`
		Price                float64   `json:"price"`
		BasePrice            float64   `json:"base_price"`
		Qty                  float64   `json:"qty"`
		RowTotal             float64   `json:"row_total"`
		BaseRowTotal         float64   `json:"base_row_total"`
		RowTotalWithDiscount float64   `json:"row_total_with_discount"`
		TaxAmount            float64   `json:"tax_amount"`
		BaseTaxAmount        float64   `json:"base_tax_amount"`
		TaxPercent           float64   `json:"tax_percent"`
		DiscountAmount       float64   `json:"discount_amount"`
		BaseDiscountAmount   float64   `json:"base_discount_amount"`
		DiscountPercent      float64   `json:"discount_percent"`
		PriceInclTax         float64   `json:"price_incl_tax"`
		BasePriceInclTax     float64   `json:"base_price_incl_tax"`
		RowTotalInclTax      float64   `json:"row_total_incl_tax"`
		BaseRowTotalInclTax  float64   `json:"base_row_total_incl_tax"`
		WeeTaxAppliedAmount  float64   `json:"wee_tax_applied_amount"`
		WeeTaxApplied        float64   `json:"wee_tax_applied"`
		Name                 string    `json:"name"`
		Options              []*Option `json:"options"`
	}

	Option struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	Segment struct {
		Code               string             `json:"code"`
		Title              string             `json:"title"`
		Value              float64            `json:"value"`
		Area               string             `json:"area,omitempty"`
		ExtensionAttribute ExtensionAttribute `json:"extension_attribute,omitempty"`
	}

	ExtensionAttribute struct {
		TaxGrandtotalDetails []TaxGrandtotalDetail `json:"tax_grandtotal_details"`
	}

	TaxGrandtotalDetail struct {
		Amount  float64 `json:"amount"`
		Rates   []Rate  `json:"rates"`
		GroupId int     `json:"group_id"`
	}

	Rate struct {
		Percent string `json:"percent"`
		Title   string `json:"title"`
	}

	AddressData struct {
		AddressInformation struct {
			ShippingAddress struct {
				CountryId string `json:"countryId"`
			}
			ShippingCarrierCode string `json:"shippingCarrierCode"`
			ShippingMethodCode  string `json:"shippingMethodCode"`
		} `json:"addressInformation"`
	}
)
