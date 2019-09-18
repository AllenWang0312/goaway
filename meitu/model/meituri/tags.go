package meituri

type Tags struct {
	Id        int    `gorm:"primary_key;index:id" json:"id"`
	Name      string `gorm:"type:varchar(100);index:name" json:"name"`
	Shortname string `gorm:"type:varchar(8);index:shortname" json:"shortname"`
	Des       string `gorm:"type:varchar(200);index:des" json:"des"`
	Nums      int    `gorm:"index:nums" json:"nums"`
}
