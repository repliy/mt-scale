package ctrls

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/syslog"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

// CreateFishTag path: tag/crt
func CreateFishTag(ctx *gin.Context) interface{} {
	var tag entitys.FishTag
	if err := ctx.ShouldBindJSON(&tag); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(tag)
	return models.AddTag(tag)
}
