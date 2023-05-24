package http

import "net/http"

type HTTPError struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	Internal   string `json:"internal"`
	StatusCode int    `json:"statusCode"`
}

var (
	StatusError         = "ERROR"
	DefaultErrorMessage = "Sedang terjadi kendala pada server kami. Silahkan coba beberapa saat lagi"
)

var (
	ErrAutoComplete = NewHTTPError(http.StatusInternalServerError, StatusError, DefaultErrorMessage, StatusError)
)

// NewHTTPError ...
func NewHTTPError(statusCode int, code, message, internal string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Internal:   internal,
	}
}

func (h *HTTPError) Error() string {
	return h.Message
}
