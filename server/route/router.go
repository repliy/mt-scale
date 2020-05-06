package route

import (
	"mt-scale/ctrls"
	"mt-scale/exception"
	"mt-scale/middleware/jwt"
	"mt-scale/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Router Gin router
var Router *gin.Engine

// ResultHandlerFunc Controller return result data handler
type ResultHandlerFunc func(c *gin.Context) (interface{}, error)

// ResultData Http request return data struct
type ResultData struct {
	data string
	code string
	msg  string
}

func wrapper(handler ResultHandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := handler(c)

		var mtException *exception.MtException

		if err != nil {
			if h, ok := err.(*exception.MtException); ok {
				mtException = h
			} else if e, ok := err.(error); ok {
				mtException = exception.UnknownError(e.Error())
			} else {
				mtException = exception.ServerError()
			}
			c.JSON(mtException.HTTPCode, mtException)
			return
		}
	}
}

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
	orderRouter.GET("/list", wrapper(ctrls.GetOrderList))

	return Router
}
