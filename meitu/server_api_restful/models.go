package api_restful

import (
	"../../conf"
	model "../model/meituri"
	"../util"

	"github.com/gin-gonic/gin"
	"strconv"
)

func GetModelList(c *gin.Context) {
	//search := c.PostForm("search")
	//if tokenEnable(c) {
	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if nil == err1 && nil == err2 {
		var models = []model.Model{}
		//if len(search) == 0 {
		//} else {
		//}
		db.Table("models_cn").Order("hot desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models) //.Order("created_at desc")
		//c.String(200,)
		for i, m := range models {
			exit, _ := util.PathExists(conf.FSMuri + "/" + strconv.Itoa(m.ID))
			models[i].Get = exit
		}
		c.JSON(200, gin.H{"data": models})
	} else {
		c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
	}
	//}
}

type ModelDetail struct {
	Info   model.Model   `json:"info"`
	Colums []model.Colum `json:"colums"`
}

func GetModelHomePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	m := model.Model{
		ID: id,
	}
	if nil == err {
		db.First(&m)
	}
	colums := []model.Colum{}
	db.Where("modelid = ?", id).Find(&colums)

	c.JSON(200, gin.H{"data": ModelDetail{
		Info:   m,
		Colums: colums,
	},
	})

	//search := c.PostForm("search")
}
