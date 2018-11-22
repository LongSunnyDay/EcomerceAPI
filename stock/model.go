package stock

import (
	"go-api-ws/config"
	"go-api-ws/helpers"
)

type stockData struct {
	Sku                            string `json:"sku"`
	ItemId                         int    `json:"item_id"`
	ProductId                      int    `json:"product_id"`
	StockId                        int    `json:"stock_id"`
	QTY                            int    `json:"qty"`
	IsInStock                      bool   `json:"is_in_stock"`
	IsQtyDecimal                   bool   `json:"is_qty_decimal"`
	ShowDefaultNotificationMessage bool   `json:"show_default_notification_message"`
	UseConfigMinQty                bool   `json:"use_config_min_qty"`
	MinQty                         int    `json:"min_qty"`
	UseConfigMinSaleQty            int    `json:"use_config_min_sale_qty"`
	MinSaleQty                     int    `json:"min_sale_qty"`
	UseConfigMaxSaleQty            bool   `json:"use_config_max_sale_qty"`
	MaxSaleQty                     int    `json:"max_sale_qty"`
	UseConfigBackorders            bool   `json:"use_config_backorders"`
	Backorders                     int    `json:"backorders"`
	UseConfigNotifyStockQty        bool   `json:"use_config_notify_stock_qty"`
	NotifyStockQty                 int    `json:"notify_stock_qty"`
	UseConfigQtyIncrements         bool   `json:"use_config_qty_increments"`
	QtyIncrements                  int    `json:"qty_increments"`
	UseConfigEnableQtyInc          bool   `json:"use_config_enable_qty_inc"`
	EnableQtyIncrements            bool   `json:"enable_qty_increments"`
	UseConfigManageStock           bool   `json:"use_config_manage_stock"`
	ManageStock                    bool   `json:"manage_stock"`
	LowStockDate                   string `json:"low_stock_date"`
	IsDecimalDivided               bool   `json:"is_decimal_divided"`
	StockStatusChangedAuto         int    `json:"stock_status_changed_auto"`
}

func getDataFromDbBySku(itemSku string) stockData {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	var itemData stockData
	err = db.QueryRow("SELECT * FROM stock WHERE sku = ?", itemSku).Scan(
		&itemData.Sku,
		&itemData.Backorders,
		&itemData.EnableQtyIncrements,
		&itemData.IsInStock,
		&itemData.IsQtyDecimal,
		&itemData.ItemId,
		&itemData.LowStockDate,
		&itemData.ManageStock,
		&itemData.MaxSaleQty,
		&itemData.MinQty,
		&itemData.MinSaleQty,
		&itemData.NotifyStockQty,
		&itemData.ProductId,
		&itemData.QTY,
		&itemData.QtyIncrements,
		&itemData.ShowDefaultNotificationMessage,
		&itemData.StockId,
		&itemData.StockStatusChangedAuto,
		&itemData.UseConfigBackorders,
		&itemData.UseConfigEnableQtyInc,
		&itemData.UseConfigManageStock,
		&itemData.UseConfigMaxSaleQty,
		&itemData.UseConfigMinQty,
		&itemData.UseConfigMinSaleQty,
		&itemData.UseConfigNotifyStockQty,
		&itemData.UseConfigQtyIncrements,
		&itemData.IsDecimalDivided)
	helpers.PanicErr(err)
	return itemData
}

func insertDataToStock(data stockData) {
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
		"stock_status_changed_auto)"+
		" VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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
