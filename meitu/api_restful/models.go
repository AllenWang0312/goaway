package api_restful

import (
	model "../model/meituri"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func likeModel(c *gin.Context) {
	userId, err1 := strconv.Atoi(c.PostForm("userId"))
	modelId, err2 := strconv.Atoi(c.PostForm("modelId"))
	if nil == err1 && nil == err2 {
		like := model.Like{
			Userid:   userId,
			Modelid:  modelId,
			Relation: strconv.Itoa(userId) + "_" + strconv.Itoa(modelId),
		}
		db.Create(&like)
		model := model.Models{
			ID: modelId,
		}
		db.Find(&model)
		model.Hot = model.Hot + 100
		fmt.Print(model.Name, model.Hot)
		db.Update(&model)
		c.JSON(200, gin.H{"message": "收藏成功"})
	} else {
		c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
	}
}
func GetModelList(c *gin.Context) {
	//search := c.PostForm("search")
	if tokenEnable(c) {
		pageNo, err1 := strconv.Atoi(c.PostForm("pageNo"))
		pageSize, err2 := strconv.Atoi(c.PostForm("pageSize"))
		if nil == err1 && nil == err2 {
			var models = []model.Models{}
			//if len(search) == 0 {
			//} else {
			//}
			db.Table("models_cn").Order("hot desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models) //.Order("created_at desc")
			//c.String(200,)
			c.JSON(200, gin.H{"data": models})
		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}
}

type ModelDetail struct {
	Info   model.Models
	Colums []model.Colums
}

func GetModelHomePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	m := model.Models{
		ID: id,
	}
	if nil == err {
		db.First(&m)
	}
	colums := []model.Colums{}
	db.Where("modelid = ?", id).Find(&colums)

	c.JSON(200, gin.H{"data": ModelDetail{
		Info:   m,
		Colums: colums,
	},
	})

	//search := c.PostForm("search")
}
