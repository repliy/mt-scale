package ctrls

import (
	"mt-scale/models"
	"mt-scale/models/vo"

	"github.com/gin-gonic/gin"
)

// StatWeight path: stat/weight
func StatWeight(ctx *gin.Context) interface{} {
	species := models.StatSpecieszWeight()
	box := models.StatBoxWeight()
	total := models.StatTotalWeight()

	return vo.StatWeightVo{
		Species: species,
		Box:     box,
		Total:   total,
	}
}
