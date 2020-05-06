package route

import (
	"mt-scale/common"
	"mt-scale/ctrls"
	"mt-scale/exception"
	"mt-scale/middleware/jwt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router Gin router
var Router *gin.Engine

// ResultHandlerFunc Controller return result data handler
type ResultHandlerFunc func(c *gin.Context) *common.ResultData

func wrapper(handler ResultHandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := handler(c)
		c.JSON(result.HTTPCode, result)
	}
}

func init() {
	Router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*", "jwt"}
	config.ExposeHeaders = []string{"jwt"}

	Router.Use(cors.New(config))

	Router.Use(jwt.Middleware(jwt.Options{}))

	Router.Use(exception.MiddleWare())
}

// SetupRouter Setup path
func SetupRouter() *gin.Engine {

	// static
	staticPath := common.GetConfStr("router.static")
	Router.Static("/static", staticPath)

	// tpl
	viewPath := common.GetConfStr("router.view")
	Router.LoadHTMLGlob(viewPath)

	// 404
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", "")
	})

	// order paths
	orderRouter := Router.Group("/order")
	orderRouter.GET("/list", wrapper(ctrls.GetOrderList))

	return Router
}
