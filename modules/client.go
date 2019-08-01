package modules

import (
	"io"
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
func GenReqWithHeader(method string, url string, headers map[string]string, timeLimit time.Duration, body io.Reader) (*http.Response, error) {
	timeout := time.Duration(timeLimit * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, body)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	return resp, err
}
