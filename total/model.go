package total

import (
	"go-api-ws/attribute"
	"go-api-ws/cart"
	"go-api-ws/helpers"
	"go-api-ws/shipping"
	"go-api-ws/tax"
	"strconv"
)

type Totals struct {
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
	Items                      []Item    `json:"items"`
	TotalSegments              []Segment `json:"total_segments"`
}

type Item struct {
	ItemId               int      `json:"item_id"`
	Price                float64  `json:"price"`
	BasePrice            float64  `json:"base_price"`
	Qty                  float64  `json:"qty"`
	RowTotal             float64  `json:"row_total"`
	BaseRowTotal         float64  `json:"base_row_total"`
	RowTotalWithDiscount float64  `json:"row_total_with_discount"`
	TaxAmount            float64  `json:"tax_amount"`
	BaseTaxAmount        float64  `json:"base_tax_amount"`
	TaxPercent           float64  `json:"tax_percent"`
	DiscountAmount       float64  `json:"discount_amount"`
	BaseDiscountAmount   float64  `json:"base_discount_amount"`
	DiscountPercent      float64  `json:"discount_percent"`
	PriceInclTax         float64  `json:"price_incl_tax"`
	BasePriceInclTax     float64  `json:"base_price_incl_tax"`
	RowTotalInclTax      float64  `json:"row_total_incl_tax"`
	BaseRowTotalInclTax  float64  `json:"base_row_total_incl_tax"`
	WeeTaxAppliedAmount  float64  `json:"wee_tax_applied_amount"`
	WeeTaxApplied        float64  `json:"wee_tax_applied"`
	Name                 string   `json:"name"`
	Options              []Option `json:"options"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Segment struct {
	Code               string             `json:"code"`
	Title              string             `json:"title"`
	Value              float64            `json:"value"`
	Area               string             `json:"area,omitempty"`
	ExtensionAttribute ExtensionAttribute `json:"extension_attribute,omitempty"`
}

type ExtensionAttribute struct {
	TaxGrandtotalDetails []TaxGrandtotalDetail `json:"tax_grandtotal_details"`
}

type TaxGrandtotalDetail struct {
	Amount  float64 `json:"amount"`
	Rates   []Rate  `json:"rates"`
	GroupId int     `json:"group_id"`
}

type Rate struct {
	Percent string `json:"percent"`
	Title   string `json:"title"`
}

type AddressData struct {
	AddressInformation struct {
		ShippingAddress struct {
			CountryId string `json:"countryId"`
		}
		ShippingCarrierCode string `json:"shippingCarrierCode"`
		ShippingMethodCode  string `json:"shippingMethodCode"`
	} `json:"addressInformation"`
}

func (t *Totals) calculateTotals() {

}

func (t *Totals) getItems(cartId string) {
	cartItems := cart.GetUserCartFromMongoByID(cartId)
	for _, item := range cartItems {
		totalsItem := Item{
			ItemId: item.ItemID,
			Qty:    item.QTY,
			Name:   item.Name,
			Price:  item.Price}
		var totalItemOptions []attribute.ItemAttribute

		for _, option := range item.ProductOption.ExtensionAttributes.ConfigurableItemOptions {
			totalItemOptions = append(totalItemOptions, attribute.GetAttributeNameFromSolr(option.OptionsID, option.OptionValue))
		}

		for _, option := range totalItemOptions {
			totalOptions := Option{
				Value: option.Label,
				Label: option.Name}
			totalsItem.Options = append(totalsItem.Options, totalOptions)
		}
		t.Items = append(t.Items, totalsItem)
	}
	//for _, item := range t.Items {
	//	fmt.Println(item.ItemId)
	//	fmt.Println(item.Qty)
	//	fmt.Println(item.Name)
	//	fmt.Println(item.Price)
	//	fmt.Println(item.Options)
	//}
}

func (t *Totals) getSubtotalTotal() {
	var subValue float64
	for _, item := range t.Items {
		price := item.Price
		item.RowTotal = item.Price * item.Qty
		subValue = price * item.Qty
	}
	segment := Segment{
		Code:  "subtotal",
		Title: "Subtotal",
		Value: subValue}
	t.TotalSegments = append(t.TotalSegments, segment)
}

func (t *Totals) getTaxRates(groupId float64) tax.Rules {
	var taxRates tax.Rules
	taxRates.GroupId = int(groupId)
	rules := taxRates.GetRates()
	return rules
}

func (t *Totals) calculateTax(rules tax.Rules) float64 {
	rateInt, err := strconv.Atoi(rules.Rates.Percent)
	helpers.PanicErr(err)
	rateFloat := float64(rateInt)
	rateFloat = rateFloat / 100
	taxAmount := t.Subtotal * rateFloat
	segment := Segment{
		Code:  "tax",
		Title: "Tax",
		Value: taxAmount,
		Area:  "taxes",
		ExtensionAttribute: ExtensionAttribute{
			TaxGrandtotalDetails: []TaxGrandtotalDetail{
				TaxGrandtotalDetail{
					Amount: taxAmount,
					Rates: []Rate{
						Rate{
							Percent: rules.Rates.Percent,
							Title:   "VAT23-pl"}},
					GroupId: rules.GroupId}}}}
	t.TotalSegments = append(t.TotalSegments, segment)
	t.TaxAmount = taxAmount
	for _, item := range t.Items {
		item.TaxPercent = float64(rateInt)
		item.TaxAmount = item.RowTotal * item.TaxPercent
		item.PriceInclTax = item.Price + item.TaxAmount
		item.BasePriceInclTax = item.Price * (1 + rateFloat)
		item.RowTotalInclTax = item.BasePriceInclTax * item.Qty
	}
	return taxAmount
}

func (t *Totals) getShipping(information AddressData) {
	method := shipping.GetShippingMethod(information.AddressInformation.ShippingCarrierCode, information.AddressInformation.ShippingMethodCode)
	segment := Segment{
		Code:  "shipping",
		Title: "Shipping & Handling (" + method.CarrierTitle + " - " + method.MethodTitle + ")",
		Value: method.PriceInclTax}
	t.TotalSegments = append(t.TotalSegments, segment)
	t.ShippingInclTax = method.PriceInclTax

}

func (t *Totals) calculateGrandtotal() {
	t.GrandTotal = t.Subtotal + t.TaxAmount + t.ShippingInclTax
	segment := Segment{
		Code:  "grand_total",
		Title: "Grand Total",
		Value: t.GrandTotal,
		Area:  "footer"}
	t.TotalSegments = append(t.TotalSegments, segment)

	//fmt.Println(t.TotalSegments)
}

func getDiscount() {

}
