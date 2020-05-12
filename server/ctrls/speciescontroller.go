package ctrls

import (
	"mt-scale/common"
	"mt-scale/entitys"
	"mt-scale/exception"
	"mt-scale/models"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

// CreateSpecies path:/species/crt { name: "lobster", "tag": "#1000"}
func CreateSpecies(ctx *gin.Context) interface{} {
	var spec entitys.Species
	if err := ctx.ShouldBindJSON(&spec); err != nil {
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(spec)
	specID := models.AddSpecies(spec)
	return specID
}

// GetAllSpecies path:/species/fetchall
func GetAllSpecies(ctx *gin.Context) interface{} {
	return models.FetchAllSpeciesInfo()
}
