package sql

import (
	"fmt"
	"database/sql"
	"time"
)
const (
	USERNAME = "root"
	PASSWORD = "qunsi003"
	NETWORK = "tcp"
	SERVER = "localhost"
	PORT = 3306
	DATABASE = "meitu"
)
func main(){
	dsn:=fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB,err:=sql.Open("mysql",dsn)
	if err!=nil{
		fmt.Printf("Open mysql failed,err:%v\n",err)
		return
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数

}
func insertModelData(DB *sql.DB) {
	result, err := DB.Exec("insert INTO models(name,age) values(?,?)", "YDZ", 23)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return
	}
	fmt.Println("LastInsertID:", lastInsertID)
	//rowsaffected,err := result.RowsAffected()  //影响行数
	//if err != nil {
	//	fmt.Printf("Get RowsAffected failed,err:%v",err)
	//	return
	//}
	//fmt.Println("RowsAffected:",rowsaffected)
}