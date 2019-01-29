package attribute

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-ws/helpers"
	"io/ioutil"
	"net/http"
)

// Apache Solr client function
func GetAttributeNameFromSolr(attributeId string, attributeValue string) ItemAttribute {
	request := map[string]interface{}{
		"query":  "_type:attribute",
		"filter": "id:" + attributeId,
		"fields": "*, [child parentFilter=_type:attribute childFilter=value:" + attributeValue + " limit=30]"}
	requestBytes := new(bytes.Buffer)
	err := json.NewEncoder(requestBytes).Encode(request)
	helpers.PanicErr(err)
	resp, err := http.Post(
		SolrQueryUrl,
		ContentType,
		requestBytes)
	helpers.PanicErr(err)
	b, _ := ioutil.ReadAll(resp.Body)
	var solrResp solrResponse
	err = json.Unmarshal(b, &solrResp)
	helpers.PanicErr(err)
	fmt.Println( solrResp.Response.NumFound)

	if solrResp.Response.NumFound == 1 {
		itemAttribute := ItemAttribute{
			Name:  solrResp.Response.Docs[0].AttributeCode,
			Label: solrResp.Response.Docs[0].ChildDocuments[0].Label}
		return itemAttribute
	}
	return ItemAttribute{}
}