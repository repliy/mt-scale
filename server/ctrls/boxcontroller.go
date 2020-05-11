package ctrls

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/models/dto"
	"mt-scale/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateBox Add a new box
func CreateBox(ctx *gin.Context) interface{} {
	var box entitys.Box
	if err := ctx.ShouldBindJSON(&box); err != nil {
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	if err := utils.ValidateStructParams(box); err != nil {
		exception.ThrowBusinessErrorMsg(err.Error())
	}
	timeNow := time.Now()
	box.CreateTime = timeNow
	box.UpdateTime = timeNow

	boxID := models.AddBox(box)
	return boxID
}

// GetBoxByType Query a new box
func GetBoxByType(ctx *gin.Context) interface{} {
	var param dto.QueryBoxDto
	if err := ctx.ShouldBindJSON(&param); err != nil {
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	if err := utils.ValidateStructParams(param); err != nil {
		exception.ThrowBusinessErrorMsg(err.Error())
	}

	return models.FetchBoxes(param.BoxType)
}
