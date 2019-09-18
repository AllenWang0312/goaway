package main

import (
	"../../api_restful"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	api_restful.InitApiDB()
	// 默认启动方式，包含 Logger、Recovery 中间件
	r := gin.Default()
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
	v1:=r.Group("/v1")
	api := v1.Group("/api")
	{
		//todo resources
		api.POST("/models", api_restful.GetModelList)
		api.POST("/colums", api_restful.GetColumsList)
		//todo account
		account:=api.Group("/account")
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
