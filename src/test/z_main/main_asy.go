package z_main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

//const (
//	SaveUserInfo  = true
//	SaveColumInfo = true
//)
type DownloadTask
{
durl:string
path:string
name:string
}

//var db *gorm.DB
//var client *http.Client
//var tasks []DownloadTask
//
//fun (task *DownloadTask)download(){
//
//
//}
func main() {
	var err error
	db, err = gorm.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/meitu?charset=utf8&parseTime=True&loc=Local") //?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	for {
		if tasks != nil && len(task) == 0 {
			return
		}

	}
