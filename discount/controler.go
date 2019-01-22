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
	"time"
	"log"
)

var CouponDiscountPercent float64
var CouponDiscountAmount float64
var CouponUsed bool

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

func createCoupon(w http.ResponseWriter, r *http.Request){
	var coupon m.Coupon
	err := json.NewDecoder(r.Body).Decode(&coupon)
	helpers.CheckErr(err)
	fmt.Println(coupon)

	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)
	result, err := db.Exec("INSERT INTO coupon("+
		"id, "+
		"code, "+
		"discountPercent, "+
		"discountAmount, "+
		"expirationDate, "+
		"usageLimit, "+
		"timesUsed) "+
		" VALUES(?, ?, ?, ?, ?, ?, ?)",
		coupon.Id,
		coupon.Code,
		coupon.DiscountPercent,
		coupon.DiscountAmount,
		coupon.ExpirationDate,
		coupon.UsageLimit,
		coupon.TimesUsed)
	fmt.Println(result)
	helpers.PanicErr(err)

}

func getCoupon(w http.ResponseWriter, r *http.Request){
	var coupon m.Coupon
	couponID := chi.URLParam(r, "couponID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	err = db.QueryRow("SELECT * FROM coupon c WHERE id=?", couponID).
		Scan(&coupon.Id, &coupon.Code, &coupon.DiscountPercent, &coupon.DiscountAmount, &coupon.ExpirationDate, &coupon.UsageLimit, &coupon.TimesUsed, &coupon.CreatedAt)
	helpers.CheckErr(err)
	json.NewEncoder(w).Encode(coupon)
}

func getCouponList(w http.ResponseWriter, r *http.Request) {
	var coupon m.Coupon
	var coupons []m.Coupon
	coupons = []m.Coupon{}

	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	rows, err := db.Query ("SELECT id, code, discountPercent, discountAmount, expirationDate, usageLimit, timesUsed, createdAt FROM coupon")
	helpers.CheckErr(err)
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(
			&coupon.Id,
			&coupon.Code,
			&coupon.DiscountPercent,
			&coupon.DiscountAmount,
			&coupon.ExpirationDate,
			&coupon.UsageLimit,
			&coupon.TimesUsed,
			&coupon.CreatedAt)
		helpers.CheckErr(err)
		coupons = append(coupons, coupon)
	}
	err = rows.Err()
	helpers.CheckErr(err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coupons)
}

func removeCoupon(w http.ResponseWriter, r *http.Request) {
	couponID := chi.URLParam(r, "couponID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	res, err := db.Exec("DELETE c FROM coupon c WHERE c.id=?", couponID)
	fmt.Println(res)
	helpers.CheckErr(err)
}

func updateCoupon(w http.ResponseWriter, r *http.Request) {
	couponID := chi.URLParam(r, "couponID")
	var coupon m.Coupon
	err := json.NewDecoder(r.Body).Decode(&coupon)
	helpers.PanicErr(err)

	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)
	query, err := db.Prepare("Update coupon set code=?, discountPercent=?, discountAmount=?, expirationDate=?, usageLimit=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(coupon.Code, coupon.DiscountPercent, coupon.DiscountAmount, coupon.ExpirationDate, coupon.UsageLimit, couponID)
	helpers.CheckIfRowExistsInMysql(db, "coupon", "id", couponID)
	helpers.PanicErr(er)
	if er == nil{
		fmt.Println(coupon.Code + " updated in mysql")
	}

}

func applyCoupon(w http.ResponseWriter, r *http.Request){
	couponCode := r.URL.Query()["coupon"][0]
	//cartId := r.URL.Query()["cartId"][0]
	db, err := c.Conf.GetDb()
	var coupon m.Coupon
	helpers.CheckErr(err)
	helpers.CheckIfRowExistsInMysql(db, "coupon", "code", couponCode)

	err = db.QueryRow("SELECT * FROM coupon c WHERE code=?", couponCode).
		Scan(&coupon.Id, &coupon.Code, &coupon.DiscountPercent, &coupon.DiscountAmount, &coupon.ExpirationDate, &coupon.UsageLimit, &coupon.TimesUsed, &coupon.CreatedAt)
	helpers.CheckErr(err)
	//json.NewEncoder(w).Encode(coupon)
	//fmt.Println(coupon.Code)

	diff,err := time.Parse(time.RFC3339, coupon.ExpirationDate)

	helpers.CheckErr(err)
	var responseResult bool

	if diff.Sub(time.Now()) > 0 {
		fmt.Println("Coupon valid")
		if coupon.UsageLimit > coupon.TimesUsed {
			query, err := db.Prepare("Update coupon set timesUsed=? where code=?")
			helpers.PanicErr(err)

			_, er := query.Exec(coupon.TimesUsed + 1, couponCode)
			helpers.PanicErr(er)
			CouponDiscountAmount = coupon.DiscountAmount
			CouponDiscountPercent = coupon.DiscountPercent
			CouponUsed = true

			responseResult = true
		}else{
			log.Fatal("Coupon code used up")
		}

	}else {
		fmt.Println(diff.Sub(time.Now()), diff)
		log.Fatal("Coupon expired")

	}
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: responseResult}
	//json.NewEncoder(w).Encode(response)
	//fmt.Println(response)
	response.SendResponse(w)
}