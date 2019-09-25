package modules

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/sepulsa/canopus-gope/response"
	"github.com/sepulsa/canopus-gope/status"
)

// GenerateMockServer is function to generate mock server with,
// Custom statusCode, headers, and body.
func GenerateMockServer(statusCode int, headersParam interface{}, body interface{}) *httptest.Server {
	// Generate mock server.
	mockserver := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set Http status code.
			w.WriteHeader(statusCode)
			// Set headers if any.
			if headersParam != "" {
				headers := headersParam.(map[string]string)
				for k, v := range headers {
					w.Header().Set(k, v)
				}
			}
			// Generate body json.
			json.NewEncoder(w).Encode(body)
		}))
	return mockserver
}

// CloseMockServer is function to close connection current,
// mock server.
func CloseMockServer(mockserver *httptest.Server) {
	defer mockserver.Close()
}

// InitMuxServer is initiate function to generate multiple mock.
func InitMuxServer() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

// CreateMocMuxServerAPI is mock for api with no spesific query param.
func CreateMocMuxServerAPI(muxserver *http.ServeMux, path string, headersParam interface{}, body interface{}) {
	muxserver.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		// Set Http status code.
		w.WriteHeader(200)
		// Create mock header.
		// Set headers if any.
		if headersParam != "" {
			headers := headersParam.(map[string]string)
			for k, v := range headers {
				w.Header().Set(k, v)
			}
		}
		// Create mock body.
		// Generate body json.
		json.NewEncoder(w).Encode(body)
	})
}

// CreateMockGetPathByID is initiate function to generate multiple mock.
func CreateMockGetPathByID(muxserver *http.ServeMux, path string, headersParam interface{}, publicBody interface{}, privateBody interface{}) {
	muxserver.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		// Set Http status code.
		w.WriteHeader(200)
		// Create mock header.
		// Set headers if any.
		if headersParam != "" {
			headers := headersParam.(map[string]string)
			for k, v := range headers {
				w.Header().Set(k, v)
			}
		}
		query := r.URL.Query()
		switch typeQuery := query.Get("type"); typeQuery {
		case "private":
			// Generate body json.
			json.NewEncoder(w).Encode(privateBody)
		case "public":
			// Create mock body.
			json.NewEncoder(w).Encode(publicBody)
			// Generate body json.
		default:
			sErr := response.BuildError(response.NewErrorInfo(
				"Mock Error",
				"There is no type",
				"modules/mockServer.go"), status.BadRequestError)
			// Generate body json.
			json.NewEncoder(w).Encode(sErr)
		}
	})
}

// GenerateMockMuxServer is function to generate mock mux server.
func GenerateMockMuxServer(muxserver *http.ServeMux) *httptest.Server {
	ts := httptest.NewServer(muxserver)
	return ts
}
