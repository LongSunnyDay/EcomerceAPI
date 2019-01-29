package currency

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/xeipuuv/gojsonschema"
	c "go-api-ws/config"
	m "go-api-ws/currency/models"
	"go-api-ws/helpers"
	"net/http"
)

func createCurrency(w http.ResponseWriter, r *http.Request) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://currency/models/currencySchema.json")
	var currency m.Currency
	var currencies []m.Currency
	err := json.NewDecoder(r.Body).Decode(&currency)
	helpers.CheckErr(err)
	documentLoader := gojsonschema.NewGoLoader(currency)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {
		db, err := c.Conf.GetDb()
		helpers.PanicErr(err)

		result, err := db.Exec("INSERT INTO currency("+
			"id,"+
			"name, "+
			"code, "+
			"sign, "+
			"defaultCurrency) "+
			" VALUES(?, ?, ?, ?, ?)",
			currency.Id,
			currency.Name,
			currency.Code,
			currency.Sign,
			currency.DefaultCurrency)
		fmt.Println(result)
		helpers.PanicErr(err)

		currencies = append(currencies, currency)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Currency: " + currency.Name + " has been registered. ")

	} else {
		json.NewEncoder(w).Encode("There is and error registering currency:")
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func getCurrency(w http.ResponseWriter, r *http.Request) {
	var currency m.Currency
	currencyID := chi.URLParam(r, "currencyID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	err = db.QueryRow("SELECT * FROM currency c WHERE id=?", currencyID).
		Scan(&currency.Id, &currency.Name, &currency.Code, &currency.Sign, &currency.DefaultCurrency)
	helpers.CheckErr(err)
	json.NewEncoder(w).Encode(currency)
}

func getCurrencyList(w http.ResponseWriter, r *http.Request) {
	var currency m.Currency
	var currencies []m.Currency
	currencies = []m.Currency{}

	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	rows, err := db.Query("SELECT id, name, code, sign, defaultCurrency FROM currency")
	helpers.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&currency.Id,
			&currency.Name,
			&currency.Code,
			&currency.Sign,
			&currency.DefaultCurrency)
		helpers.CheckErr(err)
		currencies = append(currencies, currency)
	}
	err = rows.Err()
	helpers.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currencies)

}

func removeCurrency(w http.ResponseWriter, r *http.Request) {
	currencyID := chi.URLParam(r, "currencyID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	res, err := db.Exec("DELETE c FROM currency c WHERE c.id=?", currencyID)
	fmt.Println(res)
	helpers.CheckErr(err)
}

func updateCurrency(w http.ResponseWriter, r *http.Request) {
	currencyID := chi.URLParam(r, "currencyID")
	var currency m.Currency
	err := json.NewDecoder(r.Body).Decode(&currency)
	helpers.PanicErr(err)

	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)

	query, err := db.Prepare("Update currency set name=?, code=?, sign=?, defaultCurrency=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(currency.Name, currency.Code, currency.Sign, currency.DefaultCurrency, currencyID)
	helpers.PanicErr(er)
	fmt.Println(currency.Name + " updated in mysql")
}
