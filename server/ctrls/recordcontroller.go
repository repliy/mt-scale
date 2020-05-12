package ctrls

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

// AddWeightRecord upload one weight record
func AddWeightRecord(ctx *gin.Context) interface{} {
	var record entitys.WeightRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(record)
	return models.AddWeightRecord(record)
}
