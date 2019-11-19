package meituri

type Group struct {
	Id       int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(20);index:name"`
	Belong   int    `gorm:"type:int(20);index:belong"`
	Homepage string `gorm:"type:varchar(255);index:homepage"`
}
