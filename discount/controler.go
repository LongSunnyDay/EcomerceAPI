package discount

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/xeipuuv/gojsonschema"
	c "go-api-ws/config"
	"go-api-ws/helpers"
	"log"
	"net/http"
	"time"
)

var (
	CouponDiscountPercent float64
	CouponDiscountAmount  float64
	CouponUsed            bool
)

func createDiscount(w http.ResponseWriter, r *http.Request) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file://discount/jsonSchemaModels/discountSchema.json")
	var discount Discount
	_ = json.NewDecoder(r.Body).Decode(&discount)
	documentLoader := gojsonschema.NewGoLoader(discount)
	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	helpers.PanicErr(err)

	if result.Valid() {
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
	} else {
		err = json.NewEncoder(w).Encode("There is and error registering discount:")
		helpers.PanicErr(err)
		fmt.Printf("The document is not valid. See errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

func getDiscount(w http.ResponseWriter, r *http.Request) {
	var discount Discount
	discountID := chi.URLParam(r, "discountID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	err = db.QueryRow("SELECT * FROM discount c WHERE id=?", discountID).
		Scan(&discount.Id, &discount.Sku, &discount.DiscountPercent, &discount.DiscountAmount)
	helpers.CheckErr(err)
	err = json.NewEncoder(w).Encode(discount)
	helpers.PanicErr(err)
}

func getDiscountList(w http.ResponseWriter, r *http.Request) {
	var discount Discount
	var discounts []Discount
	discounts = []Discount{}

	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	rows, err := db.Query("SELECT id, sku, discountPercent, discountAmount FROM discount")
	helpers.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
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
	err = json.NewEncoder(w).Encode(discounts)
	helpers.PanicErr(err)

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
	var discount Discount
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

func createCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon Coupon
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

func getCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon Coupon
	couponID := chi.URLParam(r, "couponID")
	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	err = db.QueryRow("SELECT * FROM coupon c WHERE id=?", couponID).
		Scan(&coupon.Id, &coupon.Code, &coupon.DiscountPercent, &coupon.DiscountAmount, &coupon.ExpirationDate, &coupon.UsageLimit, &coupon.TimesUsed, &coupon.CreatedAt)
	helpers.CheckErr(err)
	err = json.NewEncoder(w).Encode(coupon)
	helpers.PanicErr(err)
}

func getCouponList(w http.ResponseWriter, r *http.Request) {
	var coupon Coupon
	var coupons []Coupon
	coupons = []Coupon{}

	db, err := c.Conf.GetDb()
	helpers.CheckErr(err)

	rows, err := db.Query("SELECT id, code, discountPercent, discountAmount, expirationDate, usageLimit, timesUsed, createdAt FROM coupon")
	helpers.CheckErr(err)
	defer rows.Close()

	for rows.Next() {
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
	err = json.NewEncoder(w).Encode(coupons)
	helpers.PanicErr(err)
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
	var coupon Coupon
	err := json.NewDecoder(r.Body).Decode(&coupon)
	helpers.PanicErr(err)

	db, err := c.Conf.GetDb()
	helpers.PanicErr(err)
	query, err := db.Prepare("Update coupon set code=?, discountPercent=?, discountAmount=?, expirationDate=?, usageLimit=? where id=?")
	helpers.PanicErr(err)

	_, er := query.Exec(coupon.Code, coupon.DiscountPercent, coupon.DiscountAmount, coupon.ExpirationDate, coupon.UsageLimit, couponID)
	helpers.CheckIfRowExistsInMysql(db, "coupon", "id", couponID)
	helpers.PanicErr(er)
	if er == nil {
		fmt.Println(coupon.Code + " updated in mysql")
	}

}

func applyCoupon(w http.ResponseWriter, r *http.Request) {
	couponCode := r.URL.Query()["coupon"][0]
	db, err := c.Conf.GetDb()
	var coupon Coupon
	helpers.CheckErr(err)
	helpers.CheckIfRowExistsInMysql(db, "coupon", "code", couponCode)

	err = db.QueryRow("SELECT * FROM coupon c WHERE code=?", couponCode).
		Scan(&coupon.Id, &coupon.Code, &coupon.DiscountPercent, &coupon.DiscountAmount, &coupon.ExpirationDate, &coupon.UsageLimit, &coupon.TimesUsed, &coupon.CreatedAt)
	helpers.CheckErr(err)

	diff, err := time.Parse(time.RFC3339, coupon.ExpirationDate)

	helpers.CheckErr(err)
	var responseResult bool

	if diff.Sub(time.Now()) > 0 {
		fmt.Println("Coupon valid")
		if coupon.UsageLimit > coupon.TimesUsed {
			query, err := db.Prepare("Update coupon set timesUsed=? where code=?")
			helpers.PanicErr(err)

			_, er := query.Exec(coupon.TimesUsed+1, couponCode)
			helpers.PanicErr(er)
			CouponDiscountAmount = coupon.DiscountAmount
			CouponDiscountPercent = coupon.DiscountPercent
			CouponUsed = true

			responseResult = true
		} else {
			log.Fatal("Coupon code used up")
		}

	} else {
		fmt.Println(diff.Sub(time.Now()), diff)
		log.Fatal("Coupon expired")

	}
	response := helpers.Response{
		Code:   http.StatusOK,
		Result: responseResult}
	response.SendResponse(w)
}
