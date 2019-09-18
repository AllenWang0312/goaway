package orm

import (
	model "../../../model/mm131"
	"database/sql"
	"fmt"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "Qunsi003"
	NETWORK  = "tcp"
	SERVER   = "rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com"
	PORT     = 3306
	DATABASE = "mm131"
)
var DB sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
}

func SaveColum(c *model.Colums,tableName string) {
	result,err := DB.Exec("insert INTO "+tableName+"(id,title,time,fenlei) values(?,?,?,?)",c.ID,c.Title,c.Time,c.Fenlei)
	if err != nil{
		fmt.Printf("Insert failed,err:%v",err)
		return
	}
	lastInsertID,err := result.LastInsertId()  //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v",err)
		return
	}
	fmt.Println("LastInsertID:",lastInsertID)
	rowsaffected,err := result.RowsAffected()  //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v",err)
		return
	}
	fmt.Println("RowsAffected:",rowsaffected)
}
