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

// AddUser path: user/add
func AddUser(ctx *gin.Context) interface{} {
	var user entitys.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(user)
	return models.AddUser(user)
}

// Login path: user/login
func Login(ctx *gin.Context) interface{} {
	var dto dto.LoginDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		syslog.Error(err)
		exception.ThrowBusinessError(common.JSONFormatErrorCode)
	}
	utils.ValidateStructParams(dto)
	user, tokenStr := models.Login(dto)
	ctx.Header("jwt", tokenStr)
	return user
}
