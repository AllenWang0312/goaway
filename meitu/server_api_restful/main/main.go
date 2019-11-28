package main

import (
	"../../cache"
	"../../server_api_restful"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

//var Pool redis.Pool

//func init() {
//	Pool = redis.Pool{
//		MaxIdle:     16,
//		MaxActive:   32,
//		IdleTimeout: 120,
//		Dial: func() (redis.Conn, error) {
//			return redis.Dial("tcp", "192.168.0.100:6379")
//		},
//	}
//}

func main() {

	cache.InitConn()

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

	m:=api.Group("/m")
	{
		tabs:=m.Group("/tabs")
		{
			tabs.GET("/",api_restful.GetTandomHotTab)
			tabs.POST("/follow",api_restful.FollowTabs)
			tabs.GET("/followed",api_restful.FollowedTabs)

		}
		home:=m.Group("/home")
		{
			home.GET("/",api_restful.GetHomeData)
			home.GET("/zone",api_restful.GetZoneHistroy)
		}
	}

	{
		model := api.Group("/model")
		{ //todo resources
			model.GET("/list", api_restful.GetModelList)
			model.GET("", api_restful.GetModelHomePage)
		}
		colum := api.Group("/colum")
		{
			colum.GET("/list", api_restful.GetColumsList)
			colum.GET("", api_restful.GetColumDetail)
		}

		tag := api.Group("/tag")
		{
			tag.GET("/hot", api_restful.GetHotTag)
		}
		//group := api.Group("/group")
		//{
		//	group.GET("/hot", api_restful.GetHotTag)
		//}

		config := api.Group("/config")
		{
			config.GET("/splash", api_restful.GetSplashInfo)
		}
		like := api.Group("/like")
		{
			like.GET("", api_restful.Like)
			like.POST("/models", api_restful.LikeModelList)
			like.POST("/colums", api_restful.LikeColumList)
		}
		cmd := api.Group("/cmd")
		{
			cmd.GET("/hot", api_restful.ManageHot)
			cmd.GET("/zone",api_restful.CreateHistryForAlbum)
		}


		feedback := api.Group("/feedback")
		{
			feedback.GET("/list", api_restful.GetCommits)
			feedback.POST("/commit", api_restful.Commit)
			feedback.POST("/like", api_restful.LikeCommit)
		}
		//todo account
		account := api.Group("/account")
		{
			account.POST("/register", api_restful.Regist)
			//account.PUT("/", api_restful.RegistAccount)

			account.POST("/info", api_restful.GetUser)
			account.POST("/user/info", api_restful.GetUser)
			//account.POST("/user",api)
			account.POST("/tokenlogin", api_restful.TokenLogin)
			account.POST("/login", api_restful.Login)
		}
		api.POST("/file", api_restful.UploadFile)
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
