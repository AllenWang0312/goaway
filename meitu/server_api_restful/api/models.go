package api

import (
	"../../../conf"
	model "../../model/meituri"
	"../../../util"
	"strings"

	"github.com/gin-gonic/gin"
	"strconv"
)

func GetModelList(c *gin.Context) {
	//search := c.PostForm("search")
	//if tokenEnable(c) {
	contry:=c.Query("contry")
	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))

	if nil == err1 && nil == err2 {
		var models []model.Model
		//if len(search) == 0 {
		//} else {
		//}
		if len(contry)>0 {
			if strings.EqualFold(contry,"cn") {
				db.Table("models_cn").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"jp") {
				db.Where("address like ?", "日本").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"cn_hk"){
				db.Where("address like ?", "香港").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"cn_mo"){
				db.Where("address like ?", "澳门").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"cn_tw"){
				db.Where("address like ?", "台湾").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"kr"){
				db.Where("address like ?", "韩国").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"tha"){
				db.Where("address like ?", "泰国").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}else if strings.EqualFold(contry,"usa"){
				db.Where("address like ?", "美国").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
			}

		}else{
			db.Table("models").Order("hot desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&models)
		}
	 //.Order("created_at desc")
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
	Albums []model.Album `json:"albums"`
}

func GetModelHomePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	m := model.Model{}
	m.ID = id
	if nil == err {
		db.First(&m)
	}
	var albums []model.Album
	db.Where("modelid = ?", id).Find(&albums)
	c.JSON(200, gin.H{"data": ModelDetail{
		Info:   m,
		Albums: albums,
	},
	})
	//search := c.PostForm("search")
}
