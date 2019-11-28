package api_restful

import (
	"../../conf"
	model "../model/meituri"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

//func GetTandomHotTab(c *gin.Context){
//
//}

//伪随机 15 个tag
func GetTandomHotTab(c *gin.Context) {
	gender := c.Query("gender")
	pageNo, err := strconv.Atoi(c.Query("pageNo"))

	if err == nil {
		var companies [] model.Company
		var groups [] model.Group
		var models []model.Model
		var tags []model.Tag
		var tabs [] model.Tab

		if gender == "" || gender == "man" {
			db.Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&companies)
			for _, c := range companies {
				c.Type = conf.Company
				tabs = append(tabs, c.Tab)
			}
			db.Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&groups)
			for _, g := range groups {
				g.Type = conf.Group
				tabs = append(tabs, g.Tab)
			}
			db.Table("models_cn").Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&models)
			for _, m := range models {
				m.Type = conf.Model
				tabs = append(tabs, m.Tab)
			}
			db.Select("id,name,hot").Order("hot desc").Offset((pageNo - 1) * 5).Limit(5).Find(&tags)
			for _, t := range tags {
				t.Type = conf.Tag
				tabs = append(tabs, t.Tab)
			}
		}
		c.JSON(200, gin.H{"data": tabs})
	}
}

func FollowTabs(c *gin.Context) {
	var user_id = getUserIdWithToken(c)
	if user_id > 0 {
		//var tabstr=c.PostForm("tabs")
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		tabstr := string(buf[0:n])

		var tabs [] model.Tab
		err := json.NewDecoder(strings.NewReader(string(tabstr))).Decode(&tabs)
		if err == nil {
			for _, t := range tabs {
				var follow = model.FollowTab{
					Userid:   user_id,
					Resid:    t.ID,
					Type:     t.Type,
					Alias:    t.Alias,
					Relation: strconv.Itoa(user_id) + "_" + strconv.Itoa(t.Type) + "_" + strconv.Itoa(t.ID),
				}
				var new = db.NewRecord(&follow)
				if new {
					db.Create(&follow)
				} else {
					db.Model(&follow).Where("id = ?", follow.ID).Update("alias", follow.Alias)
					//db.Model(&follow).UpdateColumn("alias", follow.Alias)
				}
			}
			c.JSON(200, gin.H{"toast": "关注成功"})
		}
	}
}
func FollowedTabs(c *gin.Context) {
	var userId = getUserIdWithToken(c)
	if userId > 0 {
		var followed []model.Tab
		var tabs []model.FollowTab
		db.Where("userid = ", userId).Find(&tabs)
		for _, t := range tabs {
			if t.Type == 0 {
				com := model.Company{}
				db.Where("id = ", t.Resid).First(&com)
				followed = append(followed, com.Tab)
			} else if t.Type == 1 {
				gro := model.Group{}
				db.Where("id = ", t.Resid).First(&gro)
				followed = append(followed, gro.Tab)
			} else if t.Type == 2 {
				mo := model.Model{}
				db.Where("id = ", t.Resid).First(&mo)
				followed = append(followed, mo.Tab)
			} else if t.Type == 3 {
				tag := model.Tag{}
				db.Where("id = ", t.Resid).First(&tag)
				followed = append(followed, tag.Tab)
			}
		}
		c.JSON(200, gin.H{"data": followed})
	} else {
		c.JSON(200, gin.H{"toast": "token 失效"})
	}

}

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
		pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
		pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
		if err1 == nil && err2 == nil {
			var tabs []model.FollowTab
			db.Where("userid = ? And type = 2", user_id).Find(&tabs)
			for _, f := range tabs {
				db = db.Or("modelid = ?", f.Resid)
			}
			var zones = []model.Zone{}
			db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&zones)
			c.JSON(200,gin.H{"data":zones})
		}else{
			c.JSON(200,gin.H{"toast":"参数错误"})
		}
	}else {

	}
}
