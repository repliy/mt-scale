package ctrls

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/models/dto"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

// CreateBox Add a new box
func CreateBox(ctx *gin.Context) interface{} {
	var box entitys.Box
	if err := ctx.ShouldBindJSON(&box); err != nil {
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(box)
	boxID := models.AddBox(box)
	return boxID
}

// GetBoxByType Query a new box
func GetBoxByType(ctx *gin.Context) interface{} {
	var param dto.QueryBoxDto
	if err := ctx.ShouldBindQuery(&param); err != nil {
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(param)

	return models.FetchBoxes(param.BoxType)
}
