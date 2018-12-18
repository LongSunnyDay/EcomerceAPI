package total

import (
	"go-api-ws/attribute"
	"go-api-ws/cart"
	"go-api-ws/helpers"
	"go-api-ws/payment"
	"go-api-ws/shipping"
	"go-api-ws/tax"
	"strconv"
)

type TotalsResp struct {
	Totals         Totals           `json:"totals"`
	PaymentMethods []payment.Method `json:"payment_methods"`
}

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
	Items                      []*Item   `json:"items"`
	TotalSegments              []Segment `json:"total_segments"`
}

type Item struct {
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

func (t *Totals) CalculateTotals(urlCartId string, addressInformation AddressData, groupId float64) {
	t.GetItems(urlCartId)
	t.GetSubtotal()
	t.GetShipping(addressInformation)
	rates := t.GetTaxRates(groupId)
	t.CalculateTax(rates)
	t.CalculateGrandtotal(rates)
}

func (t *Totals) GetItems(cartId string) {
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
			totalsItem.Options = append(totalsItem.Options, &totalOptions)
		}
		t.Items = append(t.Items, &totalsItem)
	}
}

func (t *Totals) GetSubtotal() {
	for _, item := range t.Items {
		qtyFloat := float64(item.Qty)
		item.BasePrice = item.Price
		item.BaseRowTotal = item.BasePrice * qtyFloat
		item.RowTotal = item.Price * qtyFloat

		t.ItemsQty = t.ItemsQty + qtyFloat
		t.Subtotal = t.Subtotal + item.RowTotal
		t.BaseSubtotal = t.Subtotal
	}
}

func (t *Totals) GetTaxRates(groupId float64) tax.Rules {
	var taxRates tax.Rules
	taxRates.GroupId = int(groupId)
	rules := taxRates.GetRates()
	return rules
}

func (t *Totals) CalculateTax(rules tax.Rules) {

	rateInt, err := strconv.Atoi(rules.Rates.Percent)
	helpers.PanicErr(err)

	rateFloat := float64(rateInt)
	rateFloat = rateFloat / 100
	taxAmount := t.Subtotal * rateFloat

	t.TaxAmount = taxAmount
	for _, item := range t.Items {
		item.TaxPercent = float64(rateInt)

		//item.BasePriceInclTax = item.BasePrice * (1 + rateFloat)
		item.BasePriceInclTax = item.BasePrice

		item.RowTotalInclTax = item.RowTotal * (1 + rateFloat)
		//item.RowTotalInclTax = item.RowTotal

		item.BaseTaxAmount = item.RowTotal * (item.TaxPercent / 100)

		//item.PriceInclTax = item.Price * (1 + rateFloat)
		item.PriceInclTax = item.Price

		item.TaxAmount = item.RowTotal * (item.TaxPercent / 100)
	}
}

func (t *Totals) GetShipping(information AddressData) {
	method := shipping.GetShippingMethod(information.AddressInformation.ShippingCarrierCode, information.AddressInformation.ShippingMethodCode)

	t.ShippingInclTax = method.PriceInclTax
	t.ShippingAmount = method.Amount
	t.BaseShippingAmount = method.BaseAmount
	t.BaseShippingInclTax = method.PriceInclTax
}

func (t *Totals) CalculateGrandtotal(rules tax.Rules) {
	for _, item := range t.Items {
		t.BaseTaxAmount = t.BaseTaxAmount + item.TaxAmount

	}
	t.BaseGrandTotal = t.Subtotal + t.DiscountAmount + t.TaxAmount + t.ShippingInclTax
	t.BaseSubtotalWithDiscount = t.Subtotal + t.DiscountAmount
	t.SubtotalInclTax = t.Subtotal + t.TaxAmount
	t.SubtotalWithDiscount = t.Subtotal - t.DiscountAmount

	segment := Segment{
		Code:  "subtotal",
		Title: "Subtotal",
		Value: t.SubtotalInclTax}
	t.TotalSegments = append(t.TotalSegments, segment)

	segment = Segment{
		Code:  "shipping",
		Title: "Shipping & Handling (Flat Rate - Fixed)",
		Value: t.ShippingInclTax}
	t.TotalSegments = append(t.TotalSegments, segment)

	segment = Segment{
		Code:  "discount",
		Title: "Discount",
		Value: 0}
	t.TotalSegments = append(t.TotalSegments, segment)

	segment = Segment{
		Code:  "tax",
		Title: "Tax",
		Value: t.TaxAmount,
		Area:  "taxes",
		ExtensionAttribute: ExtensionAttribute{
			TaxGrandtotalDetails: []TaxGrandtotalDetail{
				TaxGrandtotalDetail{
					Amount: t.TaxAmount,
					Rates: []Rate{
						Rate{
							Percent: rules.Rates.Percent,
							Title:   "VAT23-pl"}},
					GroupId: rules.GroupId}}}}
	t.TotalSegments = append(t.TotalSegments, segment)

	segment = Segment{
		Code:  "grand_total",
		Title: "Grand Total",
		Value: t.BaseGrandTotal,
		Area:  "footer"}
	t.TotalSegments = append(t.TotalSegments, segment)

}

func getDiscount() {

}
