package httperrors

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message string) *HttpError {
	return &HttpError{
		Code:    code,
		Message: message,
	}
}

func (e *HttpError) Error() string {
	return e.Message
}
