package attribute

const (
	SolrQueryUrl = "http://192.168.0.2:8983/solr/storefrontCore/query"
	ContentType  = "application/json; charset=utf-8"
)

type (
	attribute struct {
		Type                      string        `json:"_type"`
		Id                        int           `json:"id"`
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
		IsSearchable              bool          `json:"is_searchable"`
		IsVisibleInAdvancedSearch string        `json:"is_visible_in_advanced_search"`
		IsComparable              string        `json:"is_comparable"`
		IsUsedForPromoRules       string        `json:"is_used_for_promo_rules"`
		IsVisibleOnFront          string        `json:"is_visible_on_front"`
		UsedInProductListing      string        `json:"used_in_product_listing"`
		IsVisible                 bool          `json:"is_visible"`
		Scope                     string        `json:"scope"`
		AttributeID               string        `json:"attribute_id"`
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
		ChildDocuments            []options     `json:"_childDocuments_"`
	}

	options struct {
		TypeOf   string `json:"type_of"`
		ParentId int    `json:"parent_id"`
		Label    string `json:"label"`
		Value    string `json:"value"`
	}

	solrResponse struct {
		ResponseHeader struct {
			Status int `json:"status"`
			QTime  int `json:"QTime"`
			Params struct {
				JSON string `json:"json"`
			}
		} `json:"responseHeader"`
		Response struct {
			NumFound int         `json:"numFound"`
			Start    int         `json:"start"`
			Docs     []attribute `json:"docs"`
		} `json:"response"`
	}

	ItemAttribute struct {
		Name  string
		Label string
	}
)


