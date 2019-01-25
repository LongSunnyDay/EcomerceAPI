package order

import (
	"encoding/json"
	"go-api-ws/auth"
	"go-api-ws/cart"
	"go-api-ws/helpers"
	"go-api-ws/payment_methods"
	"go-api-ws/total"
	"go-api-ws/user"
	"net/http"
	"strconv"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var (
		orderData        PlaceOrderData
		customerData     user.CustomerData
		userId           int
		userIdInt64      int64
		orderTotals      total.Totals
		addressForTotals total.AddressData
	)
	err := json.NewDecoder(r.Body).Decode(&orderData)
	helpers.PanicErr(err)

	// Gets user cart from mongoDb by userId
	cartFromMongo := cart.GetUserCartFromMongoByID(orderData.CartId)

	// Does check if items send with request match items in user cart
	err = CheckIfItemsMatchInCart(cartFromMongo, orderData)
	helpers.PanicErr(err)

	// Performs stock check
	err = CheckStockStatus(orderData)
	helpers.PanicErr(err)

	// Performs FinalPrice check between items send and items in SSOT
	err = FinalPriceCheck(orderData)
	helpers.PanicErr(err)

	// Get customer data from MySql by id send in request
	if len(orderData.UserId) > 0 {
		userId, err = strconv.Atoi(orderData.UserId)
		helpers.PanicErr(err)
		userIdInt64 = int64(userId)
		customerData = user.GetUserFromMySQLById(userIdInt64)
	} else {
		customerData = user.CustomerData{
			Email:     orderData.PersonalData.Email,
			FirstName: orderData.PersonalData.Firstname,
			LastName:  orderData.PersonalData.Lastname,
			GroupID:   1}
		userId = 0
		userIdInt64 = 0
	}

	// Saves billing addresses to MySQL
	billingAddress := AssignDataToBillingAddressAndSaveIt(orderData)

	// Address saved to DB
	billingAddress.InsertOrUpdateAddressIntoMySQL(userIdInt64)

	// Calculates order totals using cart from MongoDB

	addressForTotals.AddressInformation.ShippingCarrierCode = orderData.AddressInformation.ShippingCarrierCode
	addressForTotals.AddressInformation.ShippingMethodCode = orderData.AddressInformation.ShippingMethodCode
	orderTotals.CalculateTotals(orderData.CartId, addressForTotals, customerData.GroupID)

	// Preparing order history
	orderHistory := FormatOrderHistory(orderTotals, customerData, billingAddress.ID, cartFromMongo.QuoteId)

	// Working on Items in Order history []items
	orderItems := FormatOrderHistoryItems(orderTotals, cartFromMongo.QuoteId)
	orderHistory.Items = orderItems

	// Working on Payment information
	paymentMethod := payment_methods.GetPaymentMethodFromDbByMethodCode(orderData.AddressInformation.PaymentMethodCode)
	orderPayment := FormatPaymentData(orderHistory, paymentMethod, cartFromMongo.QuoteId, userId)

	// Shipping assignment
	shippingAssignment := AssignDataToShippingAssignmentsAndSaveIt(orderData, orderHistory, orderTotals)

	// Saves order to MySQL
	orderHistory.SaveOrder()

	// Assigns OrderId to order items
	// Saves them to MySQL
	for i := 0; i < len(orderHistory.Items); i++ {
		orderHistory.Items[i].OrderId = orderHistory.ID
		orderHistory.Items[i].SaveItem()
	}

	// Saves payment_methods data to MySQL
	orderPayment.OrderId = orderHistory.ID
	orderPayment.SavePaymentData(orderHistory.ID)

	// Saves shipping address to MySQL
	shippingAssignment.Shipping.Address.SaveOrderShippingAddress(orderHistory.ID)

	// Assign data to order history
	orderHistory.BillingAddress = BillingAddress(
		Address{
			Id:          billingAddress.ID,
			CustomerId:  billingAddress.CustomerID,
			City:        billingAddress.City,
			Company:     billingAddress.Company,
			CountryId:   billingAddress.CountryID,
			Email:       billingAddress.Email,
			Firstname:   billingAddress.Firstname,
			Lastname:    billingAddress.Lastname,
			Postcode:    billingAddress.Postcode,
			Region:      billingAddress.Region.Region,
			RegionCode:  billingAddress.Region.RegionCode,
			RegionId:    billingAddress.Region.RegionID,
			Street:      billingAddress.Street,
			StreetLine0: billingAddress.StreetLine0,
			StreetLine1: billingAddress.StreetLine1,
			Telephone:   billingAddress.Telephone,
			OrderId:     orderHistory.ID})
	orderHistory.Payment = orderPayment
	shippingAssignment.Items = orderHistory.Items
	orderHistory.ExtensionAttributes.ShippingAssignments = append(orderHistory.ExtensionAttributes.ShippingAssignments, shippingAssignment)

	// Writes order to json file
	orderHistory.WriteToJsonFile()

	// Order Pickup
	orderPickup := orderHistory.BuildOrderPickupForm()
	orderPickup.MakeOrderPickup()

	// Changes cart status to "Inactive" in mongoDb
	defer cart.UpdateCartStatus(cartFromMongo.QuoteId)

	resp := map[string]int{
		"code": 200}

	helpers.WriteResultWithStatusCode(w, resp, http.StatusOK)
}

func GetCustomerOrderHistory(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	token := auth.ParseToken(urlToken)
	claims, err := auth.GetTokenClaims(token)
	helpers.CheckErr(err)
	if err != nil {
		helpers.WriteResultWithStatusCode(w, err, http.StatusBadRequest)
	} else {
		subInt, err := strconv.Atoi(claims["sub"].(string))
		helpers.PanicErr(err)
		orderHistory := GetAllCustomerOrderHistory(subInt)

		result := result{
			Items:      orderHistory,
			TotalCount: len(orderHistory),
		}
		response := response{
			Code:   http.StatusOK,
			Result: result,
		}
		helpers.WriteResultWithStatusCode(w, response, response.Code)

	}
}
