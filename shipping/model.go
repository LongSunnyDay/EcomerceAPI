package shipping

import (
	"fmt"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type (
	methods []Method

	Method struct {
		Id           int     `json:"id,omitempty"`
		CarrierCode  string  `json:"carrier_code"`
		MethodCode   string  `json:"method_code"`
		CarrierTitle string  `json:"carrier_title"`
		MethodTitle  string  `json:"method_title"`
		Amount       float64 `json:"amount"`
		BaseAmount   float64 `json:"base_amount"`
		Available    bool    `json:"available"`
		ErrorMessage string  `json:"error_message"`
		PriceExclTax int     `json:"price_excl_tax"`
		PriceInclTax float64 `json:"price_incl_tax"`
	}

	DiscountConfig struct {
		ShippingDiscountIsAvailable bool `yaml:"shippingDiscountIsAvailable"`
		ShippingDiscountFrom        float64  `yaml:"shippingDiscountFrom"`
		ShippingDiscountTo          float64  `yaml:"shippingDiscountTo"`
	}
)

var DisConf *DiscountConfig

func (m Method) insertToDb() {
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
		" FROM shippingMethods WHERE available = true")
	var methods []Method
	defer rows.Close()
	for rows.Next() {
		var method Method
		if err := rows.Scan(&method.CarrierCode, &method.MethodCode, &method.CarrierTitle, &method.MethodTitle, &method.Amount,
			&method.BaseAmount, &method.Available, &method.ErrorMessage, &method.PriceExclTax, &method.PriceInclTax); err != nil {
			helpers.PanicErr(err)
		}
		methods = append(methods, method)
	}
	return methods
}

func (m Method) updatePaymentMethodInDb() {
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

func GetShippingMethod(shippingCarrier string, shippingMethod string) Method {
	fmt.Println(shippingCarrier, "<<<<>>>>", shippingMethod)
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var method Method
	err = db.QueryRow("SELECT * FROM shippingMethods WHERE carrier_code = ? AND method_code = ?", shippingCarrier, shippingMethod).
		Scan(&method.Id, &method.CarrierCode, &method.MethodCode, &method.CarrierTitle, &method.MethodTitle,
			&method.Amount, &method.BaseAmount, &method.Available, &method.ErrorMessage, &method.PriceExclTax, &method.PriceInclTax)
	helpers.PanicErr(err)
	return method
}

func GetConfig(configFile string) *DiscountConfig {
	DisConf = &DiscountConfig{}
	if configFile != "" {
		err := DisConf.GetConfFromFile(configFile)
		helpers.PanicErr(err)
	}
	return DisConf
}

func (c *DiscountConfig) GetConfFromFile(fileName string) error {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/shipping/" + fileName)
	if err != nil {
		log.Printf("%s file read error.  #%v\n", fileName, err)
	}
	return c.GetConfFromString(string(yamlFile))
}

func (c *DiscountConfig) GetConfFromString(yamlString string) error {

	err := yaml.Unmarshal([]byte(yamlString), c)
	if err != nil {
		log.Fatalf("%s parse error %v\n", yamlString, err)
	}
	return err
}
