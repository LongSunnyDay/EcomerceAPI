package product

import "go-api-ws/attribute"

func BuildSKUFromItemAttributes(itemAttributes []attribute.ItemAttribute, sku string) (string) {

	for _, itemAttribute := range itemAttributes {
		if itemAttribute.Name == "size" {
			sku = sku + "-" + itemAttribute.Value
		}
	}

	for _, itemAttribute := range itemAttributes {
		if itemAttribute.Name == "color" {
			sku = sku + "-" + itemAttribute.Value
		}
	}
	return sku
}
