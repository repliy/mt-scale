package route

import (
	"mt-scale/ctrls"
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

	// session_name := utils.Config.Section("session").Key("sessionname").String()
	// domain := utils.Config.Section("session").Key("sessiondomain").String()
	// maxage, _ := utils.Config.Section("session").Key("sessiongcmaxlifetime").Int()
	// csrfscret := utils.Config.Section("session").Key("csrfscret").String()
	//Session配置，Redis存储
	//store,err:=redis.NewStore(10,"tcp","rs1.rs.baidu.com:6379","",[]byte("asfajfa;lskjr2"))
	// store, err := redis.NewStoreWithPool(utils.RedisPool, []byte("as&8(0fajfa;lskjr2"))
	// store.Options(sessions.Options{
	// 	"/",
	// 	domain,
	// 	maxage,
	// 	false, //https 时使用
	// 	true,  //true:JS脚本无法获取cookie信息
	// })

	// if err != nil {
	// 	// Handle the error. Probably bail out if we can't connect.
	// 	fmt.Println("redis.NewStore error")
	// }

	// Router.Use(sessions.Sessions(session_name, store))

	// Router.Use(csrf.Middleware(csrf.Options{
	// 	Secret: csrfscret,
	// 	ErrorFunc: func(c *gin.Context) {
	// 		c.String(400, "CSRF token mismatch")
	// 		c.Abort()
	// 	},
	// }))

}

// SetupRouter Setup path
func SetupRouter() *gin.Engine {

	//静态目录配置
	public := utils.Config.Section("router").Key("public").String()
	Router.Static("/public", public)

	//模板
	viewPath := utils.Config.Section("router").Key("view_path").String()
	Router.LoadHTMLGlob(viewPath)

	//Session测试-Redis存储
	Router.GET("/session", ctrls.SessionAction)
	//Cookie测试
	Router.GET("/cookie", ctrls.CookieAction)
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
