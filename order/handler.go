package order

import (
	"encoding/json"
	"fmt"
	"go-api-ws/addresses"
	"go-api-ws/cart"
	"go-api-ws/helpers"
	"go-api-ws/payment"
	"go-api-ws/product"
	"go-api-ws/stock"
	"go-api-ws/total"
	"go-api-ws/user"
	"net/http"
	"strconv"
	"time"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderData PlaceOrderData
	err := json.NewDecoder(r.Body).Decode(&orderData)
	helpers.PanicErr(err)

	// Gets user cart from mongoDb by userId
	cartItemsFromMongo := cart.GetUserCartFromMongoByID(orderData.UserId)

	// Does check if items send with request match items in user cart
	if len(cartItemsFromMongo) == len(orderData.Products) {
		for i, item := range orderData.Products {
			if cartItemsFromMongo[i].SKU != item.Sku || cartItemsFromMongo[i].QTY != item.Qty {
				fmt.Println("Items in order and in cart doesn't match by SKU or QTY")
			} else {
				fmt.Println("All good, order item SKU -> ", item.Sku)
			}
		}
	} else {
		fmt.Println("Items amount in cart and in order is not the same. Cart items -> ", len(cartItemsFromMongo),
			". Order items -> ", len(orderData.Products))
	}

	// Performs stock check
	var orderStock []stock.DataStock
	for _, item := range orderData.Products {
		var SSOTItem stock.DataStock
		SSOTItem.GetDataFromDbBySku(item.Sku)
		orderStock = append(orderStock, SSOTItem)
	}
	for _, item := range orderData.Products {
		for _, stockItem := range orderStock {
			err := stockItem.CheckSOOT(item.Sku, item.Qty)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// Performs FinalPrice check between items send and items in SSOT
	for _, item := range orderData.Products {
		checkPrice := product.GetProductPriceFromDbBySku(item.Sku, item.FinalPrice)
		if !checkPrice {
			fmt.Printf("Product %v price doesn't match with price in db.", item.Name)
		}
	}

	// Get customer data from MySql by id send in request
	userId, err := strconv.Atoi(orderData.UserId)
	helpers.PanicErr(err)
	userIdInt64 := int64(userId)
	customerData := user.GetUserFromMySQLById(userIdInt64)

	// Saves billing addresses to MySQL
	var billingAddress addresses.Address
	billingAddress.RegionID = orderData.AddressInformation.BillingAddress.RegionId
	billingAddress.CountryID = orderData.AddressInformation.BillingAddress.CountryId
	billingAddress.StreetLine0 = orderData.AddressInformation.BillingAddress.Street[0]
	billingAddress.StreetLine1 = orderData.AddressInformation.BillingAddress.Street[1]
	billingAddress.Postcode = orderData.AddressInformation.BillingAddress.Postcode
	billingAddress.City = orderData.AddressInformation.BillingAddress.City
	billingAddress.Firstname = orderData.AddressInformation.BillingAddress.Firstname
	billingAddress.Lastname = orderData.AddressInformation.BillingAddress.Lastname
	billingAddress.Email = orderData.AddressInformation.BillingAddress.Email
	billingAddress.DefaultBilling = true
	// Address saved to DB
	billingAddress.InsertOrUpdateAddressIntoMySQL(userIdInt64)

	// Calculates order totals using cart from MongoDB
	var orderTotals total.Totals
	var addressForTotals total.AddressData
	addressForTotals.AddressInformation.ShippingCarrierCode = orderData.AddressInformation.ShippingCarrierCode
	addressForTotals.AddressInformation.ShippingMethodCode = orderData.AddressInformation.ShippingMethodCode
	orderTotals.CalculateTotals(orderData.CartId, addressForTotals, customerData.GroupID)

	// Preparing order history
	var orderHistory History
	orderHistory.BaseCurrencyCode = orderTotals.BaseCurrencyCode
	orderHistory.BaseDiscountAmount = orderTotals.BaseDiscountAmount
	orderHistory.BaseGrandTotal = orderTotals.BaseGrandTotal
	orderHistory.BaseShippingAmount = orderTotals.BaseShippingAmount
	orderHistory.BaseShippingDiscountAmount = orderTotals.BaseShippingDiscountAmount
	orderHistory.BaseShippingInclTax = orderTotals.BaseShippingInclTax
	orderHistory.BaseShippingTaxAmount = orderTotals.BaseShippingTaxAmount
	orderHistory.BaseSubtotal = orderTotals.BaseSubtotal
	orderHistory.BaseTaxAmount = orderTotals.BaseTaxAmount
	// orderHistory.BaseSubtotalInclTax
	// orderHistory.BaseTotalDue
	// orderHistory.BaseToGlobalRate
	// orderHistory.BaseToOrderRate
	orderHistory.BillingAddressId = billingAddress.ID
	orderHistory.CreatedAt = time.Now().UTC()
	orderHistory.UpdatedAt = time.Now().UTC()

	// Customer data
	orderHistory.CustomerEmail = customerData.Email
	orderHistory.CustomerFirstname = customerData.FirstName
	orderHistory.CustomerGroupId = customerData.GroupID
	orderHistory.CustomerId = customerData.ID
	orderHistory.CustomerIsGuest = 0 // ToDo Orders for Guest users needs to be implemented
	orderHistory.CustomerLastname = customerData.LastName
	// orderHistory.CustomerNoteNotify
	orderHistory.DiscountAmount = orderTotals.DiscountAmount
	orderHistory.EmailSent = 1 // ToDo Email service implementation needed
	orderHistory.EntityId = customerData.ID
	orderHistory.GlobalCurrencyCode = orderTotals.BaseCurrencyCode // ToDo probably will need to change some time later
	orderHistory.GrandTotal = orderTotals.GrandTotal
	// orderHistory.DiscountTaxCompensationAmount
	// orderHistory.IncrementId
	// orderHistory.IsVirtual ToDo virtual products needs some kind of identification
	// orderHistory.OrderCurrencyCode
	// orderHistory.ProtectCode
	orderHistory.QuoteId = cartItemsFromMongo[0].QuoteId
	orderHistory.ShippingAmount = orderTotals.ShippingAmount
	// orderHistory.ShippingDescription
	orderHistory.ShippingDiscountAmount = orderTotals.ShippingDiscountAmount
	// orderHistory.ShippingDiscountTaxCompensationAmount
	orderHistory.ShippingInclTax = orderTotals.ShippingInclTax
	orderHistory.ShippingTaxAmount = orderTotals.ShippingTaxAmount
	orderHistory.State = "new"
	orderHistory.Status = "pending"
	// orderHistory.StoreCurrencyCode
	// orderHistory.StoreId
	// orderHistory.StoreName
	// orderHistory.StoreToBaseRate
	// orderHistory.StoreToOrderRate
	orderHistory.Subtotal = orderTotals.Subtotal
	orderHistory.SubtotalInclTax = orderTotals.SubtotalInclTax
	orderHistory.TaxAmount = orderTotals.TaxAmount
	// orderHistory.TotalDue
	orderHistory.TotalItemCount = orderTotals.ItemsQty
	orderHistory.TotalQtyOrdered = orderTotals.ItemsQty
	// orderHistory.Weight

	// Working on Items in Order history []items
	var orderItems []Item
	for _, itemFromTotals := range orderTotals.Items {
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
		// orderItem.OrderId  ToDo Assign orderId to items after order has been placed to db
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
		// orderItem.QuoteItemId
		// orderItem.RowInvoiced
		orderItem.RowTotal = itemFromTotals.RowTotal
		// orderItem.RowWeight
		// orderItem.Sku ToDo Assign Sku from MySQL or Cart service
		// orderItem.StoreId
		orderItem.TaxAmount = itemFromTotals.TaxAmount
		orderItem.TaxPercent = itemFromTotals.TaxPercent
		// orderItem.TaxInvoiced
		orderItem.UpdatedAt = time.Now().UTC()
		// orderItem.Weight
		orderItems = append(orderItems, orderItem)
	}
	orderHistory.Items = orderItems

	// Working on Payment information
	paymentMethod := payment.GetPaymentMethodFromDbByMethodCode(orderData.AddressInformation.PaymentMethodCode)
	var orderPayment Payment
	orderPayment.AdditionalInformation = []string{paymentMethod.Title}
	orderPayment.AmountOrdered = orderHistory.BaseGrandTotal
	orderPayment.BaseAmountOrdered = orderHistory.BaseGrandTotal
	orderPayment.ShippingAmount = orderHistory.ShippingAmount
	orderPayment.EntityId = userId
	orderPayment.Method = paymentMethod.Code
	orderPayment.ParentId = userId
	orderPayment.ShippingAmount = orderHistory.ShippingAmount
	// orderPayment.AccountStatus
	// orderPayment.CcLast4

	// Shipping assignment
	var shippingAssignment ShippingAssignment
	shippingAssignment.Items = orderHistory.Items
	shippingAssignment.Shipping.Address = orderData.AddressInformation.ShippingAddress
	// Method constructed from Carrier Code and Method Code
	shippingAssignment.Shipping.Method = orderData.AddressInformation.ShippingCarrierCode + "_" + orderData.AddressInformation.ShippingMethodCode

	shippingAssignment.Shipping.Total.BaseShippingAmount = orderTotals.BaseShippingAmount
	shippingAssignment.Shipping.Total.BaseShippingDiscountAmount = orderTotals.BaseShippingDiscountAmount
	shippingAssignment.Shipping.Total.BaseShippingInclTax = orderTotals.BaseShippingInclTax
	shippingAssignment.Shipping.Total.BaseShippingTaxAmount = orderTotals.BaseShippingTaxAmount
	shippingAssignment.Shipping.Total.ShippingAmount = orderTotals.ShippingAmount
	shippingAssignment.Shipping.Total.ShippingDiscountAmount = orderTotals.ShippingDiscountAmount
	// shippingAssignment.Shipping.Total.ShippingDiscountTaxCompensationAmount
	shippingAssignment.Shipping.Total.ShippingInclTax = orderTotals.ShippingInclTax
	shippingAssignment.Shipping.Total.ShippingTaxAmount = orderTotals.ShippingTaxAmount

	// Saves order to MySQL
	orderHistory.SaveOrder()

	// Assigns OrderId to order items
	// Saves them to MySQL
	for i := 0; i < len(orderHistory.Items); i++ {
		orderHistory.Items[i].OrderId = orderHistory.ID
		orderHistory.Items[i].SaveItem()
	}

	// Saves payment data to MySQL
	orderPayment.SavePaymentData(orderHistory.ID)



}
