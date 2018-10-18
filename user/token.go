package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
	"net/http"
	m "./models"
	"../helpers"
	)

func LoginEndpoint (w http.ResponseWriter, req *http.Request) {
	var user m.User
	_ = json.NewDecoder(req.Body).Decode(&user)





	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer": user.Customer,
		"password": user.Password,
	})
	fmt.Println(user)
	fmt.Println(token)
	tokenString, err := token.SignedString([]byte("secret"))
	helpers.CheckErr(err)
	json.NewEncoder(w).Encode(tokenString)
}

func ProtectedEndpoint(w http.ResponseWriter, req *http.Request)  {
	params := req.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There is an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user m.User
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		json.NewEncoder(w).Encode("Invalid authorization token")
	}
}
