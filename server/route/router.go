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

// ResultData Http request return data struct
type ResultData struct {
	HTTPCode int         `json:"-"`
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Msg      string      `json:"msg"`
}

// ResultHandlerFunc Controller return result data handler
type ResultHandlerFunc func(c *gin.Context) interface{}

func wrapper(handler ResultHandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		result := handler(c)
		retData := &ResultData{
			HTTPCode: http.StatusOK,
			Data:     result,
			Code:     common.BusinessSuccessCode,
			Msg:      common.StatusText(common.BusinessSuccessCode),
		}
		c.JSON(retData.HTTPCode, retData)
	}
}

func init() {
	Router = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*", "jwt"}
	config.ExposeHeaders = []string{"jwt"}

	Router.Use(cors.New(config))
	// exception middleware need to before others
	Router.Use(exception.MiddleWare())

	Router.Use(jwt.Middleware(jwt.Options{}))
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

	// box
	boxRouter := Router.Group("/box")
	boxRouter.POST("/crt", wrapper(ctrls.CreateBox))
	boxRouter.GET("/fetchbytype", wrapper(ctrls.GetBoxByType))

	return Router
}
