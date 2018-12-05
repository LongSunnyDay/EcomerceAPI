package user

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-api-ws/auth"
	"go-api-ws/cart"
	"go-api-ws/config"
	"go-api-ws/core"
	"go-api-ws/helpers"
	"net/http"
	"strconv"
	"time"
)

var userModule core.ApiModule

func init() {
	userModule = core.ApiModule{
		Name:        "User module",
		Description: "User module. Supports username and email authentication. Categories are stored as a flat list.",
		Version:     "0.1",
		Author:      "Matas Cereskevicius @ JivaLabs",
	}

}

const adminRole = "admin"
const userRole = "user"

// Get Order History
// Path: /api/user/order-history
func getOrderHistory(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	token, _ := auth.ParseToken(urlToken)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			orderHistory := getUserOrderHistoryFromMongo(claims["sub"].(string))
			response := helpers.Response{
				Code:   http.StatusOK,
				Result: orderHistory}
			response.SendResponse(w)
		}
	} else {
		helpers.WriteResultWithStatusCode(w, "Invalid token", http.StatusForbidden)
	}
}

// Me endpoint
// Path /api/user/me
func meEndpoint(w http.ResponseWriter, r *http.Request) {
	urlToken, err := helpers.GetTokenFromUrl(r)
	helpers.PanicErr(err)
	token, _ := auth.ParseToken(urlToken)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			userId, err := strconv.Atoi(claims["sub"].(string))
			helpers.PanicErr(err)
			userId64 := int64(userId)
			userInfo := getUserFromMySQLById(userId64)
			userInfo.Addresses = getAddressesFromMySQL(userId64)
			response := helpers.Response{
				Code:   http.StatusOK,
				Result: userInfo}
			response.SendResponse(w)
		} else {
			helpers.WriteResultWithStatusCode(w, "Token expired", http.StatusForbidden)
		}
	} else {
		helpers.WriteResultWithStatusCode(w, "Invalid token", http.StatusBadRequest)
	}
}

// RegisterUser function
// Path: api/user/create
func registerUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://user/jsonSchemaModels/userRegister.schema.json", user)
	if validationResult.Valid() {
		id := insertUserIntoMySQL(user)
		customer := getUserFromMySQLById(id)
		//userInfo := CustomerData{
		//	Address:              []*Address{},
		//	CreatedAt:              time.Now().Unix(),
		//	CreatedIn:              "Default Store View",
		//	DisableAutoGroupChange: 0,
		//	GroupID:                1,
		//	ID:                     user.ID,
		//	WebsiteID:              1,
		//	UpdatedAt:              time.Now().Unix(),
		//	StoreID:                1,
		//	FirstName:              user.Customer.FirstName,
		//	LastName:               user.Customer.LastName,
		//	Email:                  user.Customer.Email,
		//}
		//insertUserIntoMongo(userInfo)
		cart.CreateCartInMongoDB(user.ID)
		response := helpers.Response{
			Code:   http.StatusOK,
			Result: customer}
		response.SendResponse(w)
	} else {
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)
	}
}

// Path: /api/user/update
//Method: post
func updateUser(w http.ResponseWriter, r *http.Request) {
	var user UpdatedCustomer
	err := json.NewDecoder(r.Body).Decode(&user)
	helpers.PanicErr(err)
	//fmt.Printf("%+v\n",user)
	for i := range user.UpdateUser.Addresses {
		user.UpdateUser.Addresses[i].insertOrUpdateAddressIntoMySQL(user.UpdateUser.ID)
	}
	for _, address := range user.UpdateUser.Addresses {
		if address.DefaultShipping == true {
			user.UpdateUser.DefaultShipping = strconv.Itoa(int(address.ID))
		}
	}
	user.UpdateUser.UpdateUserByIdMySQL()
	response := helpers.Response{
		Result: user.UpdateUser,
		Code:   http.StatusOK}
	response.SendResponse(w)
}

// Path: /api/user/refresh
func refreshToken(w http.ResponseWriter, req *http.Request) {
	var jsonBody map[string]string
	_ = json.NewDecoder(req.Body).Decode(&jsonBody)
	token, _ := auth.ParseToken(jsonBody["refreshToken"])
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {

			groupId := getGroupIdFromMySQLById(claims["sub"].(int))

			//role := roleByGroupId(groupId)

			//authToken := auth.GetNewAuthToken(claims["sub"].(string), role)
			authToken := auth.GetNewAuthToken(claims["sub"].(string), groupId)

			authTokenString, err := authToken.SignedString([]byte(config.MySecret))
			helpers.PanicErr(err)

			refreshToken := auth.GetNewRefreshToken(claims["sub"].(string))
			refreshTokenString, err := refreshToken.SignedString([]byte(config.MySecret))
			helpers.PanicErr(err)

			response := helpers.Response{
				Code:   http.StatusOK,
				Result: authTokenString,
				Meta: map[string]string{
					"refreshToken": refreshTokenString}}
			response.SendResponse(w)

		} else {
			helpers.WriteResultWithStatusCode(w, "Token expired", http.StatusForbidden)
		}
	} else {
		helpers.WriteResultWithStatusCode(w, "Invalid token", http.StatusBadRequest)
	}
}

// Path: /api/user/login
func loginEndpoint(w http.ResponseWriter, req *http.Request) {
	var userLogin LoginForm

	_ = json.NewDecoder(req.Body).Decode(&userLogin)
	validationResult := helpers.CheckJSONSchemaWithGoStruct("file://user/jsonSchemaModels/userLogin.schema.json", userLogin)

	pswd := userLogin.Password
	userLogin.Password = ""

	if validationResult.Valid() {
		userFromDb := getUserFromMySQLByEmail(userLogin.Username)

		if checkPasswordHash(pswd, userFromDb.Password) {

			//role := roleByGroupId(userFromDb.GroupId)

			//authToken := auth.GetNewAuthToken(userFromDb.ID, role)
			authToken := auth.GetNewAuthToken(userFromDb.ID, userFromDb.GroupId)

			authTokenString, err := authToken.SignedString([]byte(config.MySecret))
			helpers.PanicErr(err)

			refreshToken := auth.GetNewRefreshToken(userFromDb.ID)
			refreshTokenString, err := refreshToken.SignedString([]byte(config.MySecret))
			helpers.PanicErr(err)

			response := helpers.Response{
				Code:   http.StatusOK,
				Result: authTokenString,
				Meta: map[string]string{
					"refreshToken": refreshTokenString,
				}}
			response.SendResponse(w)
		} else {
			helpers.WriteResultWithStatusCode(w, "Password is invalid", http.StatusUnauthorized)
		}
	} else {
		helpers.WriteResultWithStatusCode(w, validationResult.Errors(), http.StatusBadRequest)

	}
}
