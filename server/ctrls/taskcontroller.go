package ctrls

import (
	"mt-scale/models"

	"github.com/gin-gonic/gin"
)

// GetCurrentTask path: task/current
func GetCurrentTask(ctx *gin.Context) interface{} {
	return models.GetCurrentTask()
}
