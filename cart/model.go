package cart

import (
	"time"
)

const (
	collectionName = "cart"
)

type (
	Cart struct {
		CartId    string    `json:"cart_id" bson:"cart_id"`
		QuoteId   int64     `json:"quote_id" bson:"quote_id"`
		UserId    string    `json:"user_id,omitempty" bson:"user_id,omitempty"`
		Items     []Item    `json:"items" bson:"items"`
		CreatedAt time.Time `json:"created_at,omitempty" bson:"createdAt,omitempty"`
		Status    string    `json:"status" bson:"status"`
	}

	CustomerCart struct {
		Item Item `json:"cartItem,omitempty" bson:"cartItem"`
	}

	Item struct {
		SKU           string  `json:"sku,omitempty" bson:"sku"`
		QTY           float64 `json:"qty,omitempty" bson:"qty"`
		Price         float64 `json:"price,omitempty" bson:"price"`
		ProductType   string  `json:"product_type,omitempty" bson:"product_type"`
		Name          string  `json:"name,omitempty" bson:"name"`
		ItemID        int     `json:"item_id,omitempty" bson:"item_id,omitempty"`
		QuoteId       string  `json:"quoteId,omitempty" bson:"quoteId"`
		ProductOption struct {
			ExtensionAttributes struct {
				ConfigurableItemOptions []Options `json:"configurable_item_options,omitempty" bson:"configurable_item_options"`
			} `json:"extension_attributes,omitempty" bson:"extension_attributes"`
		} `json:"product_option,omitempty" bson:"product_options"`
	}

	Options struct {
		OptionsID   string `json:"option_id,omitempty" bson:"option_id"`
		OptionValue string `json:"option_value,omitempty" bson:"option_value"`
	}
)

