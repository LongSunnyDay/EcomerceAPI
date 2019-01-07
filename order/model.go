package order

import (
	"fmt"
	"go-api-ws/config"
	"go-api-ws/helpers"
	"time"
)

type History struct {
	ID                                    int64              `json:"id,omitempty"`
	AppliedRuleIds                        string             `json:"applied_rule_ids"`
	BaseCurrencyCode                      string             `json:"base_currency_code"`
	BaseDiscountAmount                    float64            `json:"base_discount_amount"`
	BaseGrandTotal                        float64            `json:"base_grand_total"`
	BaseDiscountTaxCompensationAmount     float64            `json:"base_discount_tax_compensation_amount"`
	BaseShippingAmount                    float64            `json:"base_shipping_amount"`
	BaseShippingDiscountAmount            float64            `json:"base_shipping_discount_amount"`
	BaseShippingInclTax                   float64            `json:"base_shipping_incl_tax"`
	BaseShippingTaxAmount                 float64            `json:"base_shipping_tax_amount"`
	BaseSubtotal                          float64            `json:"base_subtotal"`
	BaseSubtotalInclTax                   float64            `json:"base_subtotal_incl_tax"`
	BaseTaxAmount                         float64            `json:"base_tax_amount"`
	BaseTotalDue                          float64            `json:"base_total_due"`
	BaseToGlobalRate                      float64            `json:"base_to_global_rate"`
	BaseToOrderRate                       float64            `json:"base_to_order_rate"`
	BillingAddressId                      int64              `json:"billing_address_id"`
	CreatedAt                             time.Time          `json:"created_at"`
	CustomerEmail                         string             `json:"customer_email"`
	CustomerFirstname                     string             `json:"customer_firstname"`
	CustomerGroupId                       int64              `json:"customer_group_id"`
	CustomerId                            int64              `json:"customer_id"`
	CustomerIsGuest                       int                `json:"customer_is_guest"`
	CustomerLastname                      string             `json:"customer_lastname"`
	CustomerNoteNotify                    int                `json:"customer_note_notify"`
	DiscountAmount                        float64            `json:"discount_amount"`
	EmailSent                             int                `json:"email_sent"`
	EntityId                              int64              `json:"entity_id"`
	GlobalCurrencyCode                    string             `json:"global_currency_code"`
	GrandTotal                            float64            `json:"grand_total"`
	DiscountTaxCompensationAmount         float64            `json:"discount_tax_compensation_amount"`
	IncrementId                           string             `json:"increment_id"`
	IsVirtual                             int                `json:"is_virtual"`
	OrderCurrencyCode                     string             `json:"order_currency_code"`
	ProtectCode                           string             `json:"protect_code"`
	QuoteId                               int64              `json:"quote_id"`
	ShippingAmount                        float64            `json:"shipping_amount"`
	ShippingDescription                   string             `json:"shipping_description"`
	ShippingDiscountAmount                float64            `json:"shipping_discount_amount"`
	ShippingDiscountTaxCompensationAmount float64            `json:"shipping_discount_tax_compensation_amount"`
	ShippingInclTax                       float64            `json:"shipping_incl_tax"`
	ShippingTaxAmount                     float64            `json:"shipping_tax_amount"`
	State                                 string             `json:"state"`
	Status                                string             `json:"status"`
	StoreCurrencyCode                     string             `json:"store_currency_code"`
	StoreId                               int                `json:"store_id"`
	StoreName                             string             `json:"store_name"`
	StoreToBaseRate                       float64            `json:"store_to_base_rate"`
	StoreToOrderRate                      float64            `json:"store_to_order_rate"`
	Subtotal                              float64            `json:"subtotal"`
	SubtotalInclTax                       float64            `json:"subtotal_incl_tax"`
	TaxAmount                             float64            `json:"tax_amount"`
	TotalDue                              float64            `json:"total_due"`
	TotalItemCount                        float64            `json:"total_item_count"`
	TotalQtyOrdered                       float64            `json:"total_qty_ordered"`
	UpdatedAt                             time.Time          `json:"updated_at"`
	Weight                                float64            `json:"weight"`
	Items                                 []Item             `json:"items"`
	BillingAddress                        BillingAddress     `json:"billing_address"`
	Payment                               Payment            `json:"payment"`
	StatusHistories                       []string           `json:"status_histories"`
	ExtensionAttributes                   ExtensionAttribute `json:"extension_attributes"`
}

type ExtensionAttribute struct {
	ShippingAssignments []ShippingAssignment `json:"shipping_assignments"`
}

type Item struct {
	AmountRefunded                    float64   `json:"amount_refunded"`
	AppliedRuleIds                    string    `json:"applied_rule_ids"`
	BaseAmountRefunded                float64   `json:"base_amount_refunded"`
	BaseDiscountAmount                float64   `json:"base_discount_amount"`
	BaseDiscountInvoiced              float64   `json:"base_discount_invoiced"`
	BaseDiscountTaxCompensationAmount float64   `json:"base_discount_tax_compensation_amount"`
	BaseOriginalPrice                 float64   `json:"base_original_price"`
	BasePrice                         float64   `json:"base_price"`
	BasePriceInclTax                  float64   `json:"base_price_incl_tax"`
	BaseRowInvoiced                   float64   `json:"base_row_invoiced"`
	BaseRowTotal                      float64   `json:"base_row_total"`
	BaseRowTotalInclTax               float64   `json:"base_row_total_incl_tax"`
	BaseTaxAmount                     float64   `json:"base_tax_amount"`
	BaseTaxInvoiced                   float64   `json:"base_tax_invoiced"`
	CreatedAt                         time.Time `json:"created_at"`
	DiscountAmount                    float64   `json:"discount_amount"`
	DiscountInvoiced                  float64   `json:"discount_invoiced"`
	DiscountPercent                   float64   `json:"discount_percent"`
	FreeShipping                      int       `json:"free_shipping"`
	DiscountTaxCompensationAmount     float64   `json:"discount_tax_compensation_amount"`
	IsQtyDecimal                      int       `json:"is_qty_decimal"`
	IsVirtual                         int       `json:"is_virtual"`
	ItemId                            int       `json:"item_id"`
	Name                              string    `json:"name"`
	NoDiscount                        int       `json:"no_discount"`
	OrderId                           int64     `json:"order_id"`
	OriginalPrice                     float64   `json:"original_price"`
	ParentItemId                      int       `json:"parent_item_id,omitempty"`
	Price                             float64   `json:"price"`
	ProductId                         int       `json:"product_id"`
	ProductType                       string    `json:"product_type"`
	QtyCanceled                       float64   `json:"qty_canceled"`
	QtyInvoiced                       float64   `json:"qty_invoiced"`
	QtyOrdered                        float64   `json:"qty_ordered"`
	QtyRefunded                       float64   `json:"qty_refunded"`
	QtyShipped                        float64   `json:"qty_shipped"`
	QuoteItemId                       int64     `json:"quote_item_id"`
	RowInvoiced                       float64   `json:"row_invoiced"`
	RowTotal                          float64   `json:"row_total"`
	RowTotalInclTax                   float64   `json:"row_total_incl_tax,omitempty"`
	RowWeight                         float64   `json:"row_weight"`
	Sku                               string    `json:"sku"`
	StoreId                           int       `json:"store_id"`
	TaxAmount                         float64   `json:"tax_amount"`
	TaxInvoiced                       float64   `json:"tax_invoiced"`
	TaxPercent                        float64   `json:"tax_percent"`
	UpdatedAt                         time.Time `json:"updated_at"`
	Weight                            float64   `json:"weight"`
	//ParentItem                        ParentItem `json:"parent_item,omitempty"`
}

type ParentItem struct {
	AmountRefunded                    float64   `json:"amount_refunded"`
	AppliedRuleIds                    string    `json:"applied_rule_ids,omitempty"`
	BaseAmountRefunded                float64   `json:"base_amount_refunded"`
	BaseDiscountAmount                float64   `json:"base_discount_amount"`
	BaseDiscountInvoiced              float64   `json:"base_discount_invoiced"`
	BaseDiscountTaxCompensationAmount float64   `json:"base_discount_tax_compensation_amount,omitempty"`
	BaseOriginalPrice                 float64   `json:"base_original_price,omitempty"`
	BasePrice                         float64   `json:"base_price"`
	BasePriceInclTax                  float64   `json:"base_price_incl_tax,omitempty"`
	BaseRowInvoiced                   float64   `json:"base_row_invoiced"`
	BaseRowTotal                      float64   `json:"base_row_total"`
	BaseRowTotalInclTax               float64   `json:"base_row_total_incl_tax,omitempty"`
	BaseTaxAmount                     float64   `json:"base_tax_amount"`
	BaseTaxInvoiced                   float64   `json:"base_tax_invoiced"`
	CreatedAt                         time.Time `json:"created_at"`
	DiscountAmount                    float64   `json:"discount_amount"`
	DiscountInvoiced                  float64   `json:"discount_invoiced"`
	DiscountPercent                   float64   `json:"discount_percent"`
	FreeShipping                      int       `json:"free_shipping"`
	DiscountTaxCompensationAmount     float64   `json:"discount_tax_compensation_amount,omitempty"`
	IsQtyDecimal                      int       `json:"is_qty_decimal"`
	IsVirtual                         int       `json:"is_virtual"`
	ItemId                            int       `json:"item_id"`
	Name                              string    `json:"name"`
	NoDiscount                        int       `json:"no_discount"`
	OrderId                           int       `json:"order_id"`
	OriginalPrice                     float64   `json:"original_price"`
	ParentItemId                      int       `json:"parent_item_id,omitempty"`
	Price                             float64   `json:"price"`
	PriceInclTax                      float64   `json:"price_incl_tax,omitempty"`
	ProductId                         int       `json:"product_id"`
	ProductType                       string    `json:"product_type"`
	QtyCanceled                       float64   `json:"qty_canceled"`
	QtyInvoiced                       float64   `json:"qty_invoiced"`
	QtyOrdered                        float64   `json:"qty_ordered"`
	QtyRefunded                       float64   `json:"qty_refunded"`
	QtyShipped                        float64   `json:"qty_shipped"`
	QuoteItemId                       int       `json:"quote_item_id"`
	RowInvoiced                       float64   `json:"row_invoiced"`
	RowTotal                          float64   `json:"row_total"`
	RowTotalInclTax                   float64   `json:"row_total_incl_tax,omitempty"`
	RowWeight                         float64   `json:"row_weight"`
	Sku                               string    `json:"sku"`
	StoreId                           int       `json:"store_id"`
	TaxAmount                         float64   `json:"tax_amount"`
	TaxInvoiced                       float64   `json:"tax_invoiced"`
	TaxPercent                        float64   `json:"tax_percent"`
	UpdatedAt                         time.Time `json:"updated_at"`
	Weight                            float64   `json:"weight"`
}

type BillingAddress Address

type Payment struct {
	Id                    int64    `json:"id,omitempty"`
	OrderId               int64    `json:"order_id"`
	AccountStatus         string   `json:"account_status"`
	AdditionalInformation []string `json:"additional_information"`
	AmountOrdered         float64  `json:"amount_ordered"`
	BaseAmountOrdered     float64  `json:"base_amount_ordered"`
	BaseShippingAmount    float64  `json:"base_shipping_amount"`
	CcLast4               int      `json:"cc_last4"`
	EntityId              int64    `json:"entity_id"`
	Method                string   `json:"method"`
	ParentId              int      `json:"parent_id"`
	ShippingAmount        float64  `json:"shipping_amount"`
}

type ShippingAssignment struct {
	Shipping Shipping `json:"shipping"`
	Items    []Item   `json:"items"`
}

type Shipping struct {
	Address Address `json:"address"`
	Method  string  `json:"method"`
	Total   Total   `json:"total"`
}

type Total struct {
	BaseShippingAmount                    float64 `json:"base_shipping_amount"`
	BaseShippingDiscountAmount            float64 `json:"base_shipping_discount_amount"`
	BaseShippingInclTax                   float64 `json:"base_shipping_incl_tax"`
	BaseShippingTaxAmount                 float64 `json:"base_shipping_tax_amount"`
	ShippingAmount                        float64 `json:"shipping_amount"`
	ShippingDiscountAmount                float64 `json:"shipping_discount_amount"`
	ShippingDiscountTaxCompensationAmount float64 `json:"shipping_discount_tax_compensation_amount"`
	ShippingInclTax                       float64 `json:"shipping_incl_tax"`
	ShippingTaxAmount                     float64 `json:"shipping_tax_amount"`
}

type PlaceOrderData struct {
	UserId             string             `json:"user_id"`
	CartId             string             `json:"cart_id"`
	OrderId            string             `json:"order_id"`
	CreatedAt          string             `json:"created_at"`
	UpdatedAt          string             `json:"updated_at"`
	Transmited         bool               `json:"transmited"`
	TransmitedAt       string             `json:"transmited_at"`
	AddressInformation AddressInformation `json:"addressInformation"`
	Products           []Product          `json:"products"`
}

type AddressInformation struct {
	ShippingAddress         Address     `json:"shippingAddress"`
	BillingAddress          Address     `json:"billingAddress"`
	ShippingMethodCode      string      `json:"shipping_method_code"`
	ShippingCarrierCode     string      `json:"shipping_carrier_code"`
	ShippingExtraFields     interface{} `json:"shippingExtraFields"`
	PaymentMethodCode       string      `json:"payment_method_code"`
	PaymentMethodAdditional interface{} `json:"payment_method_additional"`
}

type Address struct {
	Id              int64
	CustomerId      int64    `json:"customer_id"`
	AddressType     string   `json:"address_type"`
	City            string   `json:"city"`
	Company         string   `json:"company"`
	CountryId       string   `json:"country_id"`
	Email           string   `json:"email"`
	Firstname       string   `json:"firstname"`
	Lastname        string   `json:"lastname"`
	Postcode        string   `json:"postcode"`
	Region          string   `json:"region"`
	RegionCode      string   `json:"region_code"`
	RegionId        int64    `json:"region_id"`
	Street          []string `json:"street"`
	Telephone       string   `json:"telephone"`
	StreetLine0     string   `json:"street_line_0,omitempty"`
	StreetLine1     string   `json:"street_line_1,omitempty"`
	OrderId         int64    `json:"order_id"`
	EntityId        string   `json:"entity_id"`
	DefaultShipping bool     `json:"default_shipping"`
	DefaultBilling  bool     `json:"default_billing"`
	ParentId        int64    `json:"parent_id"`
}

type Product struct {
	Pattern             string      `json:"pattern"`
	EcoCollection       string      `json:"eco_collection"`
	TierPrices          interface{} `json:"tier_prices"`
	Tsk                 int64       `json:"tsk"`
	CustomAttributes    interface{} `json:"custom_attributes"`
	SizeOptions         []int64     `json:"size_options"`
	RegularPrice        float64     `json:"regular_price"`
	FinalPrice          float64     `json:"final_price"`
	ErinRecommends      string      `json:"erin_recommends"`
	Price               float64     `json:"price"`
	ColorOptions        []int64     `json:"color_options"`
	Id                  string      `json:"id"`
	Sku                 string      `json:"sku"`
	Image               string      `json:"image"`
	New                 string      `json:"new"`
	Thumbnail           string      `json:"thumbnail"`
	Visibility          int64       `json:"visibility"`
	TypeId              string      `json:"type_id"`
	TaxClassId          int         `json:"tax_class_id"`
	Climate             []string      `json:"climate"`
	StyleGeneral        string      `json:"style_general"`
	UrlKey              string      `json:"url_key"`
	PerformanceFabric   string      `json:"performance_fabric"`
	Sale                string      `json:"sale"`
	MaxPrice            float64     `json:"max_price"`
	MinimalRegularPrice float64     `json:"minimal_regular_price"`
	RequiredOptions     []int       `json:"required_options"`
	Material            string      `json:"material"`
	SpecialPrice        int64       `json:"special_price"`
	MinimalPrice        float64     `json:"minimal_price"`
	Name                string      `json:"name"`
	MaxRegularPrice     float64     `json:"max_regular_price"`
	Status              int64       `json:"status"`
	PriceTax            float64     `json:"price_tax"`
	PriceInclTax        float64     `json:"priceInclTax"`
	SpecialPriceTax     float64     `json:"specialPriceTax"`
	SpecialPriceInclTax float64     `json:"specialPriceInclTax"`
	Sgn                 string      `json:"sgn"`
	Score               int64       `json:"_score"`
	Slug                string      `json:"slug"`
	Errors              interface{} `json:"errors"`
	Info                struct {
		Stock string `json:"stock"`
	} `json:"info"`
	ParentSku                  string        `json:"parentSku"`
	ProductOption              ProductOption `json:"product_option"`
	Qty                        float64       `json:"qty"`
	IsConfigured               bool          `json:"is_configured"`
	Color                      string        `json:"color"`
	SmallImage                 string        `json:"small_image"`
	HasOptions                 []int         `json:"has_options"`
	MsrpDisplayActualPriceType string        `json:"msrp_display_actual_price_type"`
	Size                       string        `json:"size"`
	OnlineStockCheckId         string        `json:"onlineStockCheckid"`
	IsInStock                  bool          `json:"is_in_stock"`
	ServerItemId               int64         `json:"server_item_id"`
	ServerCartId               string        `json:"server_cart_id"`
	PrevQty                    float64       `json:"prev_qty"`
	Totals                     Totals        `json:"totals"`
	Stock                      Stock         `json:"stock"`
}

type Stock struct {
	ItemId                         int64   `json:"item_id"`
	ProductId                      int64   `json:"product_id"`
	StockId                        int64   `json:"stock_id"`
	Qty                            float64 `json:"qty"`
	IsInStock                      bool    `json:"is_in_stock"`
	IsQtyDecimal                   bool    `json:"is_qty_decimal"`
	ShowDefaultNotificationMessage bool    `json:"show_default_notification_message"`
	UseConfigMinQty                bool    `json:"use_config_min_qty"`
	MinQty                         float64 `json:"min_qty"`
	UseConfigMinSaleQty            int64   `json:"use_config_min_sale_qty"`
	MinSaleQty                     float64 `json:"min_sale_qty"`
	UseConfigMaxSaleQty            bool    `json:"use_config_max_sale_qty"`
	MaxSaleQty                     float64 `json:"max_sale_qty"`
	UseConfigBackorders            bool    `json:"use_config_backorders"`
	Backorders                     int64   `json:"backorders"`
	UseConfigNotifyStockQty        bool    `json:"use_config_notify_stock_qty"`
	NotifyStockQty                 float64 `json:"notify_stock_qty"`
	UseConfigQtyIncrements         bool    `json:"use_config_qty_increments"`
	QtyIncrements                  float64 `json:"qty_increments"`
	UseConfigEnableQtyInc          bool    `json:"use_config_enable_qty_inc"`
	EnableQtyIncrements            bool    `json:"enable_qty_increments"`
	UseConfigManageStock           bool    `json:"use_config_manage_stock"`
	ManageStock                    bool    `json:"manage_stock"`
	LowStockDate                   string  `json:"low_stock_date"`
	IsDecimalDivided               bool    `json:"is_decimal_divided"`
	StockStatusChangeAuto          int64   `json:"stock_status_change_auto"`
}

type Totals struct {
	ItemId               int64    `json:"item_id"`
	Price                float64  `json:"price"`
	BasePrice            float64  `json:"base_price"`
	Qty                  float64  `json:"qty"`
	RowTotal             float64  `json:"row_total"`
	BaseRowTotal         float64  `json:"base_row_total"`
	RowTotalWithDiscount float64  `json:"row_total_with_discount"`
	TaxAmount            float64  `json:"tax_amount"`
	BaseTaxAmount        float64  `json:"base_tax_amount"`
	TaxPercent           float64  `json:"tax_percent"`
	DiscountAmount       float64  `json:"discount_amount"`
	BaseDiscountAmount   float64  `json:"base_discount_amount"`
	DiscountPercent      float64  `json:"discount_percent"`
	PriceInclTax         float64  `json:"price_incl_tax"`
	BasePriceInclTax     float64  `json:"base_price_incl_tax"`
	RowTotalInclTax      float64  `json:"row_total_incl_tax"`
	BaseRowTotalInclTax  float64  `json:"base_row_total_incl_tax"`
	Options              []Option `json:"options"`
	WeeeTaxAppliedAmount float64  `json:"weee_tax_applied_amount"`
	WeeeTaxApplied       float64  `json:"weee_tax_applied"`
	Name                 string   `json:"name"`
}

type ProductOption struct {
	ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
}

type ExtensionAttributes struct {
	ConfigurableItemOptions []ConfigurableItemOption `json:"configurable_item_options"`
}

type Option struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type ConfigurableItemOption struct {
	OptionId    string `json:"option_id"`
	OptionValue string `json:"option_value"`
}

type response struct {
	Code   int    `json:"code"`
	Result result `json:"result"`
}

type result struct {
	Items      []History `json:"items"`
	TotalCount int       `json:"total_count"`
}

func (order *History) SaveOrder() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	res, err := db.Exec("INSERT INTO `order` ("+
		"applied_rule_ids, "+
		"base_currency_code, "+
		"base_discount_amount, "+
		"base_grand_total, "+
		"base_discount_tax_compensation_amount, "+
		"base_shipping_amount, "+
		"base_shipping_discount_amount, "+
		"base_shipping_incl_tax, "+
		"base_shipping_tax_amount, "+
		"base_subtotal, "+
		"base_subtotal_incl_tax, "+
		"base_tax_amount, "+
		"base_total_due, "+
		"base_to_global_rate, "+
		"base_to_order_rate, "+
		"billing_address_id, "+
		"created_at, "+
		"customer_email, "+
		"customer_firstname, "+
		"customer_group_id, "+
		"customer_id, "+
		"customer_is_guest, "+
		"customer_lastname, "+
		"customer_note_notify, "+
		"discount_amount, "+
		"email_sent, "+
		"entity_id, "+
		"global_currency_code, "+
		"grand_total, "+
		"discount_tax_compensation_amount, "+
		"increment_id, "+
		"is_virtual, "+
		"order_currency_code, "+
		"protect_code, "+
		"quote_id, "+
		"shipping_amount, "+
		"shipping_description, "+
		"shipping_discount_amount, "+
		"shipping_discount_tax_compensation_amount, "+
		"shipping_incl_tax, "+
		"shipping_tax_amount, "+
		"state, "+
		"status, "+
		"store_currency_code, "+
		"store_id, "+
		"store_name, "+
		"store_to_base_rate, "+
		"store_to_order_rate, "+
		"subtotal, "+
		"subtotal_incl_tax, "+
		"tax_amount, "+
		"total_due, "+
		"total_item_count, "+
		"total_qty_ordered, "+
		"updated_at, "+
		"weight) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		order.AppliedRuleIds,
		order.BaseCurrencyCode,
		order.BaseDiscountAmount,
		order.BaseGrandTotal,
		order.BaseDiscountTaxCompensationAmount,
		order.BaseShippingAmount,
		order.BaseShippingDiscountAmount,
		order.BaseShippingInclTax,
		order.BaseShippingTaxAmount,
		order.BaseSubtotal,
		order.BaseSubtotalInclTax,
		order.BaseTaxAmount,
		order.BaseTotalDue,
		order.BaseToGlobalRate,
		order.BaseToOrderRate,
		order.BillingAddressId,
		order.CreatedAt,
		order.CustomerEmail,
		order.CustomerFirstname,
		order.CustomerGroupId,
		order.CustomerId,
		order.CustomerIsGuest,
		order.CustomerLastname,
		order.CustomerNoteNotify,
		order.DiscountAmount,
		order.EmailSent,
		order.EntityId,
		order.GlobalCurrencyCode,
		order.GrandTotal,
		order.DiscountTaxCompensationAmount,
		order.IncrementId,
		order.IsVirtual,
		order.OrderCurrencyCode,
		order.ProtectCode,
		order.QuoteId,
		order.ShippingAmount,
		order.ShippingDescription,
		order.ShippingDiscountAmount,
		order.ShippingDiscountTaxCompensationAmount,
		order.ShippingInclTax,
		order.ShippingTaxAmount,
		order.State,
		order.Status,
		order.StoreCurrencyCode,
		order.StoreId,
		order.StoreName,
		order.StoreToBaseRate,
		order.StoreToOrderRate,
		order.Subtotal,
		order.SubtotalInclTax,
		order.TaxAmount,
		order.TotalDue,
		order.TotalItemCount,
		order.TotalQtyOrdered,
		order.UpdatedAt,
		order.Weight)
	helpers.PanicErr(err)
	id, err := res.LastInsertId()
	helpers.PanicErr(err)
	order.ID = id
}

func GetAllCustomerOrderHistory(customerId int) (customerHistoryArray []History) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT * FROM `order` WHERE customer_id = ?", customerId)
	helpers.PanicErr(err)
	for rows.Next() {
		var customerHistory History
		if err := rows.Scan(
			&customerHistory.AppliedRuleIds,
			&customerHistory.BaseCurrencyCode,
			&customerHistory.BaseDiscountAmount,
			&customerHistory.BaseGrandTotal,
			&customerHistory.BaseDiscountTaxCompensationAmount,
			&customerHistory.BaseShippingAmount,
			&customerHistory.BaseShippingDiscountAmount,
			&customerHistory.BaseShippingInclTax,
			&customerHistory.BaseShippingTaxAmount,
			&customerHistory.BaseSubtotal,
			&customerHistory.BaseSubtotalInclTax,
			&customerHistory.BaseTaxAmount,
			&customerHistory.BaseTotalDue,
			&customerHistory.BaseToGlobalRate,
			&customerHistory.BaseToOrderRate,
			&customerHistory.BillingAddressId,
			&customerHistory.CreatedAt,
			&customerHistory.CustomerEmail,
			&customerHistory.CustomerFirstname,
			&customerHistory.CustomerGroupId,
			&customerHistory.CustomerId,
			&customerHistory.CustomerIsGuest,
			&customerHistory.CustomerLastname,
			&customerHistory.CustomerNoteNotify,
			&customerHistory.DiscountAmount,
			&customerHistory.EmailSent,
			&customerHistory.EntityId,
			&customerHistory.GlobalCurrencyCode,
			&customerHistory.GrandTotal,
			&customerHistory.DiscountTaxCompensationAmount,
			&customerHistory.IncrementId,
			&customerHistory.IsVirtual,
			&customerHistory.OrderCurrencyCode,
			&customerHistory.ProtectCode,
			&customerHistory.QuoteId,
			&customerHistory.ShippingAmount,
			&customerHistory.ShippingDescription,
			&customerHistory.ShippingDiscountAmount,
			&customerHistory.ShippingDiscountTaxCompensationAmount,
			&customerHistory.ShippingInclTax,
			&customerHistory.ShippingTaxAmount,
			&customerHistory.State,
			&customerHistory.Status,
			&customerHistory.StoreCurrencyCode,
			&customerHistory.StoreId,
			&customerHistory.StoreName,
			&customerHistory.StoreToBaseRate,
			&customerHistory.StoreToOrderRate,
			&customerHistory.Subtotal,
			&customerHistory.SubtotalInclTax,
			&customerHistory.TaxAmount,
			&customerHistory.TotalDue,
			&customerHistory.TotalItemCount,
			&customerHistory.TotalQtyOrdered,
			&customerHistory.UpdatedAt,
			&customerHistory.Weight,
			&customerHistory.ID); err != nil {
			helpers.PanicErr(err)
		}
		customerHistory.GetOrderItems()
		customerHistory.GetOrderPaymentData()
		customerHistory.GetOrderBillingAddress()
		customerHistory.GetOrderShippingAddress()
		customerHistoryArray = append(customerHistoryArray, customerHistory)
	}
	return
}

func GetOrder(orderId int) (order History) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM order WHERE id = ?", orderId).Scan(&order.ID,
		&order.AppliedRuleIds,
		&order.BaseCurrencyCode,
		&order.BaseDiscountAmount,
		&order.BaseGrandTotal,
		&order.BaseDiscountTaxCompensationAmount,
		&order.BaseShippingAmount,
		&order.BaseShippingDiscountAmount,
		&order.BaseShippingInclTax,
		&order.BaseShippingTaxAmount,
		&order.BaseSubtotal,
		&order.BaseSubtotalInclTax,
		&order.BaseTaxAmount,
		&order.BaseTotalDue,
		&order.BaseToGlobalRate,
		&order.BaseToOrderRate,
		&order.BillingAddressId,
		&order.CreatedAt,
		&order.CustomerEmail,
		&order.CustomerFirstname,
		&order.CustomerGroupId,
		&order.CustomerId,
		&order.CustomerIsGuest,
		&order.CustomerLastname,
		&order.CustomerNoteNotify,
		&order.DiscountAmount,
		&order.EmailSent,
		&order.GlobalCurrencyCode,
		&order.GrandTotal,
		&order.DiscountTaxCompensationAmount,
		&order.IncrementId,
		&order.IsVirtual,
		&order.OrderCurrencyCode,
		&order.ProtectCode,
		&order.QuoteId,
		&order.ShippingAmount,
		&order.ShippingDescription,
		&order.ShippingDiscountAmount,
		&order.ShippingDiscountTaxCompensationAmount,
		&order.ShippingInclTax,
		&order.ShippingTaxAmount,
		&order.State,
		&order.Status,
		&order.StoreCurrencyCode,
		&order.StoreId,
		&order.StoreName,
		&order.StoreToBaseRate,
		&order.StoreToOrderRate,
		&order.Subtotal,
		&order.SubtotalInclTax,
		&order.TaxAmount,
		&order.TotalDue,
		&order.TotalItemCount,
		&order.TotalQtyOrdered,
		&order.UpdatedAt,
		&order.Weight)
	helpers.PanicErr(err)
	return
}

func RemoveOrder(orderId int) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("DELETE FROM order WHERE id = ?", orderId)
	helpers.PanicErr(err)
}

func (order History) UpdateOrder() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	order.UpdatedAt = time.Now().UTC()
	res, err := db.Exec("UPDATE order o SET "+
		"o.applied_rule_ids = ?, "+
		"o.base_currency_code  = ?, "+
		"o.base_discount_amount = ?, "+
		"o.base_grand_total = ?, "+
		"o.base_discount_tax_compensation_amount = ?, "+
		"o.base_shipping_amount = ?, "+
		"o.base_shipping_discount_amount = ?, "+
		"o.base_shipping_incl_tax = ?, "+
		"o.base_shipping_tax_amount = ?, "+
		"o.base_subtotal = ?, "+
		"o.base_subtotal_incl_tax = ?, "+
		"o.base_tax_amount = ?, "+
		"o.base_total_due = ?, "+
		"o.base_to_global_rate = ?, "+
		"o.base_to_order_rate = ?, "+
		"o.billing_address_id = ?, "+
		"o.created_at = ?, "+
		"o.customer_email = ?, "+
		"o.customer_firstname = ?, "+
		"o.customer_group_id = ?, "+
		"o.customer_id = ?, "+
		"o.customer_is_guest = ?, "+
		"o.customer_lastname = ?, "+
		"o.customer_note_notify = ?, "+
		"o.discount_amount = ?, "+
		"o.email_sent = ?, "+
		"o.entity_id = ?, "+
		"o.global_currency_code = ?, "+
		"o.grand_total = ?, "+
		"o.discount_tax_compensation_amount = ?, "+
		"o.increment_id = ?, "+
		"o.is_virtual = ?, "+
		"o.order_currency_code = ?, "+
		"o.protect_code = ?, "+
		"o.quote_id = ?, "+
		"o.shipping_amount = ?, "+
		"o.shipping_description = ?, "+
		"o.shipping_discount_amount = ?, "+
		"o.shipping_discount_tax_compensation_amount = ?, "+
		"o.shipping_incl_tax = ?, "+
		"o.shipping_tax_amount = ?, "+
		"o.state = ?, "+
		"o.status = ?, "+
		"o.store_currency_code = ?, "+
		"o.store_id = ?, "+
		"o.store_name = ?, "+
		"o.store_to_base_rate = ?, "+
		"o.store_to_order_rate = ?, "+
		"o.subtotal = ?, "+
		"o.subtotal_incl_tax = ?, "+
		"o.tax_amount = ?, "+
		"o.total_due = ?, "+
		"o.total_item_count = ?, "+
		"o.total_qty_ordered = ?, "+
		"o.updated_at = ?, "+
		"o.weight = ? "+
		"WHERE o.id = ?",
		order.AppliedRuleIds,
		order.BaseCurrencyCode,
		order.BaseDiscountAmount,
		order.BaseGrandTotal,
		order.BaseDiscountTaxCompensationAmount,
		order.BaseShippingAmount,
		order.BaseShippingDiscountAmount,
		order.BaseShippingInclTax,
		order.BaseShippingTaxAmount,
		order.BaseSubtotal,
		order.BaseSubtotalInclTax,
		order.BaseTaxAmount,
		order.BaseTotalDue,
		order.BaseToGlobalRate,
		order.BaseToOrderRate,
		order.BillingAddressId,
		order.CreatedAt,
		order.CustomerEmail,
		order.CustomerFirstname,
		order.CustomerGroupId,
		order.CustomerId,
		order.CustomerIsGuest,
		order.CustomerLastname,
		order.CustomerNoteNotify,
		order.DiscountAmount,
		order.EmailSent,
		order.EntityId,
		order.GlobalCurrencyCode,
		order.GrandTotal,
		order.DiscountTaxCompensationAmount,
		order.IncrementId,
		order.IsVirtual,
		order.OrderCurrencyCode,
		order.ProtectCode,
		order.QuoteId,
		order.ShippingAmount,
		order.ShippingDescription,
		order.ShippingDiscountAmount,
		order.ShippingDiscountTaxCompensationAmount,
		order.ShippingInclTax,
		order.ShippingTaxAmount,
		order.State,
		order.Status,
		order.StoreCurrencyCode,
		order.StoreId,
		order.StoreName,
		order.StoreToBaseRate,
		order.StoreToOrderRate,
		order.Subtotal,
		order.SubtotalInclTax,
		order.TaxAmount,
		order.TotalDue,
		order.TotalItemCount,
		order.TotalQtyOrdered,
		order.UpdatedAt,
		order.Weight,
		order.ID)
	helpers.PanicErr(err)
	rowsAffected, err := res.RowsAffected()
	helpers.PanicErr(err)
	fmt.Println("Order ID: ", order.ID, " got ", rowsAffected, " rows updated")
}

func (item Item) SaveItem() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO order_items ("+
		"amount_refunded, "+
		"applied_rule_ids, "+
		"base_amount_refunded, "+
		"base_discount_amount, "+
		"base_discount_invoiced, "+
		"base_discount_tax_compensation_amount, "+
		"base_original_price, "+
		"base_price, "+
		"base_price_incl_tax, "+
		"base_row_invoiced, "+
		"base_row_total, "+
		"base_row_total_incl_tax, "+
		"base_tax_amount, "+
		"base_tax_invoiced, "+
		"created_at, "+
		"discount_amount, "+
		"discount_invoiced, "+
		"discount_percent, "+
		"free_shipping, "+
		"discount_tax_compensation_amount, "+
		"is_qty_decimal, "+
		"is_virtual, "+
		"name, "+
		"no_discount, "+
		"order_id, "+
		"original_price, "+
		"parent_item_id, "+
		"product_id, "+
		"product_type, "+
		"qty_canceled, "+
		"qty_invoiced, "+
		"qty_ordered, "+
		"qty_refunded, "+
		"qty_shipped, "+
		"quote_item_id, "+
		"row_invoiced, "+
		"row_total, "+
		"row_total_incl_tax, "+
		"row_weight, "+
		"sku, "+
		"store_id, "+
		"tax_amount, "+
		"tax_invoiced, "+
		"tax_percent, "+
		"updated_at, "+
		"weight) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		item.AmountRefunded,
		item.AppliedRuleIds,
		item.BaseAmountRefunded,
		item.BaseDiscountAmount,
		item.BaseDiscountInvoiced,
		item.BaseDiscountTaxCompensationAmount,
		item.BaseOriginalPrice,
		item.BasePrice,
		item.BasePriceInclTax,
		item.BaseRowInvoiced,
		item.BaseRowTotal,
		item.BaseRowTotalInclTax,
		item.BaseTaxAmount,
		item.BaseTaxInvoiced,
		item.CreatedAt,
		item.DiscountAmount,
		item.DiscountInvoiced,
		item.DiscountPercent,
		item.FreeShipping,
		item.BaseDiscountTaxCompensationAmount,
		item.IsQtyDecimal,
		item.IsVirtual,
		item.Name,
		item.NoDiscount,
		item.OrderId,
		item.OriginalPrice,
		item.ParentItemId,
		item.ProductId,
		item.ProductType,
		item.QtyCanceled,
		item.QtyInvoiced,
		item.QtyOrdered,
		item.QtyRefunded,
		item.QtyShipped,
		item.QuoteItemId,
		item.RowInvoiced,
		item.RowTotal,
		item.RowTotalInclTax,
		item.RowWeight,
		item.Sku,
		item.StoreId,
		item.TaxAmount,
		item.TaxInvoiced,
		item.TaxPercent,
		item.UpdatedAt,
		item.Weight)
	helpers.PanicErr(err)
}

func (paymentData Payment) SavePaymentData(orderId int64) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO payment ("+
		"account_status, "+
		"amount_ordered, "+
		"base_amount_ordered, "+
		"base_shipping_amount, "+
		"cc_last4, "+
		"entity_id, "+
		"method, "+
		"parent_id ,"+
		"shipping_amount, "+
		"order_id, "+
		"additional_information) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		paymentData.AccountStatus,
		paymentData.AmountOrdered,
		paymentData.BaseAmountOrdered,
		paymentData.BaseShippingAmount,
		paymentData.CcLast4,
		paymentData.EntityId,
		paymentData.Method,
		paymentData.ParentId,
		paymentData.ShippingAmount,
		orderId,
		paymentData.AdditionalInformation[0])
	helpers.PanicErr(err)

}

func (order *History) GetOrderPaymentData() {
	fmt.Println("order id - ", order.ID)
	order.Payment.AdditionalInformation = make([]string, 2)
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM payment WHERE order_id = ?", order.ID).
		Scan(
			&order.Payment.Id,
			&order.Payment.AccountStatus,
			&order.Payment.AmountOrdered,
			&order.Payment.BaseAmountOrdered,
			&order.Payment.BaseShippingAmount,
			&order.Payment.CcLast4,
			&order.Payment.EntityId,
			&order.Payment.Method,
			&order.Payment.ParentId,
			&order.Payment.ShippingAmount,
			&order.Payment.OrderId,
			&order.Payment.AdditionalInformation[0])
	helpers.PanicErr(err)
}

func (address *Address) SaveOrderShippingAddress(orderId int64) {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	_, err = db.Exec("INSERT INTO order_shipping_address ("+
		"city, "+
		"company, "+
		"country_id, "+
		"email, "+
		"firstname, "+
		"lastname, "+
		"postcode, "+
		"region, "+
		"region_code, "+
		"region_id, "+
		"telephone, "+
		"street_line_0, "+
		"street_line_1, "+
		"address_type, "+
		"entity_id, "+
		"parent_id) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		address.City,
		address.Company,
		address.CountryId,
		address.Email,
		address.Firstname,
		address.Lastname,
		address.Postcode,
		address.Region,
		address.RegionCode,
		address.RegionId,
		address.Telephone,
		address.Street[0],
		address.Street[1],
		"shipping",
		address.EntityId,
		orderId)
	helpers.PanicErr(err)
}

func (order *History) GetOrderItems() {
	order.Items = []Item{}
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	rows, err := db.Query("SELECT * FROM order_items WHERE order_id = ?", order.ID)
	helpers.PanicErr(err)
	for rows.Next() {
		var item Item
		if err := rows.Scan(
			&item.AmountRefunded,
			&item.AppliedRuleIds,
			&item.BaseAmountRefunded,
			&item.BaseDiscountAmount,
			&item.BaseDiscountInvoiced,
			&item.BaseDiscountTaxCompensationAmount,
			&item.BaseOriginalPrice,
			&item.BasePrice,
			&item.BasePriceInclTax,
			&item.BaseRowInvoiced,
			&item.BaseRowTotal,
			&item.BaseRowTotalInclTax,
			&item.BaseTaxAmount,
			&item.BaseTaxInvoiced,
			&item.CreatedAt,
			&item.DiscountAmount,
			&item.DiscountInvoiced,
			&item.DiscountPercent,
			&item.FreeShipping,
			&item.DiscountTaxCompensationAmount,
			&item.IsQtyDecimal,
			&item.IsVirtual,
			&item.Name,
			&item.NoDiscount,
			&item.OrderId,
			&item.OriginalPrice,
			&item.ParentItemId,
			&item.ProductId,
			&item.ProductType,
			&item.QtyCanceled,
			&item.QtyInvoiced,
			&item.QtyOrdered,
			&item.QtyRefunded,
			&item.QtyShipped,
			&item.QuoteItemId,
			&item.RowInvoiced,
			&item.RowTotal,
			&item.RowTotalInclTax,
			&item.RowWeight,
			&item.Sku,
			&item.StoreId,
			&item.TaxAmount,
			&item.TaxInvoiced,
			&item.TaxPercent,
			&item.UpdatedAt,
			&item.Weight,
			&item.ItemId); err != nil {
			helpers.PanicErr(err)
		}
		order.Items = append(order.Items, item)
	}
}

func (order *History) GetOrderBillingAddress() {
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM addresses WHERE id = ?", order.BillingAddressId).
		Scan(
			&order.BillingAddress.Id,
			&order.BillingAddress.CustomerId,
			&order.BillingAddress.RegionId,
			&order.BillingAddress.CountryId,
			&order.BillingAddress.Telephone,
			&order.BillingAddress.Postcode,
			&order.BillingAddress.City,
			&order.BillingAddress.Firstname,
			&order.BillingAddress.Lastname,
			&order.BillingAddress.DefaultShipping,
			&order.BillingAddress.StreetLine0,
			&order.BillingAddress.StreetLine1,
			&order.BillingAddress.DefaultBilling,
			&order.BillingAddress.Email)
	helpers.PanicErr(err)
	order.BillingAddress.Street = formatStreet(order.BillingAddress.StreetLine0, order.BillingAddress.StreetLine1)
}

func formatStreet(line0 string, line1 string) []string {
	lines := []string{line0, line1}
	return lines
}

func (order *History) GetOrderShippingAddress() {
	var sa ShippingAssignment
	db, err := config.Conf.GetDb()
	helpers.PanicErr(err)
	err = db.QueryRow("SELECT * FROM order_shipping_address WHERE parent_id = ?", order.ID).
		Scan(&sa.Shipping.Address.Id,
			&sa.Shipping.Address.City,
			&sa.Shipping.Address.Company,
			&sa.Shipping.Address.CountryId,
			&sa.Shipping.Address.Email,
			&sa.Shipping.Address.Firstname,
			&sa.Shipping.Address.Lastname,
			&sa.Shipping.Address.Postcode,
			&sa.Shipping.Address.Region,
			&sa.Shipping.Address.RegionCode,
			&sa.Shipping.Address.RegionId,
			&sa.Shipping.Address.Telephone,
			&sa.Shipping.Address.StreetLine0,
			&sa.Shipping.Address.StreetLine1,
			&sa.Shipping.Address.AddressType,
			&sa.Shipping.Address.EntityId,
			&sa.Shipping.Address.ParentId)
	helpers.PanicErr(err)
	sa.Shipping.Address.Street = formatStreet(sa.Shipping.Address.StreetLine0, sa.Shipping.Address.StreetLine1)
	order.ExtensionAttributes.ShippingAssignments = append(order.ExtensionAttributes.ShippingAssignments, sa)
}
