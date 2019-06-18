// Package status is used to write status based on Sedekahnesia standard
package status

// Status holds data for status info
type Status struct {
	Type    string `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewSuccess create success status
func NewSuccess(code int, message string) Status {
	return Status{
		Type:    "Success",
		Code:    code,
		Message: message,
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
	OKSuccess      = NewSuccess(200200, "Request Successed")
	CreatedSuccess = NewSuccess(200201, "The request has been fulfilled")

	// List Of Error (Request)
	BadRequestError           = NewError(400400, "Bad Request")
	UnauthorizedError         = NewError(400401, "Unauthorized to access this endpoint resource")
	ForbiddenError            = NewError(400402, "Forbidden to access this endpoint resource")
	APINotFoundError          = NewError(400404, "The endpoint resource is not found")
	MethodNotAllowedError     = NewError(400405, "This method type is not currently supported for this endpoint resource")
	NotAcceptableError        = NewError(400406, "Acceptance header is invalid for this endpoint resource.")
	RequestTimeoutError       = NewError(400408, "Request timeout")
	UnsupportedMediaTypeError = NewError(400415, "The endpoint resource does not support the media type provided")
	MissingArgumentsError     = NewError(400419, "The endpoint resource is missing required arguments")
	InvalidArgumentsError     = NewError(400420, "The endpoint resource does not support one or more of the given parameters")
	UnprocessableEntityError  = NewError(400422, "The request body is not appropriate")

	// List Of Error (Server)
	InternalServerError = NewError(500500, "Internal server error")
)
