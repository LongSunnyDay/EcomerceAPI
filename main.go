package main

import (
	"encoding/json"
	"fmt"
	fr "github.com/DATA-DOG/fastroute"
	"jec-product-api/helpers"
	"net/http"
	"strconv"
)

type Product struct {
	Id    int     `json:"id"`
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var checkErr = helpers.CheckErr
var productList []Product
var addr string
var version string
var author string
var routes map[string]fr.Router
var router http.Handler

func init() {
	version = "0.1"
	author = "Marius"
	addr = ":8080"
	initProducts()

	routes = initRoutes()
	router = fr.RouterFunc(func(req *http.Request) http.Handler {
		return routes[req.Method] // fastroute.Router is also http.Handler
	})
}

func initProducts() {
	productList = make([]Product, 0)
	productList = append(productList, Product{1, "s-1", "test product no 1.", 23.99})
	productList = append(productList, Product{2, "s-2", "test product no 1.", 24.99})
	productList = append(productList, Product{3, "s-3", "test product no 1.", 25.99})
	productList = append(productList, Product{4, "s-4", "test product no 1.", 26.99})
}

func initRoutes() map[string]fr.Router {

	routes = map[string]fr.Router{
		"GET": fr.Chain(
			fr.New("/api/product/list", ApiProductListHandler),
			fr.New("/api/product/:id", ApiProductHandler),
			fr.New("/about", AboutHandler),
		),
		"POST": fr.Chain(),
	}
	return routes
}

func getProductById(productId int) *Product {
	for _, p := range productList {
		if p.Id == productId {
			return &p
		}
	}
	return nil
}

func main() {
	fmt.Printf("Server was started on http://localhost" + addr)
	http.ListenAndServe(addr, router)
}

func GetIntParameter(r *http.Request, name string) (int, error) {
	parameter := fr.Parameters(r).ByName(name)
	return strconv.Atoi(parameter)
}

func ApiProductHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := GetIntParameter(r, "id")
	checkErr(err)

	product := getProductById(productId)
	result, err := json.Marshal(product)
	checkErr(err)
	w.Write(result)
}

func ApiProductListHandler(w http.ResponseWriter, req *http.Request) {
	result, err := json.Marshal(productList)
	checkErr(err)
	w.Write(result)
}

func AboutHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "This is a Jiva Labs minimalistic Product Service implementation. "+
		"ver.: %s created by %s", version, author)
}
