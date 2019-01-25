package stock

import (
	"github.com/bugsnag/bugsnag-go/errors"
	"go-api-ws/config"
	"go-api-ws/helpers"
)

func (data *DataStock) GetDataFromDbBySku(itemSku string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM stock WHERE sku = ?", itemSku).Scan(
		&data.Sku,
		&data.Backorders,
		&data.EnableQtyIncrements,
		&data.IsInStock,
		&data.IsQtyDecimal,
		&data.ItemId,
		&data.LowStockDate,
		&data.ManageStock,
		&data.MaxSaleQty,
		&data.MinQty,
		&data.MinSaleQty,
		&data.NotifyStockQty,
		&data.ProductId,
		&data.QTY,
		&data.QtyIncrements,
		&data.ShowDefaultNotificationMessage,
		&data.StockId,
		&data.StockStatusChangedAuto,
		&data.UseConfigBackorders,
		&data.UseConfigEnableQtyInc,
		&data.UseConfigManageStock,
		&data.UseConfigMaxSaleQty,
		&data.UseConfigMinQty,
		&data.UseConfigMinSaleQty,
		&data.UseConfigNotifyStockQty,
		&data.UseConfigQtyIncrements,
		&data.IsDecimalDivided)
	helpers.PanicErr(err)
}

func (data DataStock) insertDataToStock() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO stock("+
		"sku, "+
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
		"use_config_notify_stock_qty, "+
		"notify_stock_qty, "+
		"use_config_qty_increments, "+
		"qty_increments, "+
		"use_config_enable_qty_inc, "+
		"enable_qty_increments, "+
		"use_config_manage_stock, "+
		"manage_stock, "+
		"low_stock_date, "+
		"is_decimal_divided, "+
		"stock_status_changed_auto) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.Sku,
		data.ItemId,
		data.ProductId,
		data.StockId,
		data.QTY,
		data.IsInStock,
		data.IsQtyDecimal,
		data.ShowDefaultNotificationMessage,
		data.UseConfigMinQty,
		data.MinQty,
		data.UseConfigMinSaleQty,
		data.MinSaleQty,
		data.UseConfigMaxSaleQty,
		data.MaxSaleQty,
		data.UseConfigBackorders,
		data.Backorders,
		data.UseConfigNotifyStockQty,
		data.NotifyStockQty,
		data.UseConfigQtyIncrements,
		data.QtyIncrements,
		data.UseConfigEnableQtyInc,
		data.EnableQtyIncrements,
		data.UseConfigManageStock,
		data.ManageStock,
		data.LowStockDate,
		data.IsDecimalDivided,
		data.StockStatusChangedAuto)
	helpers.PanicErr(err)
}

func (data DataStock) updateDataInDb() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("UPDATE stock s SET "+
		"s.item_id = ?, "+
		"s.product_id = ?, "+
		"s.stock_id = ?, "+
		"s.qty = ?, "+
		"s.is_in_stock = ?, "+
		"s.is_qty_decimal = ?, "+
		"s.show_default_notification_message = ?, "+
		"s.use_config_min_qty = ?, "+
		"s.min_qty = ?, "+
		"s.use_config_min_sale_qty = ?, "+
		"s.min_sale_qty = ?, "+
		"s.use_config_max_sale_qty = ?, "+
		"s.max_sale_qty = ?, "+
		"s.use_config_backorders = ?, "+
		"s.backorders = ?, "+
		"s.use_config_notify_stock_qty = ?, "+
		"s.notify_stock_qty = ?, "+
		"s.use_config_qty_increments = ?, "+
		"s.qty_increments = ?, "+
		"s.use_config_enable_qty_inc = ?, "+
		"s.enable_qty_increments = ?, "+
		"s.use_config_manage_stock = ?, "+
		"s.manage_stock = ?, "+
		"s.low_stock_date = ?, "+
		"s.is_decimal_divided = ?, "+
		"s.stock_status_changed_auto = ? "+
		"WHERE s.sku = ?",
		data.ItemId,
		data.ProductId,
		data.StockId,
		data.QTY,
		data.IsInStock,
		data.IsQtyDecimal,
		data.ShowDefaultNotificationMessage,
		data.UseConfigMinQty,
		data.MinQty,
		data.UseConfigMinSaleQty,
		data.MinSaleQty,
		data.UseConfigMaxSaleQty,
		data.MaxSaleQty,
		data.UseConfigBackorders,
		data.Backorders,
		data.UseConfigNotifyStockQty,
		data.NotifyStockQty,
		data.UseConfigQtyIncrements,
		data.QtyIncrements,
		data.UseConfigEnableQtyInc,
		data.EnableQtyIncrements,
		data.UseConfigManageStock,
		data.ManageStock,
		data.LowStockDate,
		data.IsDecimalDivided,
		data.StockStatusChangedAuto,
		data.Sku)
	helpers.PanicErr(err)
}

func removeItemFromDb(itemSku string) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("DELETE FROM stock WHERE sku = ?", itemSku)
	helpers.PanicErr(err)
}

func (data *DataStock) CheckSOOT(itemSku string, itemQty float64) (err error) {
	if data.Sku == itemSku && data.QTY >= itemQty {
		return nil
	} else {
		err := errors.New("item out of stock", 0)
		return err
	}
}
