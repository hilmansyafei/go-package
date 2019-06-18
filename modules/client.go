package modules

import (
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
