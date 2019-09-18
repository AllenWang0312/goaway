package mzitu

type Columtags struct {
	Relationid int `gorm:"primary_key;index:relationid" json:"relation_id"`
	Columid    int `gorm:"index:columid" json:"colum_id"`
	Tagid      int `gorm:"index:tagid" json:"colum_id"`
	Lock string `gorm:"size:30,index:lock" json:"lock"`
}
