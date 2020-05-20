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

// AddWeightRecord path: record/crt {}
func AddWeightRecord(ctx *gin.Context) interface{} {
	var dto dto.AddWeightRecordDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(dto)
	return models.AddWeightRecord(dto)
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

// UpdateWeightRecord path: record/upd
func UpdateWeightRecord(ctx *gin.Context) interface{} {
	var record dto.UpdWeightRecordDto
	if err := ctx.ShouldBindJSON(&record); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(record)
	return models.UpdWeightRecord(record)
}

// DelWeightRecord path: record/del
func DelWeightRecord(ctx *gin.Context) interface{} {
	var dto dto.DelRecordDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(dto)
	models.DeleteWeightRecord(dto)
	return "success"
}
