package main

import (
	"mt-scale/common"
	"mt-scale/route"
	"mt-scale/syslog"
	"mt-scale/task"

	"github.com/gin-gonic/gin"
)

func main() {
	syslog.Debug("Application Starting ...")

	task.Start()

	gin.SetMode(gin.DebugMode)
	r := route.SetupRouter()
	port := common.GetConfStr("service.port")
	r.Run(port)
}
