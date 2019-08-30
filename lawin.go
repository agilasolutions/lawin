package lawin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TransformStrToMap : To string to json map
// jsonstr - Json String
func TransformStrToMap(jsonstr string) map[string]interface{} {
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonstr), &jsonMap)
	return jsonMap
}

// ExtractRespBody : Extract response body and transform to map
// resp - http response
func ExtractRespBody(resp *http.Response) string {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Payload Transform Fail")
	}
	defer resp.Body.Close()
	bodyString := string(bodyBytes)
	return bodyString
}
