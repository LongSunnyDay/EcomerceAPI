package order

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/auth"
	"go-api-ws/cart"
	"go-api-ws/helpers"
	"go-api-ws/payment"
	"go-api-ws/total"
	"go-api-ws/user"
	"net/http"
	"strconv"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderData PlaceOrderData
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
	userId, err := strconv.Atoi(orderData.UserId)
	helpers.PanicErr(err)
	userIdInt64 := int64(userId)
	customerData := user.GetUserFromMySQLById(userIdInt64)

	// Saves billing addresses to MySQL
	billingAddress := AssignDataToBillingAddressAndSaveIt(orderData)

	// Address saved to DB
	billingAddress.InsertOrUpdateAddressIntoMySQL(userIdInt64)

	// Calculates order totals using cart from MongoDB
	var orderTotals total.Totals
	var addressForTotals total.AddressData
	addressForTotals.AddressInformation.ShippingCarrierCode = orderData.AddressInformation.ShippingCarrierCode
	addressForTotals.AddressInformation.ShippingMethodCode = orderData.AddressInformation.ShippingMethodCode
	orderTotals.CalculateTotals(orderData.CartId, addressForTotals, customerData.GroupID)

	// Preparing order history
	orderHistory := FormatOrderHistory(orderTotals, customerData, billingAddress.ID, cartFromMongo.QuoteId)

	// Working on Items in Order history []items
	orderItems := FormatOrderHistoryItems(orderTotals, cartFromMongo.QuoteId)
	orderHistory.Items = orderItems

	// Working on Payment information
	paymentMethod := payment.GetPaymentMethodFromDbByMethodCode(orderData.AddressInformation.PaymentMethodCode)
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

	// Saves payment data to MySQL
	orderPayment.OrderId = orderHistory.ID
	orderPayment.SavePaymentData(orderHistory.ID)

	// Saves shipping address to MySQL
	shippingAssignment.Shipping.Address.SaveOrderShippingAddress(orderHistory.ID)
	helpers.WriteResultWithStatusCode(w, http.StatusOK, http.StatusOK)
}

func GetCustomerOrderHistory(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	token, err := auth.ParseToken(urlToken)
	helpers.PanicErr(err)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
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
