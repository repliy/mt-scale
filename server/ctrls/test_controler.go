package ctrls

import (
	"mt-scale/common"
	"mt-scale/models"
	"mt-scale/models/dto"
	"mt-scale/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

// WriteExcelFile Test path: test/excel
func WriteExcelFile(ctx *gin.Context) interface{} {
	var dto dto.QueryTallyBoxDto = dto.QueryTallyBoxDto{
		TaskID: "5ec117aefa8ed46e85f52223",
	}
	tallyData := models.GetTallyInfo(dto)
	utils.WriteToExcelFile(tallyData)
	return "success"
}

//Redis测试
func RedisSetAction(ctx *gin.Context) {

	rds := common.CreateRedisPool().Get()

	count, _ := redis.Int(rds.Do("GET", "count"))
	count++
	rds.Do("SET", "count", count)
	ctx.JSON(200, gin.H{
		"message": count,
	})

}

//Sesssion测试
func SessionAction(ctx *gin.Context) {

	session := sessions.Default(ctx)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	session.Set("count", count)
	session.Save()
	ctx.JSON(200, gin.H{"count": count})

}
