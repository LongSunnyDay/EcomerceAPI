package shipping

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

type methods []method

type method struct {
	Id           int    `json:"id,omitempty"`
	CarrierCode  string `json:"carrier_code"`
	MethodCode   string `json:"method_code"`
	CarrierTitle string `json:"carrier_title"`
	MethodTitle  string `json:"method_title"`
	Amount       float64    `json:"amount"`
	BaseAmount   float64    `json:"base_amount"`
	Available    bool   `json:"available"`
	ErrorMessage string `json:"error_message"`
	PriceExclTax int    `json:"price_excl_tax"`
	PriceInclTax float64    `json:"price_incl_tax"`
}

func (m method) insertToDb() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO shippingMethods("+
		"carrier_code, "+
		"method_code, "+
		"carrier_title, "+
		"method_title, "+
		"amount, "+
		"base_amount, "+
		"available, "+
		"error_message, "+
		"price_excl_tax, "+
		"price_incl_tax) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		m.CarrierCode, m.MethodCode, m.CarrierTitle, m.MethodTitle, m.Amount,
		m.BaseAmount, m.Available, m.ErrorMessage, m.PriceExclTax, m.PriceInclTax)
	helpers.PanicErr(err)
}

func getShippingMethodsFromDb() methods {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT " +
		"carrier_code, " +
		"method_code, " +
		"carrier_title, " +
		"method_title, " +
		"amount, " +
		"base_amount, " +
		"available, " +
		"error_message, " +
		"price_excl_tax, " +
		"price_incl_tax" +
		" FROM shippingMethods")
	var methods []method
	defer rows.Close()
	for rows.Next() {
		var method method
		if err := rows.Scan(&method.CarrierCode, &method.MethodCode, &method.CarrierTitle, &method.MethodTitle, &method.Amount,
			&method.BaseAmount, &method.Available, &method.ErrorMessage, &method.PriceExclTax, &method.PriceInclTax); err != nil {
			helpers.PanicErr(err)
		}
		methods = append(methods, method)
	}
	return methods
}

func (m method) updatePaymentMethodInDb() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("UPDATE shippingMethods s SET "+
		"s.carrier_code = ?, "+
		"s.method_code = ?, "+
		"s.carrier_title = ?, "+
		"s.method_title = ?, "+
		"s.amount = ?, "+
		"s.base_amount = ?, "+
		"s.available = ?, "+
		"s.error_message = ?, "+
		"s.price_excl_tax = ?, "+
		"s.price_incl_tax = ? "+
		"WHERE s.Id = ?",
		m.CarrierCode, m.MethodCode, m.CarrierTitle, m.MethodTitle, m.Amount,
		m.BaseAmount, m.Available, m.ErrorMessage, m.PriceExclTax, m.PriceInclTax, m.Id)
	helpers.PanicErr(err)
}

func removePaymentMethodFromDb(id string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("DELETE FROM shippingMethods  WHERE Id = ?", id)
	helpers.PanicErr(err)
}

func GetShippingMethod(shippingCarrier string, shippingMethod string) method {
	//fmt.Println(shippingCarrier, shippingMethod)
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var method method
	err = db.QueryRow("SELECT * FROM shippingMethods WHERE carrier_code = ? AND method_code = ?", shippingCarrier, shippingMethod).
		Scan(&method.Id, &method.CarrierCode, &method.MethodCode, &method.CarrierTitle, &method.MethodTitle,
		&method.Amount, &method.BaseAmount, &method.Available, &method.ErrorMessage, &method.PriceExclTax, &method.PriceInclTax)
	helpers.PanicErr(err)
	return method
}
