package exception

import (
	"mt-scale/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MtException Mt server internal exception struct
type MtException struct {
	HTTPCode  int    `json:"-"`
	ErrorCode int    `json:"code"`
	Msg       string `json:"msg"`
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
				c.JSON(mtException.HTTPCode, mtException)
				c.Abort()
			}
		}()
		c.Next()
	}
}

// Error Implement error interface
func (e *MtException) Error() string {
	return e.Msg
}

// ThrowBusinessError Internal business system exception
func ThrowBusinessError(code int) {
	msg := common.StatusText(code)
	err := &MtException{
		HTTPCode:  http.StatusOK,
		ErrorCode: code,
		Msg:       msg,
	}
	panic(err)
}

// ThrowBusinessErrorMsg Throw exception with custom message
func ThrowBusinessErrorMsg(messgae string) {
	err := &MtException{
		HTTPCode:  http.StatusOK,
		ErrorCode: common.BusinessErrorCode,
		Msg:       messgae,
	}
	panic(err)
}

func unknownError(message string) *MtException {
	return &MtException{
		HTTPCode:  http.StatusForbidden,
		ErrorCode: common.UnknownErrorCode,
		Msg:       message,
	}
}

func serverError() *MtException {
	return &MtException{
		HTTPCode:  http.StatusInternalServerError,
		ErrorCode: common.ServerErrorCode,
		Msg:       http.StatusText(http.StatusInternalServerError),
	}
}
