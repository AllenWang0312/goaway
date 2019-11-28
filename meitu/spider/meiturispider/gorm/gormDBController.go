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

func updateTag(tagId int, shortname string) {
	tag := model.Tag{}
	tag.ID=tagId
	tag.Name=shortname
	db.Model(tag).Update("name", shortname)
}

func SaveGroupInfo(groups model.Group) int {
	if err := db.Create(groups).Error; err != nil {
		//return -3
		println(err.Error())
	} else {
		println("SaveGroupInfo,Success")
	}
	return -1
}
func SaveTagInfo(tag model.Tag) int {
	if err := db.Create(tag).Error; err != nil {
		println(err.Error())
	} else {
		println("SaveTagInfo,Success")
	}
	return -1
}
func CreateTableForModels(str string) {
	var models = [] model.Model{}
	db.Where("address like ?", "%"+str+"%").Find(&models)
	for i, m := range models {
		//new :=db.Table("models_jp").NewRecord(&m)
		//if(new){
		db.Table("models_jp").Save(&m)
		//}
		println(i, m.ID, m.Name+m.Address)
	}
}

//个人介绍页 获取资料
func SaveModelInfo(m *model.Model) {
	err1 := db.Create(m).Error
	if err1 != nil {
		fmt.Println(err1.Error())
	} else {

	}
	//createSuccess := db.NewRecord(m)
	//if createSuccess {
	//	fmt.Println("createSuccess")
	//
	//}
}
func GetCNModels() *[]model.Model {
	models := []model.Model{}
	//.Where("id >= 917")
	db.Table("models_cn").Select("id").Find(&models)
	return &models
}
func addColumToFavourite(token string, columId int) {

}
func cancelFavourite(token string, columId int) {

}
func SaveColumInfo(columId int, c *model.Album) int {
	tags := c.Tags
	tag := strings.Split(tags, ")")
	for _, str := range tag {
		vk := strings.Split(str, "(")
		if len(vk) >= 2 {
			id, _ := strconv.Atoi(vk[1])
			updateTag(id, vk[0])
		}
	}
	if err := db.Create(c).Error; err != nil {
		//return -3
		println(err.Error())
	} else {
		println("SaveColumInfo,Success" + strconv.Itoa(columId))
	}
	return -1
}
