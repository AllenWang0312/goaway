package gorm

import (
	model "../../../model/mzitu"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/mzitu?charset=utf8&parseTime=True&loc=Local") //?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func SaveColum(c *model.Colums) {
	if err := db.Create(c).Error; err != nil {
		//return -3
		println(err.Error())
	}
}
func SaveTagRelation(r *model.Columtags)int{
	var tag model.Columtags
	db.Where("lock = ?", r.Lock).First(&tag)
	if(tag.Relationid>0){
		println("record exit:"+strconv.Itoa(tag.Relationid))
		return tag.Relationid
	}else {
		db.Create(r)
		return r.Relationid
	}
}
func SaveTags(ename string,c *model.Tags) int {
	var tag model.Tags
	db.Where("ename = ?", ename).First(&tag)
	if(tag.ID>0){
		println("record exit:"+strconv.Itoa(tag.ID)+tag.Cname+tag.Ename)
		return tag.ID
	}else {
		db.Create(c)
		return c.ID
	}
}
