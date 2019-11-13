package api_restful

import (
	model "../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ManageHot(c *gin.Context) {
	model_id, err0 := strconv.Atoi(c.Query(MODELID))
	hot, err1 := strconv.Atoi(c.Query(HOT))
	if nil == err0 && nil == err1 {
		var model = model.Models{
			ID: model_id,
		}
		db.First(&model)
		model.Hot = hot
		db.Save(&model)
		c.JSON(200, gin.H{"toast": "编辑成功", "data": model.Hot})
	}
}
