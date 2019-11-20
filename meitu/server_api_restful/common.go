package api_restful

import "github.com/gin-gonic/gin"

func onError(c *gin.Context,code int){
	if code==401 {
		c.JSON(code,gin.H{"msg":"请先登录"})
	}
}