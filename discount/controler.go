package discount

import (
	"net/http"
	"github.com/xeipuuv/gojsonschema"
	m "go-api-ws/discount/models"
	"encoding/json"
	"go-api-ws/helpers"
	c "go-api-ws/config"
	"fmt"
	"github.com/go-chi/chi"
)

func createDiscount(w http.ResponseWriter, r *http.Request){
	var schemaLoader = gojsonschema.NewReferenceLoader("file://discount/models/discountSchema.json")
	var discount m.Discount
	_ = json.NewDecoder(r.Body).Decode(&discount)
	documentLoader := gojsonschema.NewGoLoader(discount)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid(){
		db, err := c.Conf.GetDb()
		helpers.PanicErr(err)
		result, err := db.Exec("INSERT INTO discount("+
			"id, "+
			"sku, "+
			"discountPercent, "+
			"discountAmount) "+
			" VALUES(?, ?, ?, ?)",
			discount.Id,
			discount.Sku,
			discount.DiscountPercent,
			discount.DiscountAmount)
		fmt.Println(result)
		helpers.PanicErr(err)
	} else{
		json.NewEncoder(w).Encode("There is and error registering discount:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func getDiscount(w http.ResponseWriter, r *http.Request) {
	var discount m.Discount
	discountID := chi.URLParam(r, "discountID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	err = db.QueryRow("SELECT * FROM discount c WHERE id=?", discountID).
		Scan(&discount.Id, &discount.Sku, &discount.DiscountPercent, &discount.DiscountAmount)
	helpers.CheckErr(err)
	json.NewEncoder(w).Encode(discount)
}

func getDiscountList (w http.ResponseWriter, r *http.Request) {
	var discount m.Discount
	var discounts []m.Discount
	discounts = []m.Discount{}

	db, err :=  c.Conf.GetDb()
	helpers.CheckErr(err)

	rows, err := db.Query ("SELECT id, sku, discountPercent, discountAmount FROM discount")
	helpers.CheckErr(err)
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(
			&discount.Id,
			&discount.Sku,
			&discount.DiscountPercent,
			&discount.DiscountAmount)
		helpers.CheckErr(err)
		discounts = append(discounts, discount)
	}
	err = rows.Err()
	helpers.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(discounts)

}

func removeDiscount(w http.ResponseWriter, r *http.Request) {
	discountID := chi.URLParam(r, "discountID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	res, err := db.Exec("DELETE d FROM discount d WHERE d.id=?", discountID)
	fmt.Println(res)
	helpers.CheckErr(err)
}

func updateDiscount(w http.ResponseWriter, r *http.Request) {
	discountID := chi.URLParam(r, "discountID")
	var discount m.Discount
	err := json.NewDecoder(r.Body).Decode(&discount)
	helpers.PanicErr(err)

	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	query, err := db.Prepare("Update discount set sku=?, discountPercent=?, discountAmount=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(discount.Sku, discount.DiscountPercent, discount.DiscountAmount, discountID)
	helpers.PanicErr(er)
	fmt.Println(discount.Sku + " updated in mysql")

}