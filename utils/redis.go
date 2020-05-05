package utils

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisPool   *redis.Pool
	maxIdle     int
	maxActive   int
	host        string
	database    int
	idleTimeout int
	timeout     int
	password    string
)

func init() {
	maxIdle = GetConfInt("maxIdle")
	maxActive = GetConfInt("maxActive")
	host = GetConfStr("host")
	database = GetConfInt("database")
	idleTimeout = GetConfInt("idleTimeout")
	timeout = GetConfInt("timeout")
	password = GetConfStr("password")
}

// CreateRedisPool Create redis
func CreateRedisPool() *redis.Pool {
	if redisPool == nil {
		idleTimeDuration := time.Duration(idleTimeout) * time.Second
		timeoutDuration := time.Duration(timeout)

		redisPool = &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			IdleTimeout: idleTimeDuration,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				con, err := redis.Dial("tcp", host,
					redis.DialPassword(password),
					redis.DialDatabase(database),
					redis.DialConnectTimeout(timeoutDuration),
					redis.DialReadTimeout(timeoutDuration),
					redis.DialWriteTimeout(timeoutDuration))
				if err != nil {
					return nil, err
				}
				return con, nil
			},
		}
	}
	return redisPool
}
