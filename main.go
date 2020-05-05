package main

import (
	"mt-scale/models"
	"mt-scale/route"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {

	gin.SetMode(gin.DebugMode)
	r := route.SetupRouter()

	//定时程序启动
	c := cron.New()
	//数据库状态检查
	c.AddFunc("0/10 * * * * ?", models.MongoDbCheck)
	c.Start()

	port := utils.GetConfStr("service.port")
	r.Run(port)
}
