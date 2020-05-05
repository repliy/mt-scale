package route

import (
	"mt-scale/ctrls"
	"mt-scale/middleware/jwt"
	"mt-scale/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router Gin router
var Router *gin.Engine

func init() {
	Router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*", "jwt"}
	config.ExposeHeaders = []string{"jwt"}

	Router.Use(cors.New(config))

	Router.Use(jwt.Middleware(jwt.Options{
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))
}

// SetupRouter Setup path
func SetupRouter() *gin.Engine {

	//静态目录配置
	staticPath := utils.GetConfStr("router.static")
	Router.Static("/static", staticPath)

	//模板
	viewPath := utils.GetConfStr("router.view")
	Router.LoadHTMLGlob(viewPath)

	//Session测试-Redis存储
	Router.GET("/session", ctrls.SessionAction)

	//Redis测试
	Router.GET("/redis", ctrls.RedisSetAction)

	Router.GET("/queue", ctrls.QueueAction)

	//获取用户订单
	Router.GET("/orders", ctrls.GetOrderList)

	Router.POST("/protected", func(c *gin.Context) {
		c.String(200, "CSRF token is valid")
	})

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	return Router
}
