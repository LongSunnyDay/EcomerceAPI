package order

import (
	"errors"
	"fmt"
	"go-api-ws/addresses"
	"go-api-ws/cart"
	"go-api-ws/payment_methods"
	"go-api-ws/product"
	"go-api-ws/stock"
	"go-api-ws/total"
	"go-api-ws/user"
	"time"
)

// Pure function
func CheckIfItemsMatchInCart(cart cart.Cart, data PlaceOrderData) (err error) {
	fmt.Println(len(cart.Items), "<- cart - data ->", len(data.Products))

	if len(cart.Items) == len(data.Products) {
		for i, item := range data.Products {
			if cart.Items[i].SKU != item.Sku || cart.Items[i].QTY != item.Qty {
				err := errors.New("items in order and in cart doesn't match by SKU or QTY")
				return err
			} else {
				fmt.Println("All good, order item SKU -> ", item.Sku)
				return nil
			}
		}
	} else {
		err := errors.New("items amount in cart and in order is not the same")
		return err
	}
	return nil
}

func CheckStockStatus(data PlaceOrderData) (err error) {
	var orderStock []stock.DataStock
	for _, item := range data.Products {
		var SSOTItem stock.DataStock
		SSOTItem.GetDataFromDbBySku(item.Sku)
		orderStock = append(orderStock, SSOTItem)
	}
	for _, item := range data.Products {
		for _, stockItem := range orderStock {
			err := stockItem.CheckSOOT(item.Sku, item.Qty)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func FinalPriceCheck(data PlaceOrderData) (err error) {
	for _, item := range data.Products {
		checkPrice := product.GetProductPriceFromDbBySku(item.Sku, item.FinalPrice)
		if !checkPrice {
			err = errors.New("Product %v price doesn't match with price in db. " + item.Name)
			return err
		}
	}
	return nil
}

// Pure function
func AssignDataToBillingAddressAndSaveIt(data PlaceOrderData) (billingAddress addresses.Address) {
	billingAddress.RegionID = data.AddressInformation.BillingAddress.RegionId
	billingAddress.CountryID = data.AddressInformation.BillingAddress.CountryId
	billingAddress.StreetLine0 = data.AddressInformation.BillingAddress.Street[0]
	billingAddress.StreetLine1 = data.AddressInformation.BillingAddress.Street[1]
	billingAddress.Postcode = data.AddressInformation.BillingAddress.Postcode
	billingAddress.City = data.AddressInformation.BillingAddress.City
	billingAddress.Firstname = data.AddressInformation.BillingAddress.Firstname
	billingAddress.Lastname = data.AddressInformation.BillingAddress.Lastname
	billingAddress.Email = data.AddressInformation.BillingAddress.Email
	billingAddress.DefaultBilling = true

	return
}

// Pure function
func FormatOrderHistoryItems(totals total.Totals, quoteId int64) (orderItems []Item) {
	for _, itemFromTotals := range totals.Items {
		var orderItem Item
		//orderItem.AmountRefunded
		//orderItem.BaseAmountRefunded
		orderItem.BaseDiscountAmount = itemFromTotals.BaseDiscountAmount
		//orderItem.BaseDiscountInvoiced
		orderItem.BasePrice = itemFromTotals.BasePrice
		//orderItem.BaseRowInvoiced
		orderItem.RowTotal = itemFromTotals.RowTotal
		orderItem.BaseTaxAmount = itemFromTotals.BaseTaxAmount
		orderItem.CreatedAt = time.Now().UTC()
		orderItem.DiscountAmount = itemFromTotals.DiscountAmount
		// orderItem.DiscountInvoiced
		orderItem.DiscountPercent = itemFromTotals.DiscountPercent
		// orderItem.FreeShipping
		// orderItem.IsQtyDecimal
		// orderItem.IsVirtual
		orderItem.ItemId = itemFromTotals.ItemId
		orderItem.Name = itemFromTotals.Name
		// orderItem.NoDiscount
		// orderItem.OrderId  ToDo Assign orderId to items after order has been placed to db  --> Done
		// orderItem.OriginalPrice
		// orderItem.ParentItemId
		// orderItem.Price
		// orderItem.ProductId
		// orderItem.ProductType
		// orderItem.QtyCanceled
		// orderItem.QtyInvoiced
		orderItem.QtyOrdered = itemFromTotals.Qty
		// orderItem.QtyRefunded
		// orderItem.QtyShipped
		orderItem.QuoteItemId = quoteId
		// orderItem.RowInvoiced
		orderItem.RowTotal = itemFromTotals.RowTotal
		// orderItem.RowWeight
		orderItem.Sku = itemFromTotals.SKU
		// orderItem.StoreId
		orderItem.TaxAmount = itemFromTotals.TaxAmount
		orderItem.TaxPercent = itemFromTotals.TaxPercent
		// orderItem.TaxInvoiced
		orderItem.UpdatedAt = time.Now().UTC()
		// orderItem.Weight
		orderItems = append(orderItems, orderItem)
	}
	return
}

// Pure function
func AssignDataToShippingAssignmentsAndSaveIt(data PlaceOrderData, history History, totals total.Totals) (shippingAssignment ShippingAssignment) {
	shippingAssignment.Items = history.Items
	shippingAssignment.Shipping.Address = data.AddressInformation.ShippingAddress
	// Method constructed from Carrier Code and Method Code
	shippingAssignment.Shipping.Method = data.AddressInformation.ShippingCarrierCode + "_" + data.AddressInformation.ShippingMethodCode

	shippingAssignment.Shipping.Total.BaseShippingAmount = totals.BaseShippingAmount
	shippingAssignment.Shipping.Total.BaseShippingDiscountAmount = totals.BaseShippingDiscountAmount
	shippingAssignment.Shipping.Total.BaseShippingInclTax = totals.BaseShippingInclTax
	shippingAssignment.Shipping.Total.BaseShippingTaxAmount = totals.BaseShippingTaxAmount
	shippingAssignment.Shipping.Total.ShippingAmount = totals.ShippingAmount
	shippingAssignment.Shipping.Total.ShippingDiscountAmount = totals.ShippingDiscountAmount
	// shippingAssignment.Shipping.Total.ShippingDiscountTaxCompensationAmount
	shippingAssignment.Shipping.Total.ShippingInclTax = totals.ShippingInclTax
	shippingAssignment.Shipping.Total.ShippingTaxAmount = totals.ShippingTaxAmount
	return
}

func FormatOrderHistory(totals total.Totals, customerData user.CustomerData, billingAddressId int64, quoteId int64) (history History) {
	history.BaseCurrencyCode = totals.BaseCurrencyCode
	history.BaseDiscountAmount = totals.BaseDiscountAmount
	history.BaseGrandTotal = totals.BaseGrandTotal
	history.BaseShippingAmount = totals.BaseShippingAmount
	history.BaseShippingDiscountAmount = totals.BaseShippingDiscountAmount
	history.BaseShippingInclTax = totals.BaseShippingInclTax
	history.BaseShippingTaxAmount = totals.BaseShippingTaxAmount
	history.BaseSubtotal = totals.BaseSubtotal
	history.BaseTaxAmount = totals.BaseTaxAmount
	// orderHistory.BaseSubtotalInclTax
	// orderHistory.BaseTotalDue
	// orderHistory.BaseToGlobalRate
	// orderHistory.BaseToOrderRate
	history.BillingAddressId = billingAddressId
	history.CreatedAt = time.Now().UTC()
	history.UpdatedAt = time.Now().UTC()

	// Customer data
	history.CustomerEmail = customerData.Email
	history.CustomerFirstname = customerData.FirstName
	history.CustomerGroupId = customerData.GroupID
	history.CustomerId = customerData.ID
	if customerData.ID != 0 {
		history.CustomerIsGuest = 0 // ToDo Orders for Guest users needs to be implemented
	} else {
		history.CustomerIsGuest = 1
	}
	history.CustomerLastname = customerData.LastName
	// orderHistory.CustomerNoteNotify
	history.DiscountAmount = totals.DiscountAmount
	history.EmailSent = 1 // ToDo Email service implementation needed
	history.EntityId = quoteId
	history.GlobalCurrencyCode = totals.BaseCurrencyCode // ToDo probably will need to change some time later
	history.GrandTotal = totals.BaseGrandTotal
	// orderHistory.DiscountTaxCompensationAmount
	// orderHistory.IncrementId
	// orderHistory.IsVirtual ToDo virtual products needs some kind of identification
	// orderHistory.OrderCurrencyCode
	// orderHistory.ProtectCode
	history.QuoteId = quoteId
	history.ShippingAmount = totals.ShippingAmount
	// orderHistory.ShippingDescription
	history.ShippingDiscountAmount = totals.ShippingDiscountAmount
	// orderHistory.ShippingDiscountTaxCompensationAmount
	history.ShippingInclTax = totals.ShippingInclTax
	history.ShippingTaxAmount = totals.ShippingTaxAmount
	history.State = "new"
	history.Status = "pending"
	// orderHistory.StoreCurrencyCode
	// orderHistory.StoreId
	// orderHistory.StoreName
	// orderHistory.StoreToBaseRate
	// orderHistory.StoreToOrderRate
	history.Subtotal = totals.Subtotal
	history.SubtotalInclTax = totals.SubtotalInclTax
	history.TaxAmount = totals.TaxAmount
	// orderHistory.TotalDue
	history.TotalItemCount = totals.ItemsQty
	history.TotalQtyOrdered = totals.ItemsQty
	// orderHistory.Weight
	return
}

func FormatPaymentData(history History, method payment_methods.Method, quoteId int64, userId int) (payment Payment) {
	payment.AdditionalInformation = []string{method.Title}
	payment.AmountOrdered = history.BaseGrandTotal
	payment.BaseAmountOrdered = history.BaseGrandTotal
	payment.ShippingAmount = history.ShippingAmount
	payment.EntityId = quoteId
	payment.Method = method.Code
	payment.ParentId = userId
	payment.ShippingAmount = history.ShippingAmount
	// orderPayment.AccountStatus
	// orderPayment.CcLast4
	return
}


