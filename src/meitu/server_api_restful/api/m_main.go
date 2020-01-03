package api

import (
	"../../../conf"
	model "../../../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func GetHomeData(c *gin.Context) {
	if tokenEnable(c) {
		var cn = c.Query("contry")
		//var t = c.Query("type")

		var banners [] model.Banner
		var companies [] model.Company

		var models []model.Model

		var apps []model.App

		db.Where("type = ?", 2).Find(&banners)
		db.Order("hot desc").Limit(10).Find(&companies)
		var modeltable = "models"
		if len(cn) > 0 {
			modeltable = modeltable + "_" + cn
		}
		db.Table(modeltable).Order("hot desc").Limit(10).Find(&models)
		db.Find(&apps)

		var home = model.Home{
			Banners:  banners,
			Apps:     apps,
			Companys: companies,
			Models:   models,
		}
		c.JSON(200, gin.H{"data": home})
	}
}
func GetZoneHistroy(c *gin.Context) {
	var user_id = getUserIdWithToken(c)
	if user_id > 0 {
		//scope := c.Query("scope")
		year := c.Query("year")
		month := c.Query("month")
		pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
		pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
		if err1 == nil && err2 == nil {
			//if strings.EqualFold(scope, "follow") {
			//	var tabs []model.FollowTab
			//	db.Where("userid = ? And type = 2", user_id).Find(&tabs)
			//	for _, f := range tabs {
			//		db = db.Or("modelid = ?", f.Resid)
			//	}
			//}
			var zones []model.Zone
			if (strings.EqualFold(month, "00")) {
				var tablename = "zones"
				if !db.HasTable(tablename) {
					//tablename = "zone"
					c.JSON(200, gin.H{"toast": "没有这个月份的数据哦"})
					return
				}
				db.Table(tablename).Where("time like ?", year).Offset((pageNo - 1) * pageSize).Limit(pageSize).Preload("Model").Preload("Album").Find(&zones)
			} else {
				var tablename = "zones" + year + "_" + month
				if !db.HasTable(tablename) {
					//tablename = "zone"
					c.JSON(200, gin.H{"toast": "没有这个月份的数据哦"})
					return
				}
				db.Table(tablename).Offset((pageNo - 1) * pageSize).Limit(pageSize).Preload("Model").Preload("Album").Find(&zones)
			}

			c.JSON(200, gin.H{"data": zones})
		} else {
			c.JSON(200, gin.H{"code": -1, "toast": "参数错误"})
		}
	} else {

	}
}
func RecordVisitHistroy(c *gin.Context){
	useridstr := getUserIdStrWithToken(c)
	userid, err := strconv.Atoi(useridstr)
	if err == nil && userid > 0 {
		//modelIdStr := c.Query("model_id")
		albumIdStr := c.Query("album_id")
		albumId, _ := strconv.Atoi(albumIdStr)
		album := model.Album{}
		db.Where("id = ?", albumIdStr).Preload("Model").First(&album)
		if album.ID > 0 {
			var now = time.Now().Format("2006-01-02")
			var record = model.VisitHistroy{
				Albumid:  albumId,
				Userid:   userid,
				Date:     now,
				Relation: albumIdStr + "_" + useridstr + "_" + now,
			}
			var tableName = conf.VisitHistroy + strconv.Itoa(userid/1000)
			if !db.HasTable(tableName) {
				db.Table(tableName).Create(model.VisitHistroy{})
			}
			db.Table(tableName).Create(&record)
			c.JSON(200,gin.H{"msg":"record success"})
		}else{
			c.JSON(200,gin.H{"toast":"album id 不存在"})
		}
	}else {
		c.JSON(200,gin.H{"toast":"获取用户信息失败"})
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
				db.Table(tableName).Preload("Album").Where("userid = ?", userId).Order("date desc").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&records)
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
