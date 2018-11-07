package cart

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"go-api-ws/user"
	"net/http"
	"time"
)

func CartRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/create", createCart)

	return r
}

func createCart(w http.ResponseWriter, req *http.Request) {
	urlToken := req.URL.Query()["token"][0]
	if len(urlToken) > 0 {
		token, _ := jwt.Parse(urlToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There is and error")
			}
			return []byte(user.MySecret), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims.VerifyExpiresAt(time.Now().Unix(), true) {
				items := []string{}
				result := map[string]interface{}{
					"items": items,
				}
				response := user.Response{
					Code:   200,
					Result: result}
				w.Header().Set("content-type", "application/json")
				w.WriteHeader(200)
				json.NewEncoder(w).Encode(response)
			}
		}
	} else {
		result := "a1a1a1a1a1a1"
		response := user.Response{
			Code:   200,
			Result: result}
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
}

func pullCart(w http.ResponseWriter, req *http.Request)  {
	urlToken := req.URL.Query()["token"][0]
	urlCartId := req.URL.Query()["cartId"][0]
	fmt.Println(urlToken, urlCartId)
}