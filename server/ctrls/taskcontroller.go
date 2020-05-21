package ctrls

import (
	"mt-scale/common"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/models/dto"
	"mt-scale/syslog"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

// GetCurrentTask path: task/current
func GetCurrentTask(ctx *gin.Context) interface{} {
	return models.GetCurrentTask()
}

// UpdTaskStatus path: task/status/upd
func UpdTaskStatus(ctx *gin.Context) interface{} {
	var dto dto.TaskUpdDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(dto)
	models.UpdTaskStatus(dto)
	return "success"
}
