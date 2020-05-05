package main

import (
	"mt-scale/route"
	"mt-scale/syslog"
	"mt-scale/task"
	"mt-scale/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	syslog.Debug("Application Starting ...")

	task.Start()

	gin.SetMode(gin.DebugMode)
	r := route.SetupRouter()
	port := utils.GetConfStr("service.port")
	r.Run(port)
}
