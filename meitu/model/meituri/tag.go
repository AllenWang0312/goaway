package meituri

type Tag struct {
	Id        int    `gorm:"primary_key;index:id" json:"id,omitempty"`
	Name      string `gorm:"type:varchar(100);index:name" json:"name,omitempty"`
	Shortname string `gorm:"type:varchar(8);index:shortname" json:"shortname,omitempty"`
	Des       string `gorm:"type:varchar(200);index:des" json:"des,omitempty"`
	Nums      int    `gorm:"index:nums" json:"nums,omitempty"`
	Hot       int    `gorm:"index:hot" json:"hot,omitempty"`
}
