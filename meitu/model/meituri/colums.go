package meituri

type Colums struct {
	ID      int `gorm:"primary_key;index:id" json:"id"`
	Modelid int `gorm:"index:modelid" json:"modelid"`

	Title   string `gorm:"type:varchar(100);index:title" json:"title"`
	Tags    string `gorm:"type:varchar(100);index:tags" json:"tags"`
	Subs    string `gorm:"type:varchar(500);index:subs" json:"subs"`
	Group   string `gorm:"type:varchar(100);index:group" json:"org"`
	Groupid int    `gorm:"index:groupid" json:"orgid"`
	No      string `gorm:"type:varchar(20);index:no" json:"no"`
	Nums    int    `gorm:"index:nums" json:"nums"`
	Time    string `gorm:"type:varchar(30);index:time" json:"time"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
}
