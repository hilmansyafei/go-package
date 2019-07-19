// Package status is used to write status based on Sedekahnesia standard
package status

// Status holds data for status info
type Status struct {
	Type    string                 `json:"type"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// NewSuccess create success status
func NewSuccess(code int, message string) Status {
	return Status{
		Type:    "Success",
		Code:    code,
		Message: message,
	}
}

// NewSuccessData create success status
func NewSuccessData(code int, message string, data map[string]interface{}) Status {
	return Status{
		Type:    "Success",
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// NewError create error status
func NewError(code int, message string) Status {
	return Status{
		Type:    "Error",
		Code:    code,
		Message: message,
	}
}

var (
	// List Of Success
	OKSuccess      = NewSuccess(200, "Request Successed")
	CreatedSuccess = NewSuccess(201, "The request has been fulfilled")
	DataNotFound   = NewSuccess(404, "Data Not Found")

	// List Of Error (Request)
	BadRequestError           = NewError(400, "Bad Request")
	UnauthorizedError         = NewError(401, "Unauthorized to access this endpoint resource")
	ForbiddenError            = NewError(402, "Forbidden to access this endpoint resource")
	APINotFoundError          = NewError(404, "The endpoint resource is not found")
	MethodNotAllowedError     = NewError(405, "This method type is not currently supported for this endpoint resource")
	NotAcceptableError        = NewError(406, "Acceptance header is invalid for this endpoint resource.")
	RequestTimeoutError       = NewError(408, "Request timeout")
	UnsupportedMediaTypeError = NewError(415, "The endpoint resource does not support the media type provided")
	MissingArgumentsError     = NewError(419, "The endpoint resource is missing required arguments")
	InvalidArgumentsError     = NewError(420, "The endpoint resource does not support one or more of the given parameters")
	UnprocessableEntityError  = NewError(422, "The request body is not appropriate")

	// List Of Error (Server)
	InternalServerError = NewError(500, "Internal server error")
)
