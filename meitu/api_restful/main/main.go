package main

import (
	"../../api_restful"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var Pool redis.Pool

func init() {
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.0.100:6379")
		},
	}
}

func main() {

	//conn := Pool.Get()
	//res, err := conn.Do("HSET", "student", "name", "jack")
	//res1, err := redis.String(conn.Do("HGET", "student", "name"))
	//
	//ok, rdsKey := redis.PutJSON("", "user", rdsVal, 1800*time.Second)

	api_restful.InitApiDB()
	r := gin.Default()
	// 默认启动方式，包含 Logger、Recovery 中间件
	//无中间件启动
	//r := gin.New()
	//r.Use(gin.Logger())
	//自定义日志格式
	//r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)

	//}))
	//r.Use(gin.Recovery())

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 1, "message": "hello"})
	})
	v1 := r.Group("/v1")
	api := v1.Group("/api")
	{

		model := api.Group("/model")
		{ //todo resources
			model.POST("/list", api_restful.GetModelList)
			model.GET("", api_restful.GetModelHomePage)
		}
		api.POST("/colums", api_restful.GetColumsList)
		api.GET("/colum", api_restful.GetColumPhotos)

		tags := api.Group("/tags")
		{
			tags.GET("/hot", api_restful.GetHotTag)
		}

		//todo account
		account := api.Group("/account")
		{
			account.PUT("/", api_restful.RegistAccount)
			account.POST("/login", api_restful.Login)
		}
		//v1.POST("/login", loginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}
	//v2 := r.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}
	//initLogger()
	r.Run()
}

func initLogger() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
