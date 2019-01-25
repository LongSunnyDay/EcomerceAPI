package solrDataContrler

import (
	"../helpers"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func cmdDeleteAll() {
	deleteUrl := "http://localhost:8983/solr/vue_storefront_core/update?commit=true"
	deleteCmd := "<delete><query>*:*</query></delete>"
	resp, err := http.Post(deleteUrl, "text/xml", strings.NewReader(deleteCmd))
	defer helpers.CloseRows(resp.Body)
	helpers.PanicErr(err)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("success response code: %d %v", resp.StatusCode, string(bodyBytes))
}
