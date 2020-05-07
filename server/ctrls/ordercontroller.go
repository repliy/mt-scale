package ctrls

import (
	"mt-scale/models"

	"github.com/gin-gonic/gin"
)

// GetOrderList Get order list
func GetOrderList(ctx *gin.Context) interface{} {
	orderList := models.FetchOrder()
	return orderList
}
