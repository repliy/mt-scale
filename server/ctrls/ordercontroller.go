package ctrls

import (
	"mt-scale/exception"
	"mt-scale/models"

	"github.com/gin-gonic/gin"
)

// GetOrderList Get order list
func GetOrderList(ctx *gin.Context) (interface{}, error) {
	orderList := models.FetchOrder()
	return orderList, exception.BusinessError()
}
