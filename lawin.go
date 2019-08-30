package lawin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TransformStrToMap : To string to json map
// jsonstr - json String
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

// HTTPGet : http GET Request
// URL - http URL
// queryString - query string params
// headers - request headers
func HTTPGet(URL string, queryString map[string]interface{},
	headers map[string]string) *http.Response {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	if queryString != nil {
		q := req.URL.Query()
		for key, value := range queryString {
			strValue, _ := value.(string)
			q.Add(key, strValue)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP GET Failed")
	}
	return resp
}

// HTTPPost : HTTP Post request
// URL - http URL
// headers - request headers
func HTTPPost(URL string, payload map[string]interface{},
	headers map[string]string) *http.Response {
	client := &http.Client{}

	jsonPayload, _ := json.Marshal(payload)
	jsonString := string(jsonPayload)

	fmt.Println(jsonString)

	req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonPayload))
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP POST Failed")
	}

	return resp
}
