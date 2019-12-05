package gorm

import (
	"../download"
	"../../../../util"
	"../../../../conf"
	model "../../../model/meituri"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

var db *gorm.DB

func InitDB() {
	var err error
	//db, err = gorm.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/meitu?charset=utf8&parseTime=True&loc=Local") //?charset=utf8&parseTime=True&loc=Local
	db, err = gorm.Open("mysql", "root:"+conf.MysqlPass+"@tcp("+conf.DBHost+":3306)/meitu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func SaveColum(userId int, c *model.Album) {
	if err := db.Create(c).Error; err != nil {
		//return -3
		println(err.Error())
	}
}

func SaveGroupInfo(groups model.Group) int {
	if !db.HasTable("groups") {
		db.CreateTable(model.Group{})
	}
	if err := db.Create(groups).Error; err != nil {
		//return -3
		println(err.Error())
	} else {
		println("SaveGroupInfo,Success")
	}
	return -1
}
func SaveTag(tag model.Tag) int {
	//if !db.HasTable("tags") {
	//	db.CreateTable(model.Tag{})
	//}
	new := db.NewRecord(&tag)
	if (new) {
		db.Create(&tag)
	} else {
		db.Save(&tag)
	}
	return -1
}
func SaveTagInfo(id int, name string) {
	tag := model.Tag{}
	tag.ID = id
	tag.Name = name
	SaveTag(tag)
}
func DownloadCoverForModel(pageNo int, pageSize int) {
	var models []model.Model
	db.Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&models)
	for _, m := range models {
		var durl=m.Cover
		println(durl)

		filename:=util.GetNameFromUri(durl)
		path := conf.FSRoot + "/muri/cover/"
		//downloadFile(durl,path,filename)
		download.WG.Add(1)
		download.DownloadImage(durl, path, filename)
	}
}
func CreateHistryForAlbum(pageNo int, pageSize int) {
	var albums []model.Album
	db.Order("time desc").Offset((pageNo - 1) * pageSize).Limit(pageSize).Find(&albums)
	var timeCache string
	var tablename string
	for _, a := range albums {
		var zone = model.Zone{}
		var time = strings.Trim(a.Time, " ")
		tablename = "zones"
		if time != timeCache {
			if len(time) > 0 {
				var chars = strings.Split(time, ".")
				if len(chars) == 3 {
					tablename = "zones" + chars[0] + "_" + chars[1]
				}
			}
			timeCache = time
		}
		if !db.HasTable(tablename) {
			db.Table(tablename).CreateTable(model.Zone{})
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
}
func UpDateHot(id int, hot int) {
	var model = model.Model{}
	model.ID = id
	db.First(&model)
	model.Hot = hot
	db.Save(&model)
}
func CreateSplashForColum(modelid int, albumid int, src string) {
	var album = model.Album{
		ID:      albumid,
		Modelid: modelid,
	}
	db.First(&album)

	if len(album.Title) > 0 {
		var splash = model.Splash{}
		splash.Type = 0
		splash.Src = conf.FILE_SERVER + "/muri/" + strconv.Itoa(modelid) + "/" + strconv.Itoa(albumid) + "/" + src
		splash.Link = "/album?id=" + strconv.Itoa(albumid)
		db.Table("banners").Create(&splash)
	}
}

func CreateTableForModels(offset int, count int) {
	var models = [] model.Model{}
	db.Offset(offset).Limit(count).Find(&models)
	for i, m := range models {
		var end = "cn"
		if strings.Contains(m.Address, "日本") {
			end = "jp"
		} else if strings.Contains(m.Address, "美国") {
			end = "usa"
		} else if strings.Contains(m.Address, "韩国") {
			end = "kr"
		} else if strings.Contains(m.Address, "泰国") {
			end = "tha"
		} else if strings.Contains(m.Address, "香港") {
			end = "cn_hk"
		} else if strings.Contains(m.Address, "台湾") {
			end = "cn_tw"
		} else if strings.Contains(m.Address, "澳门") {
			end = "cn_mo"
		}
		var tableName = "models_" + end
		if !db.HasTable(tableName) {
			db.Table(tableName).CreateTable(model.Model{})
		}
		//new := db.Table(tableName).NewRecord(&m)
		//if (new) {
		db.Table(tableName).Create(&m)
		//} else {
		//	db.Table(tableName).Save(&m)
		//}
		println(i, m.ID, m.Name+m.Address)
	}
}

func GetCNModels() *[]model.Model {
	models := []model.Model{}
	//.Where("id >= 917")
	db.Table("models_cn").Select("id,hot").Order("hot desc").Find(&models)
	return &models
}

//func addColumToFavourite(token string, columId int) {
//
//}
//func cancelFavourite(token string, columId int) {
//
//}

//个人介绍页 获取资料
func SaveModelInfo(contry string, m *model.Model) {

	if len(contry) > 0 {
		var tableName = "models_" + contry
		if !db.HasTable(tableName) {
			db.CreateTable(model.Model{})
		}
		db.Table(tableName).Create(m)
	} else {
		if !db.HasTable("models") {
			db.CreateTable(model.Model{})
		}
		db.Create(m)
	}
	//createSuccess := db.NewRecord(m)
	//if createSuccess {
	//	fmt.Println("createSuccess")
	//
	//}
}
func SaveColumInfo(columId int, c *model.Album) int {
	tags := c.Tags
	tag := strings.Split(tags, ")")
	for _, str := range tag {
		vk := strings.Split(str, "(")
		if len(vk) >= 2 {
			id, _ := strconv.Atoi(vk[1])
			SaveTagInfo(id, vk[0])
		}
	}
	if !db.HasTable("albums") {
		db.CreateTable(model.Album{})
	}
	if err := db.Create(c).Error; err != nil {
		//return -3
		println(err.Error())
	} else {
		println("SaveColumInfo,Success" + strconv.Itoa(columId))
	}
	return -1
}
