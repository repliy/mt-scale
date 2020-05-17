package ctrls

import (
	"mt-scale/common"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/models/dto"
	"mt-scale/models/vo"
	"mt-scale/syslog"
	"mt-scale/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

// StatWeight path: stat/weight
func StatWeight(ctx *gin.Context) interface{} {
	var stat dto.StatDto
	if err := ctx.ShouldBindQuery(&stat); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(stat)
	tashBsonID, _ := primitive.ObjectIDFromHex(stat.TaskID)

	species := models.StatSpecieszWeight(tashBsonID)
	box := models.StatBoxWeight(tashBsonID)
	total := models.StatTotalWeight(tashBsonID)

	return vo.StatWeightVo{
		Species: species,
		Box:     box,
		Total:   total,
	}
}
