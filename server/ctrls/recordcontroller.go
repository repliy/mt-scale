package ctrls

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/models/dto"
	"mt-scale/syslog"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

// AddWeightRecord path: record/crt {}
func AddWeightRecord(ctx *gin.Context) interface{} {
	var record entitys.WeightRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(record)
	return models.AddWeightRecord(record)
}

// GetWeightRecord path: record/fetch?box_id=&species_id=&index=&page_num=&page_size=
func GetWeightRecord(ctx *gin.Context) interface{} {
	var params dto.QueryRecordDto
	if err := ctx.ShouldBindQuery(&params); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	return models.FetchWeightRecord(params)
}

// StatWeightRecord path:/record/stat
func StatWeightRecord(ctx *gin.Context) interface{} {
	return models.StatSpecieszWeight()
}
