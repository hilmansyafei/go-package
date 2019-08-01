package modules

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// GenClient : Generate client timeout
func GenClient(timeLimit time.Duration) http.Client {
	timeout := time.Duration(timeLimit * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	return client
}

// GenReqWithHeader : Generate request with header
func GenReqWithHeader(method string, url string, headers map[string]string, timeLimit time.Duration, body map[string]interface{}) (*http.Response, error) {
	timeout := time.Duration(timeLimit * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	bodyMarshal, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyMarshal))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	return resp, err
}
