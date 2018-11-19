package product

import "go-api-ws/attribute"

func BuildSKUFromItemAttributes(itemAttributes []attribute.ItemAttribute, sku string) (string) {

	for _, attribute := range itemAttributes {
		if attribute.Name == "size" {
			sku = sku + "-" + attribute.Value
		}
	}

	for _, attribute := range itemAttributes {
		if attribute.Name == "color" {
			sku = sku + "-" + attribute.Value
		}
	}
	return sku
}
