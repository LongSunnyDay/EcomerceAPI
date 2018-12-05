package order

import (
	"time"
)

type History struct {
	AppliedRuleIds                        string    `json:"applied_rule_ids"`
	BaseCurrencyCode                      string    `json:"base_currency_code"`
	BaseDiscountAmount                    float64   `json:"base_discount_amount"`
	BaseGrandTotal                        float64   `json:"base_grand_total"`
	BaseDiscountTaxCompensationAmount     float64   `json:"base_discount_tax_compensation_amount"`
	BaseShippingAmount                    float64   `json:"base_shipping_amount"`
	BaseShippingDiscountAmount            float64   `json:"base_shipping_discount_amount"`
	BaseShippingInclTax                   float64   `json:"base_shipping_incl_tax"`
	BaseShippingTaxAmount                 float64   `json:"base_shipping_tax_amount"`
	BaseSubtotal                          float64   `json:"base_subtotal"`
	BaseSubtotalInclTax                   float64   `json:"base_subtotal_incl_tax"`
	BaseTaxAmount                         float64   `json:"base_tax_amount"`
	BaseTotalDue                          float64   `json:"base_total_due"`
	BaseToGlobalRate                      float64   `json:"base_to_global_rate"`
	BaseToOrderRate                       float64   `json:"base_to_order_rate"`
	BillingAddressId                      int       `json:"billing_address_id"`
	CreatedAt                             time.Time `json:"created_at"`
	CustomerEmail                         string    `json:"customer_email"`
	CustomerFirstname                     string    `json:"customer_firstname"`
	CustomerGroupId                       int       `json:"customer_group_id"`
	CustomerId                            int       `json:"customer_id"`
	CustomerIsGuest                       int       `json:"customer_is_guest"`
	CustomerLastname                      string    `json:"customer_lastname"`
	CustomerNoteNotify                    int       `json:"customer_note_notify"`
	DiscountAmount                        float64   `json:"discount_amount"`
	EmailSent                             int       `json:"email_sent"`
	EntityId                              int       `json:"entity_id"`
	GlobalCurrencyCode                    string    `json:"global_currency_code"`
	GrandTotal                            float64   `json:"grand_total"`
	DiscountTaxCompensationAmount         float64   `json:"discount_tax_compensation_amount"`
	IncrementId                           string    `json:"increment_id"`
	IsVirtual                             int       `json:"is_virtual"`
	OrderCurrencyCode                     string    `json:"order_currency_code"`
	ProtectCode                           string    `json:"protect_code"`
	QuoteId                               int       `json:"quote_id"`
	ShippingAmount                        float64   `json:"shipping_amount"`
	ShippingDescription                   string    `json:"shipping_description"`
	ShippingDiscountAmount                float64   `json:"shipping_discount_amount"`
	ShippingDiscountTaxCompensationAmount float64   `json:"shipping_discount_tax_compensation_amount"`
	ShippingInclTax                       float64   `json:"shipping_incl_tax"`
	ShippingTaxAmount                     float64   `json:"shipping_tax_amount"`
	State                                 string    `json:"state"`
	Status                                string    `json:"status"`
	StoreCurrencyCode                     string    `json:"store_currency_code"`
	StoreId                               int       `json:"store_id"`
	StoreName                             string    `json:"store_name"`
	StoreToBaseRate                       float64   `json:"store_to_base_rate"`
	StoreToOrderRate                      float64   `json:"store_to_order_rate"`
	Subtotal                              float64   `json:"subtotal"`
	SubtotalInclTax                       float64   `json:"subtotal_incl_tax"`
	TaxAmount                             float64   `json:"tax_amount"`
	TotalDue                              float64   `json:"total_due"`
	TotalItemCount                        float64   `json:"total_item_count"`
	TotalQtyOrdered                       float64   `json:"total_qty_ordered"`
	UpdatedAt                             time.Time `json:"updated_at"`
	Weight                                float64   `json:"weight"`
	Items                                 []Item    `json:"items"`
	BillingAddress                        Address   `json:"billing_address"`
	Payment                               Payment   `json:"payment"`
	StatusHistories                       []string  `json:"status_histories"`
	ExtensionAttributes                   struct {
		ShippingAssignments []ShippingAssignment `json:"shipping_assignments"`
	} `json:"extension_attributes"`
}

type Item struct {
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
	ParentItem                        Item      `json:"parent_item,omitempty"`
}

type Address struct {
	AddressType string   `json:"address_type"`
	City        string   `json:"city"`
	CountryId   string   `json:"country_id"`
	Email       string   `json:"email"`
	EntityId    int      `json:"entity_id"`
	Firstname   string   `json:"firstname"`
	Lastname    string   `json:"lastname"`
	ParentId    int      `json:"parent_id"`
	Postcode    string   `json:"postcode"`
	Street      []string `json:"street"`
	Telephone   string
}

type Payment struct {
	AccountStatus         string   `json:"account_status"`
	AdditionalInformation []string `json:"additional_information"`
	AmountOrdered         float64  `json:"amount_ordered"`
	BaseAmountOrdered     float64  `json:"base_amount_ordered"`
	BaseShippingAmount    float64  `json:"base_shipping_amount"`
	CcLast4               int      `json:"cc_last4"`
	EntityId              int      `json:"entity_id"`
	Method                string   `json:"method"`
	ParentId              int      `json:"parent_id"`
	ShippingAmount        float64  `json:"shipping_amount"`
}

type ShippingAssignment struct {
	Shipping struct {
		Address Address `json:"address"`
		Method  string  `json:"method"`
		Total   struct {
			BaseShippingAmount                    float64 `json:"base_shipping_amount"`
			BaseShippingDiscountAmount            float64 `json:"base_shipping_discount_amount"`
			BaseShippingInclTax                   float64 `json:"base_shipping_incl_tax"`
			BaseShippingTaxAmount                 float64 `json:"base_shipping_tax_amount"`
			ShippingAmount                        float64 `json:"shipping_amount"`
			ShippingDiscountAmount                float64 `json:"shipping_discount_amount"`
			ShippingDiscountTaxCompensationAmount float64 `json:"shipping_discount_tax_compensation_amount"`
			ShippingInclTax                       float64 `json:"shipping_incl_tax"`
			ShippingTaxAmount                     float64 `json:"shipping_tax_amount"`
		} `json:"total"`
	} `json:"shipping"`
	Items []Item `json:"items"`
}
