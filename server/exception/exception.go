package exception

import (
	"mt-scale/common"
	"mt-scale/syslog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MtException Mt server internal exception struct
type MtException struct {
	HTTPCode  int    `json:"-"`
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
}

// Error Implement error interface
func (e *MtException) Error() string {
	return e.Message
}

// New New Create
func New(code int, msg string) *MtException {
	return &MtException{
		HTTPCode:  http.StatusOK,
		ErrorCode: code,
		Message:   msg,
	}
}

// MiddleWare Gin panic middleware
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var mtException *MtException
				if h, ok := err.(*MtException); ok {
					mtException = h
				} else if e, ok := err.(error); ok {
					mtException = unknownError(e.Error())
				} else {
					mtException = serverError()
				}
				syslog.Debug("bad request end:", time.Now())
				syslog.Error(mtException, mtException.HTTPCode)
				c.JSON(mtException.HTTPCode, mtException)
				c.Abort()
			}
		}()
		c.Next()
	}
}

// ThrowBusinessError Internal business system exception
func ThrowBusinessError(code int) {
	msg := common.StatusText(code)
	err := &MtException{
		HTTPCode:  http.StatusOK,
		ErrorCode: code,
		Message:   msg,
	}
	panic(err)
}

// ThrowBusinessErrorMsg Throw exception with custom message
func ThrowBusinessErrorMsg(messgae string) {
	err := &MtException{
		HTTPCode:  http.StatusOK,
		ErrorCode: common.BusinessErrorCode,
		Message:   messgae,
	}
	panic(err)
}

func unknownError(message string) *MtException {
	return &MtException{
		HTTPCode:  http.StatusForbidden,
		ErrorCode: common.UnknownErrorCode,
		Message:   message,
	}
}

func serverError() *MtException {
	return &MtException{
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: common.ServerErrorCode,
		Message:   http.StatusText(http.StatusInternalServerError),
	}
}
