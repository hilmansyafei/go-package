package response

type JsonRespFormat struct {
	Message string                 `json:"message,omitempty"`
	Code    int                    `json:"code,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
