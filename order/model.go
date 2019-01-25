package order

import (
	"time"
)

type (
	History struct {
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

	ExtensionAttribute struct {
		ShippingAssignments []ShippingAssignment `json:"shipping_assignments"`
	}

	Item struct {
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

	ParentItem struct {
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

	BillingAddress Address

	Payment struct {
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

	ShippingAssignment struct {
		Shipping Shipping `json:"shipping"`
		Items    []Item   `json:"items"`
	}

	Shipping struct {
		Address Address `json:"address"`
		Method  string  `json:"method"`
		Total   Total   `json:"total"`
	}

	Total struct {
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

	PlaceOrderData struct {
		UserId             string             `json:"user_id"`
		CartId             string             `json:"cart_id"`
		OrderId            string             `json:"order_id"`
		CreatedAt          string             `json:"created_at"`
		UpdatedAt          string             `json:"updated_at"`
		Transmited         bool               `json:"transmited"`
		TransmitedAt       string             `json:"transmited_at"`
		AddressInformation AddressInformation `json:"addressInformation"`
		Products           []Product          `json:"products"`
		PersonalData       PersonalData       `json:"personalDetails"`
	}

	PersonalData struct {
		Email     string `json:"emailAddress"`
		Firstname string `json:"firstName"`
		Lastname  string `json:"lastName"`
	}

	AddressInformation struct {
		ShippingAddress         Address     `json:"shippingAddress"`
		BillingAddress          Address     `json:"billingAddress"`
		ShippingMethodCode      string      `json:"shipping_method_code"`
		ShippingCarrierCode     string      `json:"shipping_carrier_code"`
		ShippingExtraFields     interface{} `json:"shippingExtraFields"`
		PaymentMethodCode       string      `json:"payment_method_code"`
		PaymentMethodAdditional interface{} `json:"payment_method_additional"`
	}

	Address struct {
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

	Product struct {
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
		Climate             []string    `json:"climate"`
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

	Stock struct {
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

	Totals struct {
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

	ProductOption struct {
		ExtensionAttributes ExtensionAttributes `json:"extension_attributes"`
	}

	ExtensionAttributes struct {
		ConfigurableItemOptions []ConfigurableItemOption `json:"configurable_item_options"`
	}

	Option struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}

	ConfigurableItemOption struct {
		OptionId    string `json:"option_id"`
		OptionValue string `json:"option_value"`
	}

	response struct {
		Code   int    `json:"code"`
		Result result `json:"result"`
	}

	result struct {
		Items      []History `json:"items"`
		TotalCount int       `json:"total_count"`
	}
)

