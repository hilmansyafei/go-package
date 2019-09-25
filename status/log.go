// Package status is used to write status based on Canopus standard
package status

import "net/http"

// Log holds data for log info
type Log struct {
	IP       string      `json:"ip"`
	Protocol string      `json:"protocol"`
	Host     string      `json:"host"`
	URI      string      `json:"uri"`
	Headers  http.Header `json:"headers"`
}
