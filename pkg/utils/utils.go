package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Parse the ody of the http request:
func ParseBody(rq *http.Request, x interface{}) {
	if body, err := io.ReadAll(rq.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			log.Fatal(err)
			return
		}
	}
}