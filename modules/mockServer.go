package modules

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// GenerateMockServer is function to generate mock server with,
// Custom statusCode, headers, and body.
func GenerateMockServer(statusCode int, headersParam interface{}, body interface{}) *httptest.Server {
	mockserver := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
			if headersParam != "" {
				headers := headersParam.(map[string]string)
				for k, v := range headers {
					w.Header().Set(k, v)
				}
			}
			json.NewEncoder(w).Encode(body)
		}))
	return mockserver
}

// CloseMockServer is function to close connection current,
// mock server.
func CloseMockServer(mockserver *httptest.Server) {
	defer mockserver.Close()
}
