package gorm

import (
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
	db, err = gorm.Open("mysql", "root:qunsi003@tcp("+conf.DBHost+":3306)/meitu?charset=utf8&parseTime=True&loc=Local")
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

func CreateTableForModels(end string, str string) {
	var models = [] model.Model{}
	db.Where("address like ?", "%"+str+"%").Find(&models)
	var tableName = "models_" + end
	if !db.HasTable(tableName) {
		db.Table(tableName).CreateTable(model.Model{})
	}
	for i, m := range models {
		new := db.Table(tableName).NewRecord(&m)
		if (new) {
			db.Table(tableName).Create(&m)
		} else {
			db.Table(tableName).Save(&m)
		}
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
