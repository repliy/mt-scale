package common

import (
	"encoding/json"
	"net/http"
)

// ResultData Http request return data struct
type ResultData struct {
	HTTPCode int    `json:"-"`
	Data     string `json:"data"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}

// RetData Struct http return data
func RetData(obj interface{}) *ResultData {
	data, err := json.Marshal(obj)
	if err != nil {
		panic("json marshal error")
	}
	return &ResultData{
		HTTPCode: http.StatusOK,
		Data:     string(data),
		Code:     BusinessSuccessCode,
		Msg:      StatusText(BusinessSuccessCode),
	}
}
