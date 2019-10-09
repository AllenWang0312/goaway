package api_restful

import (
	model "../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetColumsList(c *gin.Context) {
	tag, err0 := strconv.Atoi(c.Query("tag"))

	pageNo, err1 := strconv.Atoi(c.PostForm("pageNo"))
	pageSize, err2 := strconv.Atoi(c.PostForm("pageSize"))
	var colums = []model.Colums{}

	if err0 == nil {
		db.Where("tags LIKE ?", tag).Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&colums) //.Order("created_at desc")
		//c.String(200,)
		c.JSON(200, gin.H{"data": colums})
	} else {
		if nil == err1 && nil == err2 {
			//if len(search) == 0 {
			//} else {
			//}
			db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Order("id desc").Find(&colums) //.Order("created_at desc")
			//c.String(200,)
			c.JSON(200, gin.H{"data": colums})
		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}
}
