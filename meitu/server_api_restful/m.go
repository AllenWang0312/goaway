package api_restful

import (
	model "../model/meituri"
	"github.com/gin-gonic/gin"
)

func GetHomeData(c *gin.Context) {
	if (tokenEnable(c)) {
		var cn = c.Query("contry")
		var banners = [] model.Banner{}
		var companys = [] model.Company{}
		var models = []model.Model{}
		var colums = []model.Colum{}

		db.Where("state = ?", 1).First(&banners)
		db.Order("hot desc").Limit(10).Find(&companys)
		db.Table("models_" + cn).Order("hot desc").Limit(10).Find(&models)
		db.Order("hot desc").Limit(10).Find(&colums)
		var home = model.Home{
			Banners:banners,
			Companys:companys,
			Models:models,
			Colums:colums,
		}
		c.JSON(200,gin.H{"data":home})
	}
}
