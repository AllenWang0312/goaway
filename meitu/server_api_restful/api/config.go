package api

import (
	model "../../model/meituri"
	"github.com/gin-gonic/gin"
)

func GetSplashInfo(c *gin.Context) {
	var splashs = []model.Banner{}
	//.Preload("Users").Preload("Models")
	db.Where("type = ?", 0).Find(&splashs)
	if len(splashs) > 0 {
		c.JSON(200, gin.H{"data": splashs})
	} else {
		c.JSON(200, gin.H{"msg": "获取数据失败"})
	}

}
