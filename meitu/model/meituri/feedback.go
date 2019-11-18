package meituri

type Feedback struct {
	Id         int    `gorm:"primary_key" json:"id"`
	Userid     int    `gorm:"type:int(20);index:userid" json:"user_id"`
	Content    string `gorm:"type:varchar(255);index:name" json:"content"`
	Images     string `gorm:"type:varchar(255);index:images" json:"images"`
	Like       int    `gorm:"type:int(20);index:like" json:"like"`
	State      int    `gorm:"type:int(5);index:state" json:"state"`
	Createtime string `gorm:"type:varchar(50);index:createtime" json:"createtime"`
}
