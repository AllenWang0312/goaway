package api

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../../../conf"
)

var db *gorm.DB

func InitApiDB() {
	var err error
	//db, err = gorm.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/meitu?charset=utf8&parseTime=True&loc=Local") //?charset=utf8&parseTime=True&loc=Local
	db, err = gorm.Open("mysql", "root:"+conf.MysqlPass+"@tcp("+conf.DBHost+":3306)/meitu?charset=utf8&parseTime=True&loc=Local")
	//db, err = gorm.Open("mysql", "root:qunsi003@tcp(122.51.205.29:3306)/meitu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(conf.GormDebug)

}
