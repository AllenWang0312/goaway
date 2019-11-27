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
			db.Select("id,name,hot").Order("hot desc").Offset(pageNo * 5).Limit(5).Find(&companies)
			for _, c := range companies {
				c.Type=conf.Company
				tabs = append(tabs, c.Tab)
			}
			db.Select("id,name,hot").Order("hot desc").Offset(pageNo * 5).Limit(5).Find(&groups)
			for _, g := range groups {
				g.Type=conf.Group
				tabs = append(tabs, g.Tab)
			}
			db.Table("models_cn").Select("id,name,hot").Order("hot desc").Offset(pageNo * 5).Limit(5).Find(&models)
			for _, m := range models {
				m.Type=conf.Model
				tabs = append(tabs, m.Tab)
			}
			db.Select("id,name,hot").Order("hot desc").Offset(pageNo * 5).Limit(5).Find(&tags)
			for _, t := range tags {
				t.Type=conf.Tag
				tabs = append(tabs, t.Tab)
			}
		}
		c.JSON(200, gin.H{"data":tabs})
	}
}

func FollowTabs(c *gin.Context){

	var user_id=getUserIdWithToken(c)
	var tabstr=c.PostForm("tabs")

	var tabs [] model.Tab
	err := json.NewDecoder(strings.NewReader(string(tabstr))).Decode(&tabs)
	if err==nil {
		for i,t:= range tabs{
			println(i,t.Name)
			var follow=model.FollowTab{
				UserId:user_id,
				ResId:t.ID,
				Type:t.Type,
				Relation:strconv.Itoa(user_id)+"_"+strconv.Itoa(t.Type)+"_"+strconv.Itoa(t.ID),
			}
			var new = db.NewRecord(&follow)
			if(new){
				db.Save(&follow)
			}
		}
		c.JSON(200,gin.H{"toast":"关注成功"})
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
