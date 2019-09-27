package gorm

import (
	model "../../../model/doubanmovie"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/doubandianyin?charset=utf8&parseTime=True&loc=Local") //?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func SaveMovie( c *model.Subject) {

	if err := db.Create(c).Error; err != nil {
		//return -3
		println(err.Error())
	}
}
//func updateTag(tagId int, shortname string) {
//	tag := model.Tags{
//		Id:        tagId,
//		Shortname: shortname,
//	}
//	db.Model(tag).Update("shortname", shortname)
//}
//func SaveGroupInfo(groups model.Groups) int {
//	if err := db.Create(groups).Error; err != nil {
//		//return -3
//		println(err.Error())
//	} else {
//		println("SaveGroupInfo,Success")
//	}
//	return -1
//}
//func SaveTagInfo(tag model.Tags) int {
//	if err := db.Create(tag).Error; err != nil {
//		println(err.Error())
//	} else {
//		println("SaveTagInfo,Success")
//	}
//	return -1
//}
//
////个人介绍页 获取资料
//func SaveModelInfo(m *model.Models) {
//	createSuccess := db.NewRecord(m)
//	if createSuccess {
//		fmt.Println("createSuccess")
//		err1 := db.Create(m).Error
//		if err1 != nil {
//			println(err1.Error())
//		} else {
//
//		}
//	}
//}
//func addColumToFavourite(token string, columId int) {
//
//}
//func cancelFavourite(token string, columId int) {
//
//}
//func SaveColumInfo(columId int, c *model.Colums) int {
//	tags := c.Tags
//	tag := strings.Split(tags, ")")
//	for _, str := range tag {
//		vk := strings.Split(str, "(")
//		if len(vk) >= 2 {
//			id, _ := strconv.Atoi(vk[1])
//			updateTag(id, vk[0])
//		}
//	}
//
//	if err := db.Create(c).Error; err != nil {
//		//return -3
//		println(err.Error())
//	} else {
//		println("SaveColumInfo,Success" + strconv.Itoa(columId))
//	}
//	return -1
//}
