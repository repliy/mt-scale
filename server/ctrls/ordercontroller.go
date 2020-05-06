package ctrls

import (
	"mt-scale/common"
	"mt-scale/models"

	"github.com/gin-gonic/gin"
)

// GetOrderList Get order list
func GetOrderList(ctx *gin.Context) *common.ResultData {
	orderList := models.FetchOrder()
	return common.RetData(orderList)
}
