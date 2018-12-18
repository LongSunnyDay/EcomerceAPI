package order

import (
	"encoding/json"
	"fmt"
	"go-api-ws/cart"
	"go-api-ws/helpers"
	"go-api-ws/product"
	"go-api-ws/stock"
	"net/http"
)

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderData PlaceOrderData

	err := json.NewDecoder(r.Body).Decode(&orderData)
	helpers.PanicErr(err)

	//fmt.Printf("%+v\n", orderData)

	cartItemsFromMongo := cart.GetUserCartFromMongoByID(orderData.UserId)

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
			fmt.Println(item.Sku)
		}
	}

	for _, item := range orderData.Products {
		checkPrice := product.GetProductPriceFromDbBySku(item.Sku, item.FinalPrice)
		if !checkPrice {
			fmt.Printf("Product %v price doesn't match with price in db.", item.Name)
		}
	}

}
