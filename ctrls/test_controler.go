package ctrls

import (
	"mt-scale/jobs"
	"mt-scale/queue"
	"mt-scale/utils"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

//Redis测试
func RedisSetAction(ctx *gin.Context) {

	rds := utils.RedisPool.Get()

	count, _ := redis.Int(rds.Do("GET", "count"))
	count++
	rds.Do("SET", "count", count)
	ctx.JSON(200, gin.H{
		"message": count,
	})

}

//Sesssion测试
func SessionAction(ctx *gin.Context) {

	session := sessions.Default(ctx)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count += 1
	}
	session.Set("count", count)
	session.Save()
	ctx.JSON(200, gin.H{"count": count})

}

//Cookie测试
func CookieAction(ctx *gin.Context) {

	var count int
	v, _ := ctx.Cookie("count")

	count, _ = strconv.Atoi(v)
	count++
	utils.SetCookie(ctx, "count", strconv.Itoa(count), 3600*24)

	ctx.JSON(200, gin.H{"count": v})

}

//队列生产者测试
func QueueAction(ctx *gin.Context) {
	queue.NewSender("SomeQueue", "Dosome", jobs.Subscribe{Name: "We are doing..."}).Send()
}
