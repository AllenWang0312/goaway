package gorm

import (
	"../../../../conf"
	"fmt"
	"github.com/jinzhu/gorm"
	model "../../../model/fengniao"
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
	db.LogMode(conf.GormDebug)
}
func SaveAlbum(album *model.Album){
	new :=db.Table("fengniao_album").NewRecord(album)
	if new {
		db.Table("fengniao_album").Create(album)
	}else{
		db.Table("fengniao_album").Save(album)
	}

}
