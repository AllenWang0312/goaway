package api_restful

import (
	model "../model/meituri"
	"github.com/gin-gonic/gin"
)

func GetSplashInfo(c *gin.Context) {
	var splashs = []model.Splash{}
	//.Preload("Users").Preload("Models")
	db.Where("enable = ?", 1).Find(&splashs)
	if len(splashs) > 0 {
		c.JSON(200, gin.H{"data": splashs})
	} else {
		c.JSON(200, gin.H{"msg": "获取数据失败"})
	}

}
