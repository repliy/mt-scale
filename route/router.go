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

	// static
	staticPath := utils.GetConfStr("router.static")
	Router.Static("/static", staticPath)

	// tpl
	viewPath := utils.GetConfStr("router.view")
	Router.LoadHTMLGlob(viewPath)

	// 404
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	// order paths
	orderRouter := Router.Group("/order")
	orderRouter.GET("/list", ctrls.GetOrderList)

	return Router
}
