package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-ws/attribute"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"io/ioutil"
	"math"
	"net/http"
)

func BuildSKUFromItemAttributes(itemAttributes []attribute.ItemAttribute, sku string) string {

	for _, itemAttribute := range itemAttributes {
		if itemAttribute.Name == "size" {
			sku = sku + "-" + itemAttribute.Label
		}
	}

	for _, itemAttribute := range itemAttributes {
		if itemAttribute.Name == "color" {
			sku = sku + "-" + itemAttribute.Label
		}
	}
	return sku
}

func GetProductFromSolrBySKU(sku string) ConfigurableProductStruct {
	request := map[string]interface{}{
		"query": "+_type:product +sku:'" + sku + "'",
		"limit": 1}
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(request)
	helpers.PanicErr(err)
	resp, err := http.Post(
		attribute.SolrQueryUrl,
		attribute.ContentType,
		requestBytes)
	helpers.PanicErr(err)
	b, _ := ioutil.ReadAll(resp.Body)
	var solrResp solrResponse
	err = json.Unmarshal(b, &solrResp)
	helpers.PanicErr(err)
	if solrResp.Response.NumFound > 0 {
		return solrResp.Response.Docs[0]
	}
	return ConfigurableProductStruct{}
}

func GetProductPriceFromDbBySku(sku string, userPrice float64) bool {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var price float64
	err = db.QueryRow("SELECT final_price FROM product WHERE sku = ?", sku).Scan(&price)
	helpers.PanicErr(err)
	if price == math.Round(userPrice*100)/100 {
		return true
	} else {
		return false
	}
}

func (product SimpleProductStruct) insertSimpleProductToDb() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO product("+
		"doc_type, "+
		"sku, "+
		"name, "+
		"attribute_set_id, "+
		"price, "+
		"status, "+
		"visibility, "+
		"type_id, "+
		"created_at, "+
		"updated_at, "+
		"final_price, "+
		"max_price, "+
		"max_regular_price, "+
		"minimal_regular_price, "+
		"minimal_price, "+
		"regular_price, "+
		"item_id, "+
		"product_id, "+
		"stock_id, "+
		"qty, "+
		"is_in_stock, "+
		"is_qty_decimal, "+
		"show_default_notification_message, "+
		"use_config_min_qty, "+
		"min_qty, "+
		"use_config_min_sale_qty, "+
		"min_sale_qty, "+
		"use_config_max_sale_qty, "+
		"max_sale_qty, "+
		"use_config_backorders, "+
		"backorders, "+
		"use_config_qty_increments, "+
		"qty_increments, "+
		"use_config_notify_stock_qty, "+
		"notify_stock_qty, "+
		"use_config_enable_qty_inc, "+
		"enable_qty_increments, "+
		"use_config_manage_stock, "+
		"manage_stock, "+
		"is_decimal_divided, "+
		"stock_status_changed_auto, "+
		"tsk, "+
		"description, "+
		"image, "+
		"small_image, "+
		"thumbnail, "+
		"options_container, "+
		"url_key, "+
		"tax_class_id, "+
		"gender, "+
		"category_gear, "+
		"erin_recommends, "+
		"sale, "+
		"new) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		product.DocType,
		product.Sku,
		product.Name,
		product.AttributeSetID,
		product.Price,
		product.Status,
		product.Visibility,
		product.TypeID,
		product.CreatedAt,
		product.UpdatedAt,
		product.FinalPrice,
		product.MaxPrice,
		product.MaxRegularPrice,
		product.MinimalRegularPrice,
		product.MinimalPrice,
		product.RegularPrice,
		product.ItemID,
		product.ProductID,
		product.StockID,
		product.Qty,
		product.IsInStock,
		product.IsQtyDecimal,
		product.ShowDefaultNotificationMessage,
		product.UseConfigMinQty,
		product.MinQty,
		product.UseConfigMinSaleQty,
		product.MinSaleQty,
		product.UseConfigMaxSaleQty,
		product.MaxSaleQty,
		product.UseConfigBackorders,
		product.Backorders,
		product.UseConfigQtyIncrements,
		product.QtyIncrements,
		product.UseConfigNotifyStockQty,
		product.NotifyStockQty,
		product.UseConfigEnableQtyInc,
		product.EnableQtyIncrements,
		product.UseConfigManageStock,
		product.ManageStock,
		product.IsDecimalDivided,
		product.StockStatusChangedAuto,
		product.Tsk,
		product.Description,
		product.Image,
		product.SmallImage,
		product.Thumbnail,
		product.OptionsContainer,
		product.URLKey,
		product.TaxClassID,
		product.Gender,
		product.CategoryGear,
		product.ErinRecommends,
		product.Sale,
		product.New)
	helpers.PanicErr(err)
}

func getSimpleProductFromDbBySku(sku string) SimpleProductStruct {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var product SimpleProductStruct
	err = db.QueryRow("SELECT * FROM product WHERE sku = ?", sku).Scan(
		&product.ID,
		&product.DocType,
		&product.Sku,
		&product.Name,
		&product.AttributeSetID,
		&product.Price,
		&product.Status,
		&product.Visibility,
		&product.TypeID,
		&product.CreatedAt,
		&product.UpdatedAt,
		&product.FinalPrice,
		&product.MaxPrice,
		&product.MaxRegularPrice,
		&product.MinimalRegularPrice,
		&product.MinimalPrice,
		&product.RegularPrice,
		&product.ItemID,
		&product.ProductID,
		&product.StockID,
		&product.Qty,
		&product.IsInStock,
		&product.IsQtyDecimal,
		&product.ShowDefaultNotificationMessage,
		&product.UseConfigMinQty,
		&product.MinQty,
		&product.UseConfigMinSaleQty,
		&product.MinSaleQty,
		&product.UseConfigMaxSaleQty,
		&product.MaxSaleQty,
		&product.UseConfigBackorders,
		&product.Backorders,
		&product.UseConfigQtyIncrements,
		&product.QtyIncrements,
		&product.UseConfigNotifyStockQty,
		&product.NotifyStockQty,
		&product.UseConfigEnableQtyInc,
		&product.EnableQtyIncrements,
		&product.UseConfigManageStock,
		&product.ManageStock,
		&product.IsDecimalDivided,
		&product.StockStatusChangedAuto,
		&product.Tsk,
		&product.Description,
		&product.Image,
		&product.SmallImage,
		&product.Thumbnail,
		&product.OptionsContainer,
		&product.URLKey,
		&product.TaxClassID,
		&product.Gender,
		&product.CategoryGear,
		&product.ErinRecommends,
		&product.Sale,
		&product.New)
	helpers.PanicErr(err)
	return product
}

func removeProductFromDbBySku(sku string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("DELETE FROM product WHERE sku = ?", sku)
	helpers.PanicErr(err)
}

func (product SimpleProductStruct) updateProductInDb() int64 {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	res, err := db.Exec("UPDATE product p SET "+
		"p.doc_type = ?, "+
		"p.sku = ?, "+
		"p.name = ?, "+
		"p.attribute_set_id = ?, "+
		"p.price = ?, "+
		"p.status = ?, "+
		"p.visibility = ?, "+
		"p.type_id = ?, "+
		"p.created_at = ?, "+
		"p.updated_at = ?, "+
		"p.final_price = ?, "+
		"p.max_price = ?, "+
		"p.max_regular_price = ?, "+
		"p.minimal_regular_price = ?, "+
		"p.minimal_price = ?, "+
		"p.regular_price = ?, "+
		"p.item_id = ?, "+
		"p.product_id = ?, "+
		"p.stock_id = ?, "+
		"p.qty = ?, "+
		"p.is_in_stock = ?, "+
		"p.is_qty_decimal = ?, "+
		"p.show_default_notification_message = ?, "+
		"p.use_config_min_qty = ?, "+
		"p.min_qty = ?, "+
		"p.use_config_min_sale_qty = ?, "+
		"p.min_sale_qty = ?, "+
		"p.use_config_max_sale_qty = ?, "+
		"p.max_sale_qty = ?, "+
		"p.use_config_backorders = ?, "+
		"p.backorders = ?, "+
		"p.use_config_qty_increments = ?, "+
		"p.qty_increments = ?, "+
		"p.use_config_notify_stock_qty = ?, "+
		"p.notify_stock_qty = ?, "+
		"p.use_config_enable_qty_inc = ?, "+
		"p.enable_qty_increments = ?, "+
		"p.use_config_manage_stock = ?, "+
		"p.manage_stock = ?, "+
		"p.is_decimal_divided = ?, "+
		"p.stock_status_changed_auto = ?, "+
		"p.tsk = ?, "+
		"p.description = ?, "+
		"p.image = ?, "+
		"p.small_image = ?, "+
		"p.thumbnail = ?, "+
		"p.options_container = ?, "+
		"p.url_key = ?, "+
		"p.tax_class_id = ?, "+
		"p.gender = ?, "+
		"p.category_gear = ?, "+
		"p.erin_recommends = ?, "+
		"p.sale = ?, "+
		"p.new = ? "+
		"WHERE p.id = ?",
		product.DocType,
		product.Sku,
		product.Name,
		product.AttributeSetID,
		product.Price,
		product.Status,
		product.Visibility,
		product.TypeID,
		product.CreatedAt,
		product.UpdatedAt,
		product.FinalPrice,
		product.MaxPrice,
		product.MaxRegularPrice,
		product.MinimalRegularPrice,
		product.MinimalPrice,
		product.RegularPrice,
		product.ItemID,
		product.ProductID,
		product.StockID,
		product.Qty,
		product.IsInStock,
		product.IsQtyDecimal,
		product.ShowDefaultNotificationMessage,
		product.UseConfigMinQty,
		product.MinQty,
		product.UseConfigMinSaleQty,
		product.MinSaleQty,
		product.UseConfigMaxSaleQty,
		product.MaxSaleQty,
		product.UseConfigBackorders,
		product.Backorders,
		product.UseConfigQtyIncrements,
		product.QtyIncrements,
		product.UseConfigNotifyStockQty,
		product.NotifyStockQty,
		product.UseConfigEnableQtyInc,
		product.EnableQtyIncrements,
		product.UseConfigManageStock,
		product.ManageStock,
		product.IsDecimalDivided,
		product.StockStatusChangedAuto,
		product.Tsk,
		product.Description,
		product.Image,
		product.SmallImage,
		product.Thumbnail,
		product.OptionsContainer,
		product.URLKey,
		product.TaxClassID,
		product.Gender,
		product.CategoryGear,
		product.ErinRecommends,
		product.Sale,
		product.New,
		product.ID)
	helpers.PanicErr(err)
	rowNumber, err := res.RowsAffected()
	helpers.PanicErr(err)
	fmt.Printf("Product id %v got %v modyfied", product.ID, rowNumber)
	return rowNumber
}
