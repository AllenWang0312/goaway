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
		var model = model.Model{}
		model.ID=model_id

		db.First(&model)
		model.Hot = hot
		db.Save(&model)
		c.JSON(200, gin.H{"toast": "编辑成功", "data": model.Hot})
	}
}

func createHistryForAlbum(c *gin.Context){
	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))

	var albums=[]model.Album{}
	db.Select("id,").Order("time desc").Offset((pageNo-1)*pageSize).Limit(pageSize).Find(&albums)
db.Create(model.)
}
