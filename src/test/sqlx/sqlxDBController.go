package sqlx

import (
	"../../meitu/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

func InitDB() {
	dbx, _ = sqlx.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/meitu?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
}

func SaveColumRelation(userId int, colum int) {

}
func SaveModelInfo(userId int, model *model.Models) {
	//tx := dbx.MustBegin()
	//tx.MustExec("INSERT INTO models (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
}

func SaveColumInfo(colum int, userId int, doc *goquery.Document) int {

	return 0
}