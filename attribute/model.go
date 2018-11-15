package attribute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-ws/helpers"
	"io/ioutil"
	"net/http"
)

type attribute struct {
	Type                      string        `json:"_type"`
	Id                        string        `json:"id"`
	Score                     int           `json:"_score"`
	IsWysiwygEnabled          bool          `json:"is_wysiwyg_enabled"`
	IsHTMLAllowedOnFront      bool          `json:"is_html_allowed_on_front"`
	UsedForSortBy             bool          `json:"used_for_sort_by"`
	IsFilterable              bool          `json:"is_filterable"`
	IsFilterableInSearch      bool          `json:"is_filterable_in_search"`
	IsUsedInGrid              bool          `json:"is_used_in_grid"`
	IsVisibleInGrid           bool          `json:"is_visible_in_grid"`
	IsFilterableInGrid        bool          `json:"is_filterable_in_grid"`
	Position                  int           `json:"position"`
	ApplyTo                   []interface{} `json:"apply_to"`
	IsSearchable              string        `json:"is_searchable"`
	IsVisibleInAdvancedSearch string        `json:"is_visible_in_advanced_search"`
	IsComparable              string        `json:"is_comparable"`
	IsUsedForPromoRules       string        `json:"is_used_for_promo_rules"`
	IsVisibleOnFront          string        `json:"is_visible_on_front"`
	UsedInProductListing      string        `json:"used_in_product_listing"`
	IsVisible                 bool          `json:"is_visible"`
	Scope                     string        `json:"scope"`
	AttributeID               int           `json:"attribute_id"`
	AttributeCode             string        `json:"attribute_code"`
	FrontendInput             string        `json:"frontend_input"`
	EntityTypeID              string        `json:"entity_type_id"`
	IsRequired                bool          `json:"is_required"`
	IsUserDefined             bool          `json:"is_user_defined"`
	DefaultFrontendLabel      string        `json:"default_frontend_label"`
	BackendType               string        `json:"backend_type"`
	BackendModel              string        `json:"backend_model"`
	SourceModel               string        `json:"source_model"`
	DefaultValue              string        `json:"default_value"`
	IsUnique                  string        `json:"is_unique"`
	ID                        int           `json:"id"`
	Tsk                       int64         `json:"tsk"`
}

func getAttributeFromSolr(attributeId string)  {
	request := map[string]interface{}{
		"query":  "_type:attribute",
		"filter": attributeId}
	requestBytes := new(bytes.Buffer)
	json.NewEncoder(requestBytes).Encode(request)
	resp, err := http.Post(
		"http://localhost:8983/solr/storefrontCore/query",
		"application/json; charset=utf-8",
		requestBytes)
	helpers.PanicErr(err)
	b, _ := ioutil.ReadAll(resp.Body)
	var attribute attribute
	json.Unmarshal(b, &attribute)
	fmt.Println(attribute)
}
