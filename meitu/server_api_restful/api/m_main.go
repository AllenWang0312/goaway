package api

import (
	"../../../conf"
	model "../../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetHomeData(c *gin.Context) {
	if tokenEnable(c) {
		var cn = c.Query("contry")
		var banners [] model.Banner
		var companies [] model.Company
		var models []model.Model

		db.Where("state = ?", 1).Find(&banners)
		db.Order("hot desc").Limit(10).Find(&companies)
		db.Table("models_" + cn).Order("hot desc").Limit(10).Find(&models)

		var home = model.Home{
			Banners:  banners,
			Companys: companies,
			Models:   models,
		}
		c.JSON(200, gin.H{"data": home})
	}
}
func GetZoneHistroy(c *gin.Context) {
	var user_id = getUserIdWithToken(c)
	if user_id > 0 {
		year := c.Query("year")
		month := c.Query("month")
		pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
		pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
		if err1 == nil && err2 == nil {
			var tabs []model.FollowTab
			db.Where("userid = ? And type = 2", user_id).Find(&tabs)
			for _, f := range tabs {
				db = db.Or("modelid = ?", f.Resid)
			}
			var zones []model.Zone
			var tablename = "zones" + year + "_" + month
			if !db.HasTable(tablename) {
				//tablename = "zone"
				c.JSON(200, gin.H{"toast": "没有这个月份的数据哦"})
				return
			}
			db.Table(tablename).Offset((pageNo - 1) * pageSize).Limit(pageSize).Preload("Model").Preload("Album").Find(&zones)

			c.JSON(200, gin.H{"data": zones})
		} else {
			c.JSON(200, gin.H{"code": -1, "toast": "参数错误"})
		}
	} else {

	}
}
func GetVisitHistroy(c *gin.Context) {
	var pageNo, err0 = strconv.Atoi(c.Query("pageNo"))
	var pageSize, err1 = strconv.Atoi(c.Query("pageSize"))
	if err0 == nil && err1 == nil {
		userIdStr := getUserIdStrWithToken(c)
		userId, err := strconv.Atoi(userIdStr)
		if err == nil {
			var records []model.VisitHistroy
			tableName := conf.VisitHistroy + strconv.Itoa(userId/1000)
			if db.HasTable(tableName) {
				db.Table(tableName).Where("userid = ?", userId).Order("date desc").Offset((pageNo - 1) * pageSize).Limit(pageSize).Preload("Album").Find(&records)
				c.JSON(200, gin.H{"data": records})
			} else {
				c.JSON(200, gin.H{"toast": "没有记录"})
			}

		} else {
			c.JSON(200, gin.H{"msg": err.Error()})
		}
	} else {
		c.JSON(200, gin.H{"toast": "参数错误"})
	}
}
func CleanVisitHistroy(c *gin.Context) {
	var userid = getUserIdWithToken(c)
	if userid > 0 {
		var records [] model.VisitHistroy
		db.Where("userid = ?", userid).Find(&records)
		for _, r := range records {
			db.Delete(r)
		}
		c.JSON(200, gin.H{"toast": "操作成功"})
	} else {
		c.JSON(200, gin.H{"toast": "获取用户信息失败"})
	}
}
