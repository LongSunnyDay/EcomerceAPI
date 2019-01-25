package total

import (
	"go-api-ws/attribute"
	"go-api-ws/cart"
	"go-api-ws/config"
	"go-api-ws/discount"
	"go-api-ws/helpers"
	"go-api-ws/shipping"
	"go-api-ws/tax"
	"strconv"
)

func (t *Totals) CalculateTotals(urlCartId string, addressInformation AddressData, groupId int64) {
	t.GetItems(urlCartId)
	t.GetSubtotal()
	t.GetShipping(addressInformation)
	rates := t.GetTaxRates(groupId)
	t.CalculateTax(rates)
	t.GetDiscounts(urlCartId)
	t.CalculateGrandtotal(rates)
}

//NOT tested yet
func (t *Totals) GetDiscounts(cartId string) {
	var totalDiscount discount.Discount
	db, err := config.Conf.GetDb()
	helpers.CheckErr(err)
	if discount.CouponUsed {
		t.DiscountAmount = t.DiscountAmount + discount.CouponDiscountAmount
		discount.CouponUsed = false
	}
	for _, item := range t.Items {
		err = db.QueryRow("SELECT discountPercent, discountAmount FROM totalDiscount c WHERE sku=?", item.SKU).
			Scan(&totalDiscount.DiscountPercent, &totalDiscount.DiscountAmount)
		helpers.CheckErr(err)
		percentToCurrency := totalDiscount.DiscountPercent / 100 * item.RowTotalInclTax
		t.DiscountAmount = t.DiscountAmount + percentToCurrency + totalDiscount.DiscountAmount

	}
}

func (t *Totals) GetItems(cartId string) {
	userCart := cart.GetUserCartFromMongoByID(cartId)
	for _, item := range userCart.Items {
		totalsItem := Item{
			SKU:    item.SKU,
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

func (t *Totals) GetTaxRates(groupId int64) tax.Rules {
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

//noinspection GoRedundantTypeDeclInCompositeLit,GoRedundantTypeDeclInCompositeLit
func (t *Totals) CalculateGrandtotal(rules tax.Rules) {
	for _, item := range t.Items {
		t.BaseTaxAmount = t.BaseTaxAmount + item.TaxAmount
	}
	for _, item := range t.Items {
		t.BaseTaxAmount = t.BaseTaxAmount + item.TaxAmount
	}

	dc := shipping.GetConfig("config.yml")

	t.BaseGrandTotal = t.Subtotal - t.DiscountAmount + t.TaxAmount + t.ShippingInclTax
	t.BaseSubtotalWithDiscount = t.Subtotal + t.DiscountAmount
	t.SubtotalInclTax = t.Subtotal + t.TaxAmount
	t.SubtotalWithDiscount = t.Subtotal - t.DiscountAmount

	if dc.ShippingDiscountIsAvailable && dc.ShippingDiscountFrom < t.SubtotalInclTax && dc.ShippingDiscountTo > t.SubtotalInclTax {
		t.DiscountAmount = t.ShippingInclTax + t.DiscountAmount
		t.BaseGrandTotal = t.BaseGrandTotal - t.ShippingInclTax
		t.ShippingInclTax = 0
		t.BaseShippingInclTax = 0
	}

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
		Value: t.DiscountAmount}
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
