package category

import (
	"go-api-ws/core"
)

type Category struct {
	Id string
	Code string
	Slug string
	Title string
	Description string
}

var categories []Category

var cotegoryModule core.ApiModule

func init() {
	cotegoryModule = core.ApiModule{
		Name: "Category module",
		Description: "Category module. Support unlimited count of categories. Categories are stored as a flat list.",
		Version: "0.1",
		Author: "Remigijus Bauzys @ JivaLabs",
	}
}

func addRoutes() {

}

func initCategories() {

}
