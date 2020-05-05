package ctrls

import (
	"mt-scale/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOrderList Get order list
func GetOrderList(ctx *gin.Context) {
	orderList := models.FetchOrder()

	ctx.JSON(http.StatusOK, orderList)
}
