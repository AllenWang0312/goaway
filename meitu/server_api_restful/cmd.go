package api_restful

import (
	"../../conf"
	model "../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func ManageHot(c *gin.Context) {
	model_id, err0 := strconv.Atoi(c.Query(MODELID))
	hot, err1 := strconv.Atoi(c.Query(HOT))
	if nil == err0 && nil == err1 {
		var model = model.Model{}
		model.ID = model_id

		db.First(&model)
		model.Hot = hot
		db.Save(&model)
		c.JSON(200, gin.H{"toast": "编辑成功", "data": model.Hot})
	}
}

func CreateHistryForAlbum(c *gin.Context) {
	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))

	if err1 == nil && err2 == nil {
		var albums []model.Album
		db.Order("time desc").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&albums)
		var timeCache string
		var tablename string
		for _, a := range albums {
			var zone = model.Zone{}
			var time = strings.Trim(a.Time, " ")
			if time != timeCache {
				if len(time) > 0 {
					var chars = strings.Split(time, ".")
					if len(chars) == 3 {
						tablename = "zone" + chars[0] + "_" + chars[1]
					} else {
						tablename = "zone"
					}
					if !db.HasTable(tablename) {
						db.Table(tablename).CreateTable(model.Zone{})
					}
				}
				timeCache = time
			}
			zone.Albumid = a.ID
			zone.Modelid = a.Modelid
			zone.Groupid = a.Groupid
			zone.Time = time
			zone.Type = conf.Album

			//INSERT INTO `meitu`.`zone2019_08` (`id`, `type`, `userid`, `companyid`, `groupid`, `modelid`, `albumid`, `content`, `time`, `address`, `lat`, `long`) VALUES ('1', '4', '0', '0', '58', '3156', '30212', '', '2019.08.30', '', '0', '0');
			//db.Exec("INSERT INTO `meitu`.`" + tablename + "` (`type`, `groupid`, `modelid`, `albumid`, `time`)" +
			//	" VALUES ('"+conf.Album+"', '?', '?', '?', '?');",a.Groupid,a.Modelid,a.ID,time)
			new := db.Table(tablename).NewRecord(&zone)
			if (new) {
				db.Table(tablename).Create(&zone)
			} else {
				println("record exist")
			}
		}
		c.JSON(200, gin.H{"toast": "操作成功"})

	} else {
		c.JSON(200, gin.H{"toast": "参数有误"})
	}
}
