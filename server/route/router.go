package route

import (
	"mt-scale/common"
	"mt-scale/ctrls"
	"mt-scale/exception"
	"mt-scale/middleware/jwt"
	"mt-scale/syslog"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Router Gin router
// var Router *gin.Engine

// ResultData Http request return data struct
type ResultData struct {
	HTTPCode int         `json:"-"`
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
}

// ResultHandlerFunc Controller return result data handler
type ResultHandlerFunc func(c *gin.Context) interface{}

func wrapper(handler ResultHandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		syslog.Debug("request start:", time.Now())
		result := handler(c)
		if result == "file" {
			syslog.Info("request file download")
			return
		}
		retData := &ResultData{
			HTTPCode: http.StatusOK,
			Data:     result,
			Code:     common.BusinessSuccessCode,
			Message:  common.StatusText(common.BusinessSuccessCode),
		}
		syslog.Debug("request end:", time.Now())
		c.JSON(retData.HTTPCode, retData)
	}
}

// SetupRouter Setup path
func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*", "jwt"}
	config.ExposeHeaders = []string{"jwt"}

	router.Use(cors.New(config))
	// exception middleware need to before others
	router.Use(exception.MiddleWare())

	router.Use(jwt.Middleware(jwt.Options{}))

	router.Use(static.Serve("/", static.LocalFile("./dist/", true)))

	// // static
	// staticPath := common.GetConfStr("router.static")
	// router.Static("/static", staticPath)

	// // tpl
	// viewPath := common.GetConfStr("router.view")
	// router.LoadHTMLGlob(viewPath)

	// 404
	// router.NoRoute(func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusNotFound, "404.html", "")
	// })

	// box
	boxRouter := router.Group("/api/box")
	boxRouter.POST("/crt", wrapper(ctrls.CreateBox))
	boxRouter.POST("/crtboxes", wrapper(ctrls.CreateBoxList))
	boxRouter.GET("/fetchbytype", wrapper(ctrls.GetBoxByType))
	boxRouter.GET("/latestbox", wrapper(ctrls.GetLatestBoxes))
	boxRouter.GET("/tally", wrapper(ctrls.GetVesselPlantTallyInfo))

	// species
	speciesRouter := router.Group("/api/species")
	speciesRouter.POST("/crt", wrapper(ctrls.CreateSpecies))
	speciesRouter.GET("/fetchall", wrapper(ctrls.GetAllSpecies))

	// tag
	tagRouter := router.Group("/api/tag")
	tagRouter.POST("/crt", wrapper(ctrls.CreateFishTag))

	// record
	recordRouter := router.Group("/api/record")
	recordRouter.POST("/crt", wrapper(ctrls.AddWeightRecord))
	recordRouter.GET("/fetch", wrapper(ctrls.GetWeightRecord))
	recordRouter.POST("/upd", wrapper(ctrls.UpdateWeightRecord))
	recordRouter.POST("/del", wrapper(ctrls.DelWeightRecord))

	// stat
	statRouter := router.Group("/api/stat")
	statRouter.GET("/weight", wrapper(ctrls.StatWeight))

	// task
	taskRouter := router.Group("/api/task")
	taskRouter.GET("/latest", wrapper(ctrls.GetCurrentTask))
	taskRouter.POST("/status/upd", wrapper(ctrls.UpdTaskStatus))

	// user
	userRouter := router.Group("/api/user")
	userRouter.POST("/add", wrapper(ctrls.AddUser))
	userRouter.POST("/login", wrapper(ctrls.Login))

	// test
	testRouter := router.Group("/api/test")
	testRouter.GET("/excel", wrapper(ctrls.WriteExcelFile))

	testRouter.GET("/api/local/file", func(c *gin.Context) {
		c.File("test.xlsx")
	})

	return router
}
