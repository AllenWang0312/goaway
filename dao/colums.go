package dao

type Colums struct {
	ID     int    `gorm:"primary_key"`
	Userid int    `gorm:"index:userid"`
	Title  string `gorm:"type:varchar(100);index:title"`
	Subs   string `gorm:"type:varchar(100);index:title"`
	Org    string `gorm:"type:varchar(100);index:title"`
	Orgid  int    `gorm:"index:title"`
	No     int    `gorm:"index:title"`
	Nums   int    `gorm:"index:title"`
	Time   string `gorm:"type:varchar(30);index:title"`
}
