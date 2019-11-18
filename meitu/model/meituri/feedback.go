package meituri

type Feedback struct {
	Id         int    `gorm:"primary_key"`
	Userid     int    `gorm:"type:int(20);index:userid"`
	Content    string `gorm:"type:varchar(255);index:name"`
	Images     string `gorm:"type:varchar(255);index:images"`
	Like       int    `gorm:"type:int(20);index:like"`
	State      int    `gorm:"type:int(5);index:state"`
	Createtime string `gorm:"type:varchar(50);index:createtime"`
}
