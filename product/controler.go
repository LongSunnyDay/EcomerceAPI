package product

import "go-api-ws/attribute"

func BuildSKUFromItemAttributes(itemAttributes []attribute.ItemAttribute, sku string) (string) {

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
