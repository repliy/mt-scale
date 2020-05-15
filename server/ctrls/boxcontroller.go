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

// CreateBox Add a new box
func CreateBox(ctx *gin.Context) interface{} {
	var box dto.AddBoxDto
	if err := ctx.ShouldBindJSON(&box); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(box)
	return models.AddBox(box)
}

// CreateBoxNum path: box/boxlist/add
func CreateBoxNum(ctx *gin.Context) interface{} {
	var list dto.AddBoxListDto
	if err := ctx.ShouldBindJSON(&list); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(list)
	return ""
}

// GetBoxByType Query a new box
func GetBoxByType(ctx *gin.Context) interface{} {
	var param dto.QueryBoxDto
	if err := ctx.ShouldBindQuery(&param); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(param)
	return models.FetchBoxes(param.BoxType)
}

