package exception

import (
	"net/http"
)

const (
	// logicErrorCode Server internal exception code
	businessErrorCode = 200101
	unknownErrorCode  = 200102
	serverErrorCode   = 200103
)

var statusText = map[int]string{
	businessErrorCode: "Mt Internal System Error ...",
	unknownErrorCode:  "",
	serverErrorCode:   "",
}

// MtException Mt server internal exception struct
type MtException struct {
	HTTPCode  int    `json:"-"`
	ErrorCode int    `json:"code"`
	Msg       string `json:"msg"`
}

// Error Implement error interface
func (e *MtException) Error() string {
	return e.Msg
}

// BusinessError Internal business system exception
func BusinessError() *MtException {
	return &MtException{
		HTTPCode:  http.StatusOK,
		ErrorCode: businessErrorCode,
		Msg:       statusText[businessErrorCode],
	}
}

// UnknownError Non-internal system exception
func UnknownError(message string) *MtException {
	return &MtException{
		HTTPCode:  http.StatusForbidden,
		ErrorCode: unknownErrorCode,
		Msg:       message,
	}
}

// ServerError Other system exception
func ServerError() *MtException {
	return &MtException{
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: serverErrorCode,
		Msg:       http.StatusText(http.StatusInternalServerError),
	}
}
