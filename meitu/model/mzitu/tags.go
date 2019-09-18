package mzitu

type Tags struct {
	ID    int    `gorm:"primary_key;AUTO_INCREMENT;index:id" json:"id"`
	Ename string `gorm:"size:20;index:ename" json:"ename"`
	Cname string `gorm:"size:20;index:cname" json:"cname"`
}
