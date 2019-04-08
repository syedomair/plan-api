package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/go-kit/kit/log"
)

func PostNotificationToHttpBin(logger log.Logger, url string, jsonStr map[string]string, chJob chan string, chRresult chan string) {

	start := time.Now()
	logger.Log("METHOD", "PostNotificationToHttpBin", "SPOT", "method start", "time_start", start)
	for s := range chJob {
		time.Sleep(2 * time.Second)
		logger.Log("METHOD", "PostNotificationToHttpBin", "SPOT", "Start", "Processing", s)
		requestBody, _ := json.Marshal(jsonStr)
		resp, _ := http.Post("https://httpbin.org/"+url, "application/json", bytes.NewBuffer(requestBody))
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		if strings.Contains(string(body), "data") {
			chRresult <- "success"
		}
		logger.Log("METHOD", "PostNotificationToHttpBin", "SPOT", "End", "Processing", s)
	}
	logger.Log("METHOD", "PostNotificationToHttpBin", "SPOT", "method end", "time_spent", time.Since(start))
}
