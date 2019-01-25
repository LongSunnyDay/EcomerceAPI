package order

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api-ws/addresses"
	"go-api-ws/cart"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"go-api-ws/payment_methods"
	"go-api-ws/postNord"
	"go-api-ws/product"
	"go-api-ws/stock"
	"go-api-ws/total"
	"go-api-ws/user"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// ToDo A lot of data assignments are skipped because they are not used right now. They probably will need implementation in future.

// Pure function
func CheckIfItemsMatchInCart(cart cart.Cart, data PlaceOrderData) (err error) {
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
	billingAddress.DefaultBilling = false

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

func GetAllCustomerOrderHistory(customerId int) (customerHistoryArray []History) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT * FROM `order` WHERE customer_id = ?", customerId)
	helpers.PanicErr(err)
	for rows.Next() {
		var customerHistory History
		if err := rows.Scan(
			&customerHistory.AppliedRuleIds,
			&customerHistory.BaseCurrencyCode,
			&customerHistory.BaseDiscountAmount,
			&customerHistory.BaseGrandTotal,
			&customerHistory.BaseDiscountTaxCompensationAmount,
			&customerHistory.BaseShippingAmount,
			&customerHistory.BaseShippingDiscountAmount,
			&customerHistory.BaseShippingInclTax,
			&customerHistory.BaseShippingTaxAmount,
			&customerHistory.BaseSubtotal,
			&customerHistory.BaseSubtotalInclTax,
			&customerHistory.BaseTaxAmount,
			&customerHistory.BaseTotalDue,
			&customerHistory.BaseToGlobalRate,
			&customerHistory.BaseToOrderRate,
			&customerHistory.BillingAddressId,
			&customerHistory.CreatedAt,
			&customerHistory.CustomerEmail,
			&customerHistory.CustomerFirstname,
			&customerHistory.CustomerGroupId,
			&customerHistory.CustomerId,
			&customerHistory.CustomerIsGuest,
			&customerHistory.CustomerLastname,
			&customerHistory.CustomerNoteNotify,
			&customerHistory.DiscountAmount,
			&customerHistory.EmailSent,
			&customerHistory.EntityId,
			&customerHistory.GlobalCurrencyCode,
			&customerHistory.GrandTotal,
			&customerHistory.DiscountTaxCompensationAmount,
			&customerHistory.IncrementId,
			&customerHistory.IsVirtual,
			&customerHistory.OrderCurrencyCode,
			&customerHistory.ProtectCode,
			&customerHistory.QuoteId,
			&customerHistory.ShippingAmount,
			&customerHistory.ShippingDescription,
			&customerHistory.ShippingDiscountAmount,
			&customerHistory.ShippingDiscountTaxCompensationAmount,
			&customerHistory.ShippingInclTax,
			&customerHistory.ShippingTaxAmount,
			&customerHistory.State,
			&customerHistory.Status,
			&customerHistory.StoreCurrencyCode,
			&customerHistory.StoreId,
			&customerHistory.StoreName,
			&customerHistory.StoreToBaseRate,
			&customerHistory.StoreToOrderRate,
			&customerHistory.Subtotal,
			&customerHistory.SubtotalInclTax,
			&customerHistory.TaxAmount,
			&customerHistory.TotalDue,
			&customerHistory.TotalItemCount,
			&customerHistory.TotalQtyOrdered,
			&customerHistory.UpdatedAt,
			&customerHistory.Weight,
			&customerHistory.ID); err != nil {
			helpers.PanicErr(err)
		}
		customerHistory.GetOrderItems()
		customerHistory.GetOrderPaymentData()
		customerHistory.GetOrderBillingAddress()
		customerHistory.GetOrderShippingAddress()
		customerHistoryArray = append(customerHistoryArray, customerHistory)
	}
	return
}

func GetOrder(orderId int) (order History) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM `order` WHERE id = ?", orderId).Scan(&order.ID,
		&order.AppliedRuleIds,
		&order.BaseCurrencyCode,
		&order.BaseDiscountAmount,
		&order.BaseGrandTotal,
		&order.BaseDiscountTaxCompensationAmount,
		&order.BaseShippingAmount,
		&order.BaseShippingDiscountAmount,
		&order.BaseShippingInclTax,
		&order.BaseShippingTaxAmount,
		&order.BaseSubtotal,
		&order.BaseSubtotalInclTax,
		&order.BaseTaxAmount,
		&order.BaseTotalDue,
		&order.BaseToGlobalRate,
		&order.BaseToOrderRate,
		&order.BillingAddressId,
		&order.CreatedAt,
		&order.CustomerEmail,
		&order.CustomerFirstname,
		&order.CustomerGroupId,
		&order.CustomerId,
		&order.CustomerIsGuest,
		&order.CustomerLastname,
		&order.CustomerNoteNotify,
		&order.DiscountAmount,
		&order.EmailSent,
		&order.GlobalCurrencyCode,
		&order.GrandTotal,
		&order.DiscountTaxCompensationAmount,
		&order.IncrementId,
		&order.IsVirtual,
		&order.OrderCurrencyCode,
		&order.ProtectCode,
		&order.QuoteId,
		&order.ShippingAmount,
		&order.ShippingDescription,
		&order.ShippingDiscountAmount,
		&order.ShippingDiscountTaxCompensationAmount,
		&order.ShippingInclTax,
		&order.ShippingTaxAmount,
		&order.State,
		&order.Status,
		&order.StoreCurrencyCode,
		&order.StoreId,
		&order.StoreName,
		&order.StoreToBaseRate,
		&order.StoreToOrderRate,
		&order.Subtotal,
		&order.SubtotalInclTax,
		&order.TaxAmount,
		&order.TotalDue,
		&order.TotalItemCount,
		&order.TotalQtyOrdered,
		&order.UpdatedAt,
		&order.Weight)
	helpers.PanicErr(err)
	return
}

func RemoveOrder(orderId int) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("DELETE FROM `order` WHERE id = ?", orderId)
	helpers.PanicErr(err)
}

func (order History) UpdateOrder() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	order.UpdatedAt = time.Now().UTC()
	res, err := db.Exec("UPDATE order o SET "+
		"o.applied_rule_ids = ?, "+
		"o.base_currency_code  = ?, "+
		"o.base_discount_amount = ?, "+
		"o.base_grand_total = ?, "+
		"o.base_discount_tax_compensation_amount = ?, "+
		"o.base_shipping_amount = ?, "+
		"o.base_shipping_discount_amount = ?, "+
		"o.base_shipping_incl_tax = ?, "+
		"o.base_shipping_tax_amount = ?, "+
		"o.base_subtotal = ?, "+
		"o.base_subtotal_incl_tax = ?, "+
		"o.base_tax_amount = ?, "+
		"o.base_total_due = ?, "+
		"o.base_to_global_rate = ?, "+
		"o.base_to_order_rate = ?, "+
		"o.billing_address_id = ?, "+
		"o.created_at = ?, "+
		"o.customer_email = ?, "+
		"o.customer_firstname = ?, "+
		"o.customer_group_id = ?, "+
		"o.customer_id = ?, "+
		"o.customer_is_guest = ?, "+
		"o.customer_lastname = ?, "+
		"o.customer_note_notify = ?, "+
		"o.discount_amount = ?, "+
		"o.email_sent = ?, "+
		"o.entity_id = ?, "+
		"o.global_currency_code = ?, "+
		"o.grand_total = ?, "+
		"o.discount_tax_compensation_amount = ?, "+
		"o.increment_id = ?, "+
		"o.is_virtual = ?, "+
		"o.order_currency_code = ?, "+
		"o.protect_code = ?, "+
		"o.quote_id = ?, "+
		"o.shipping_amount = ?, "+
		"o.shipping_description = ?, "+
		"o.shipping_discount_amount = ?, "+
		"o.shipping_discount_tax_compensation_amount = ?, "+
		"o.shipping_incl_tax = ?, "+
		"o.shipping_tax_amount = ?, "+
		"o.state = ?, "+
		"o.status = ?, "+
		"o.store_currency_code = ?, "+
		"o.store_id = ?, "+
		"o.store_name = ?, "+
		"o.store_to_base_rate = ?, "+
		"o.store_to_order_rate = ?, "+
		"o.subtotal = ?, "+
		"o.subtotal_incl_tax = ?, "+
		"o.tax_amount = ?, "+
		"o.total_due = ?, "+
		"o.total_item_count = ?, "+
		"o.total_qty_ordered = ?, "+
		"o.updated_at = ?, "+
		"o.weight = ? "+
		"WHERE o.id = ?",
		order.AppliedRuleIds,
		order.BaseCurrencyCode,
		order.BaseDiscountAmount,
		order.BaseGrandTotal,
		order.BaseDiscountTaxCompensationAmount,
		order.BaseShippingAmount,
		order.BaseShippingDiscountAmount,
		order.BaseShippingInclTax,
		order.BaseShippingTaxAmount,
		order.BaseSubtotal,
		order.BaseSubtotalInclTax,
		order.BaseTaxAmount,
		order.BaseTotalDue,
		order.BaseToGlobalRate,
		order.BaseToOrderRate,
		order.BillingAddressId,
		order.CreatedAt,
		order.CustomerEmail,
		order.CustomerFirstname,
		order.CustomerGroupId,
		order.CustomerId,
		order.CustomerIsGuest,
		order.CustomerLastname,
		order.CustomerNoteNotify,
		order.DiscountAmount,
		order.EmailSent,
		order.EntityId,
		order.GlobalCurrencyCode,
		order.GrandTotal,
		order.DiscountTaxCompensationAmount,
		order.IncrementId,
		order.IsVirtual,
		order.OrderCurrencyCode,
		order.ProtectCode,
		order.QuoteId,
		order.ShippingAmount,
		order.ShippingDescription,
		order.ShippingDiscountAmount,
		order.ShippingDiscountTaxCompensationAmount,
		order.ShippingInclTax,
		order.ShippingTaxAmount,
		order.State,
		order.Status,
		order.StoreCurrencyCode,
		order.StoreId,
		order.StoreName,
		order.StoreToBaseRate,
		order.StoreToOrderRate,
		order.Subtotal,
		order.SubtotalInclTax,
		order.TaxAmount,
		order.TotalDue,
		order.TotalItemCount,
		order.TotalQtyOrdered,
		order.UpdatedAt,
		order.Weight,
		order.ID)
	helpers.PanicErr(err)
	rowsAffected, err := res.RowsAffected()
	helpers.PanicErr(err)
	fmt.Println("Order ID: ", order.ID, " got ", rowsAffected, " rows updated")
}

func (item Item) SaveItem() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO order_items ("+
		"amount_refunded, "+
		"applied_rule_ids, "+
		"base_amount_refunded, "+
		"base_discount_amount, "+
		"base_discount_invoiced, "+
		"base_discount_tax_compensation_amount, "+
		"base_original_price, "+
		"base_price, "+
		"base_price_incl_tax, "+
		"base_row_invoiced, "+
		"base_row_total, "+
		"base_row_total_incl_tax, "+
		"base_tax_amount, "+
		"base_tax_invoiced, "+
		"created_at, "+
		"discount_amount, "+
		"discount_invoiced, "+
		"discount_percent, "+
		"free_shipping, "+
		"discount_tax_compensation_amount, "+
		"is_qty_decimal, "+
		"is_virtual, "+
		"name, "+
		"no_discount, "+
		"order_id, "+
		"original_price, "+
		"parent_item_id, "+
		"product_id, "+
		"product_type, "+
		"qty_canceled, "+
		"qty_invoiced, "+
		"qty_ordered, "+
		"qty_refunded, "+
		"qty_shipped, "+
		"quote_item_id, "+
		"row_invoiced, "+
		"row_total, "+
		"row_total_incl_tax, "+
		"row_weight, "+
		"sku, "+
		"store_id, "+
		"tax_amount, "+
		"tax_invoiced, "+
		"tax_percent, "+
		"updated_at, "+
		"weight) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		item.AmountRefunded,
		item.AppliedRuleIds,
		item.BaseAmountRefunded,
		item.BaseDiscountAmount,
		item.BaseDiscountInvoiced,
		item.BaseDiscountTaxCompensationAmount,
		item.BaseOriginalPrice,
		item.BasePrice,
		item.BasePriceInclTax,
		item.BaseRowInvoiced,
		item.BaseRowTotal,
		item.BaseRowTotalInclTax,
		item.BaseTaxAmount,
		item.BaseTaxInvoiced,
		item.CreatedAt,
		item.DiscountAmount,
		item.DiscountInvoiced,
		item.DiscountPercent,
		item.FreeShipping,
		item.BaseDiscountTaxCompensationAmount,
		item.IsQtyDecimal,
		item.IsVirtual,
		item.Name,
		item.NoDiscount,
		item.OrderId,
		item.OriginalPrice,
		item.ParentItemId,
		item.ProductId,
		item.ProductType,
		item.QtyCanceled,
		item.QtyInvoiced,
		item.QtyOrdered,
		item.QtyRefunded,
		item.QtyShipped,
		item.QuoteItemId,
		item.RowInvoiced,
		item.RowTotal,
		item.RowTotalInclTax,
		item.RowWeight,
		item.Sku,
		item.StoreId,
		item.TaxAmount,
		item.TaxInvoiced,
		item.TaxPercent,
		item.UpdatedAt,
		item.Weight)
	helpers.PanicErr(err)
}

func (paymentData Payment) SavePaymentData(orderId int64) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO payment ("+
		"account_status, "+
		"amount_ordered, "+
		"base_amount_ordered, "+
		"base_shipping_amount, "+
		"cc_last4, "+
		"entity_id, "+
		"method, "+
		"parent_id ,"+
		"shipping_amount, "+
		"order_id, "+
		"additional_information) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		paymentData.AccountStatus,
		paymentData.AmountOrdered,
		paymentData.BaseAmountOrdered,
		paymentData.BaseShippingAmount,
		paymentData.CcLast4,
		paymentData.EntityId,
		paymentData.Method,
		paymentData.ParentId,
		paymentData.ShippingAmount,
		orderId,
		paymentData.AdditionalInformation[0])
	helpers.PanicErr(err)

}

func (order *History) GetOrderPaymentData() {
	order.Payment.AdditionalInformation = make([]string, 2)
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM payment WHERE order_id = ?", order.ID).
		Scan(
			&order.Payment.Id,
			&order.Payment.AccountStatus,
			&order.Payment.AmountOrdered,
			&order.Payment.BaseAmountOrdered,
			&order.Payment.BaseShippingAmount,
			&order.Payment.CcLast4,
			&order.Payment.EntityId,
			&order.Payment.Method,
			&order.Payment.ParentId,
			&order.Payment.ShippingAmount,
			&order.Payment.OrderId,
			&order.Payment.AdditionalInformation[0])
	helpers.PanicErr(err)
	//fmt.Println(order.Payment.AdditionalInformation)
}

func (address *Address) SaveOrderShippingAddress(orderId int64) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO order_shipping_address ("+
		"city, "+
		"company, "+
		"country_id, "+
		"email, "+
		"firstname, "+
		"lastname, "+
		"postcode, "+
		"region, "+
		"region_code, "+
		"region_id, "+
		"telephone, "+
		"street_line_0, "+
		"street_line_1, "+
		"address_type, "+
		"entity_id, "+
		"parent_id) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		address.City,
		address.Company,
		address.CountryId,
		address.Email,
		address.Firstname,
		address.Lastname,
		address.Postcode,
		address.Region,
		address.RegionCode,
		address.RegionId,
		address.Telephone,
		address.Street[0],
		address.Street[1],
		"shipping",
		address.EntityId,
		orderId)
	helpers.PanicErr(err)
}

func (order *History) GetOrderItems() {
	order.Items = []Item{}
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", order.ID)
	helpers.PanicErr(err)
	for rows.Next() {
		var item Item
		if err := rows.Scan(
			&item.AmountRefunded,
			&item.AppliedRuleIds,
			&item.BaseAmountRefunded,
			&item.BaseDiscountAmount,
			&item.BaseDiscountInvoiced,
			&item.BaseDiscountTaxCompensationAmount,
			&item.BaseOriginalPrice,
			&item.BasePrice,
			&item.BasePriceInclTax,
			&item.BaseRowInvoiced,
			&item.BaseRowTotal,
			&item.BaseRowTotalInclTax,
			&item.BaseTaxAmount,
			&item.BaseTaxInvoiced,
			&item.CreatedAt,
			&item.DiscountAmount,
			&item.DiscountInvoiced,
			&item.DiscountPercent,
			&item.FreeShipping,
			&item.DiscountTaxCompensationAmount,
			&item.IsQtyDecimal,
			&item.IsVirtual,
			&item.Name,
			&item.NoDiscount,
			&item.OrderId,
			&item.OriginalPrice,
			&item.ParentItemId,
			&item.ProductId,
			&item.ProductType,
			&item.QtyCanceled,
			&item.QtyInvoiced,
			&item.QtyOrdered,
			&item.QtyRefunded,
			&item.QtyShipped,
			&item.QuoteItemId,
			&item.RowInvoiced,
			&item.RowTotal,
			&item.RowTotalInclTax,
			&item.RowWeight,
			&item.Sku,
			&item.StoreId,
			&item.TaxAmount,
			&item.TaxInvoiced,
			&item.TaxPercent,
			&item.UpdatedAt,
			&item.Weight,
			&item.ItemId); err != nil {
			helpers.PanicErr(err)
		}
		order.Items = append(order.Items, item)
	}
}

func (order *History) GetOrderBillingAddress() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM addresses WHERE id = ?", order.BillingAddressId).
		Scan(
			&order.BillingAddress.Id,
			&order.BillingAddress.CustomerId,
			&order.BillingAddress.RegionId,
			&order.BillingAddress.CountryId,
			&order.BillingAddress.Telephone,
			&order.BillingAddress.Postcode,
			&order.BillingAddress.City,
			&order.BillingAddress.Firstname,
			&order.BillingAddress.Lastname,
			&order.BillingAddress.DefaultShipping,
			&order.BillingAddress.StreetLine0,
			&order.BillingAddress.StreetLine1,
			&order.BillingAddress.DefaultBilling,
			&order.BillingAddress.Email)
	helpers.PanicErr(err)
	order.BillingAddress.Street = formatStreet(order.BillingAddress.StreetLine0, order.BillingAddress.StreetLine1)
}

func formatStreet(line0 string, line1 string) []string {
	lines := []string{line0, line1}
	return lines
}

func (order *History) GetOrderShippingAddress() {
	var sa ShippingAssignment
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM order_shipping_address WHERE parent_id = ?", order.ID).
		Scan(&sa.Shipping.Address.Id,
			&sa.Shipping.Address.City,
			&sa.Shipping.Address.Company,
			&sa.Shipping.Address.CountryId,
			&sa.Shipping.Address.Email,
			&sa.Shipping.Address.Firstname,
			&sa.Shipping.Address.Lastname,
			&sa.Shipping.Address.Postcode,
			&sa.Shipping.Address.Region,
			&sa.Shipping.Address.RegionCode,
			&sa.Shipping.Address.RegionId,
			&sa.Shipping.Address.Telephone,
			&sa.Shipping.Address.StreetLine0,
			&sa.Shipping.Address.StreetLine1,
			&sa.Shipping.Address.AddressType,
			&sa.Shipping.Address.EntityId,
			&sa.Shipping.Address.ParentId)
	helpers.PanicErr(err)
	sa.Shipping.Address.Street = formatStreet(sa.Shipping.Address.StreetLine0, sa.Shipping.Address.StreetLine1)
	order.ExtensionAttributes.ShippingAssignments = append(order.ExtensionAttributes.ShippingAssignments, sa)
}

func (order *History) WriteToJsonFile() {

	orderIdInt := int(order.ID)
	orderIdString := strconv.Itoa(orderIdInt) + ".json"

	orderJson, err := json.Marshal(order)
	helpers.PanicErr(err)

	err = ioutil.WriteFile("./order/orders/"+orderIdString, orderJson, os.ModePerm)
	helpers.PanicErr(err)
}

func (order *History) BuildOrderPickupForm() (form postNord.OrderPickupForm) {

	for _, item := range order.Items {
		idInt := int(item.QuoteItemId)
		idString := strconv.Itoa(idInt)
		form.Shipment.Items = append(form.Shipment.Items, &idString)
	}

	orderIdInt := int(order.ID)
	orderIdString := strconv.Itoa(orderIdInt)
	form.Order.OrderReference = orderIdString

	return form
}
func (order *History) SaveOrder() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	res, err := db.Exec("INSERT INTO `order` ("+
		"applied_rule_ids, "+
		"base_currency_code, "+
		"base_discount_amount, "+
		"base_grand_total, "+
		"base_discount_tax_compensation_amount, "+
		"base_shipping_amount, "+
		"base_shipping_discount_amount, "+
		"base_shipping_incl_tax, "+
		"base_shipping_tax_amount, "+
		"base_subtotal, "+
		"base_subtotal_incl_tax, "+
		"base_tax_amount, "+
		"base_total_due, "+
		"base_to_global_rate, "+
		"base_to_order_rate, "+
		"billing_address_id, "+
		"created_at, "+
		"customer_email, "+
		"customer_firstname, "+
		"customer_group_id, "+
		"customer_id, "+
		"customer_is_guest, "+
		"customer_lastname, "+
		"customer_note_notify, "+
		"discount_amount, "+
		"email_sent, "+
		"entity_id, "+
		"global_currency_code, "+
		"grand_total, "+
		"discount_tax_compensation_amount, "+
		"increment_id, "+
		"is_virtual, "+
		"order_currency_code, "+
		"protect_code, "+
		"quote_id, "+
		"shipping_amount, "+
		"shipping_description, "+
		"shipping_discount_amount, "+
		"shipping_discount_tax_compensation_amount, "+
		"shipping_incl_tax, "+
		"shipping_tax_amount, "+
		"state, "+
		"status, "+
		"store_currency_code, "+
		"store_id, "+
		"store_name, "+
		"store_to_base_rate, "+
		"store_to_order_rate, "+
		"subtotal, "+
		"subtotal_incl_tax, "+
		"tax_amount, "+
		"total_due, "+
		"total_item_count, "+
		"total_qty_ordered, "+
		"updated_at, "+
		"weight) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		order.AppliedRuleIds,
		order.BaseCurrencyCode,
		order.BaseDiscountAmount,
		order.BaseGrandTotal,
		order.BaseDiscountTaxCompensationAmount,
		order.BaseShippingAmount,
		order.BaseShippingDiscountAmount,
		order.BaseShippingInclTax,
		order.BaseShippingTaxAmount,
		order.BaseSubtotal,
		order.BaseSubtotalInclTax,
		order.BaseTaxAmount,
		order.BaseTotalDue,
		order.BaseToGlobalRate,
		order.BaseToOrderRate,
		order.BillingAddressId,
		order.CreatedAt,
		order.CustomerEmail,
		order.CustomerFirstname,
		order.CustomerGroupId,
		order.CustomerId,
		order.CustomerIsGuest,
		order.CustomerLastname,
		order.CustomerNoteNotify,
		order.DiscountAmount,
		order.EmailSent,
		order.EntityId,
		order.GlobalCurrencyCode,
		order.GrandTotal,
		order.DiscountTaxCompensationAmount,
		order.IncrementId,
		order.IsVirtual,
		order.OrderCurrencyCode,
		order.ProtectCode,
		order.QuoteId,
		order.ShippingAmount,
		order.ShippingDescription,
		order.ShippingDiscountAmount,
		order.ShippingDiscountTaxCompensationAmount,
		order.ShippingInclTax,
		order.ShippingTaxAmount,
		order.State,
		order.Status,
		order.StoreCurrencyCode,
		order.StoreId,
		order.StoreName,
		order.StoreToBaseRate,
		order.StoreToOrderRate,
		order.Subtotal,
		order.SubtotalInclTax,
		order.TaxAmount,
		order.TotalDue,
		order.TotalItemCount,
		order.TotalQtyOrdered,
		order.UpdatedAt,
		order.Weight)
	helpers.PanicErr(err)
	id, err := res.LastInsertId()
	helpers.PanicErr(err)
	order.ID = id
}
