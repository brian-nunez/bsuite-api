package errors

type ErrorType string

const (
	ErrInvalidRequest      ErrorType = "INVALID_REQUEST"
	ErrUnauthorized        ErrorType = "UNAUTHORIZED"
	ErrNotFound            ErrorType = "NOT_FOUND"
	ErrNotAllowed          ErrorType = "NOT_ALLOWED"
	ErrInternalServerError ErrorType = "INTERNAL_SERVER_ERROR"
	ErrServiceUnavailable  ErrorType = "SERVICE_UNAVAILABLE"
)

type ErrorMessage struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type ErrorResponse struct {
	HTTPStatusCode int          `json:"status"`
	ErrorMessage   ErrorMessage `json:"error"`
}

type errorBuilder struct {
	httpStatusCode int
	errorCode      string
	message        string
}

func (b *errorBuilder) WithStatusCode(code int) *errorBuilder {
	b.httpStatusCode = code
	return b
}

func (b *errorBuilder) WithMessage(message string) *errorBuilder {
	b.message = message
	return b
}

func (b *errorBuilder) WithErrorCode(errorCode string) *errorBuilder {
	b.errorCode = errorCode
	return b
}

func (b *errorBuilder) Build() *ErrorResponse {
	return &ErrorResponse{
		HTTPStatusCode: b.httpStatusCode,
		ErrorMessage: ErrorMessage{
			ErrorCode:    b.errorCode,
			ErrorMessage: b.message,
		},
	}
}

func Custom() *errorBuilder {
	return &errorBuilder{}
}

func InvalidRequest() *errorBuilder {
	return &errorBuilder{
		httpStatusCode: 400,
		errorCode:      string(ErrInvalidRequest),
		message:        "Invalid Request",
	}
}

func Unauthorized() *errorBuilder {
	return &errorBuilder{
		httpStatusCode: 401,
		errorCode:      string(ErrUnauthorized),
		message:        "Invalid Request",
	}
}

func NotFound() *errorBuilder {
	return &errorBuilder{
		httpStatusCode: 404,
		errorCode:      string(ErrNotFound),
		message:        "Invalid Request",
	}
}

func NotAllowed() *errorBuilder {
	return &errorBuilder{
		httpStatusCode: 405,
		errorCode:      string(ErrNotAllowed),
		message:        "Invalid Request",
	}
}

func InternalServerError() *errorBuilder {
	return &errorBuilder{
		httpStatusCode: 500,
		errorCode:      string(ErrInternalServerError),
		message:        "Invalid Request",
	}
}

func ServiceNotAvailable() *errorBuilder {
	return &errorBuilder{
		httpStatusCode: 503,
		errorCode:      string(ErrServiceUnavailable),
		message:        "Invalid Request",
	}
}

func GenerateByStatusCode(code int) *errorBuilder {
	switch code {
	case 400:
		return InvalidRequest()
	case 401:
		return Unauthorized()
	case 404:
		return NotAllowed()
	case 405:
		return NotAllowed()
	case 500:
		return InternalServerError()
	case 503:
		return ServiceNotAvailable()
	}

	return InternalServerError()
}
