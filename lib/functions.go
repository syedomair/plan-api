package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func PostNotificationToHttpBin(url string, jsonStr map[string]string, resultCh chan string) {

	requestBody, _ := json.Marshal(jsonStr)
	resp, _ := http.Post("https://httpbin.org/"+url, "application/json", bytes.NewBuffer(requestBody))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if strings.Contains(string(body), "data") {
		resultCh <- "success"
	}
}
