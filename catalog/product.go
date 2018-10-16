package catalog

import (
	"../core"
	"../helpers"
	"encoding/json"
	fr "github.com/DATA-DOG/fastroute"
	"net/http"
)

var checkErr = helpers.CheckErr
var productList []Product

type Product struct {
	Id    int     `json:"id"`
	Sku   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

func init() {
	initRoutes()
	initProducts()
}

func initRoutes() map[string]fr.Router {

	productRoutes := fr.Chain(
		fr.New("/api/product/list", ApiProductListHandler),
		fr.New("/api/product/:id", ApiProductHandler),
	)

	getRoutes, ok := core.Routes["GET"]
	if !ok {
		core.Routes["GET"] = productRoutes
	} else {
		core.Routes["GET"] = fr.Chain(getRoutes, productRoutes)
	}

}

func initProducts() {
	productList = make([]Product, 0)
	productList = append(productList, Product{1, "s-1", "test product no 1.", 23.99})
	productList = append(productList, Product{2, "s-2", "test product no 1.", 24.99})
	productList = append(productList, Product{3, "s-3", "test product no 1.", 25.99})
	productList = append(productList, Product{4, "s-4", "test product no 1.", 26.99})
}

func ApiProductHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := helpers.GetIntParameter(r, "id")
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

// Product model methods

func getProductById(productId int) *Product {
	for _, p := range productList {
		if p.Id == productId {
			return &p
		}
	}
	return nil
}
