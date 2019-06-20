package response

// SuccessBody holds data for success response
type SuccessBody struct {
	Data   interface{} `json:"data,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ErrorBody holds data for error response
type ErrorBody struct {
	Errors ErrorInfo   `json:"errors"`
	Status interface{} `json:"status"`
}

// ErrorInfo holds detail information about error
type ErrorInfo struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Source  string `json:"source"`
}

// NewErrorInfo to create single error info in Errors
func NewErrorInfo(field, message, source string) ErrorInfo {
	return ErrorInfo{
		Field:   field,
		Message: message,
		Source:  source,
	}
}

// BuildSuccess is a function to create SuccessBody
func BuildSuccess(data interface{}, stat interface{}) SuccessBody {
	return SuccessBody{
		Data:   data,
		Status: stat,
	}
}

// BuildError is a function to create ErrorBody
func BuildError(errors ErrorInfo, stat interface{}) ErrorBody {
	return ErrorBody{
		Errors: errors,
		Status: stat,
	}
}
