package solrDataContrler

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"go-api-ws/helpers"
)

func cmdDeleteAll (){
	deleteUrl := "http://localhost:8983/solr/vue_storefront_core/update?commit=true"
	deleteCmd := "<delete><query>*:*</query></delete>"
	resp, err := http.Post(deleteUrl, "text/xml", strings.NewReader(deleteCmd))
	defer resp.Body.Close()
	helpers.CheckErr(err)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("success response code: %d %v", resp.StatusCode, string(bodyBytes))
}
