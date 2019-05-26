package dao

type Models struct {
	ID            int  `gorm:"primary_key"`
	Cover         string `gorm:"type:varchar(100);index:cover"`
	Name          string `gorm:"type:varchar(20);index:name"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但string则可以接收nil值
	Nicknames     string `gorm:"type:varchar(100);index:nicknames"`
	Birthday      string `gorm:"type:varchar(20);index:birthday"`
	Constellation string `gorm:"type:varchar(10);index:constellation"`
	Height        int  `gorm:"index:height"`
	Dimensions    string `gorm:"type:varchar(30);index:dimensions"`
	Cup           string `gorm:"type:varchar(5);index:cup"`
	Address       string `gorm:"type:varchar(100);index:address"`
	Jobs          string `gorm:"type:varchar(100);index:jobs"`
	Interest      string `gorm:"type:varchar(200);index:interest"`
	More          string `gorm:"type:varchar(255);index:more"`
	Tags          string `gorm:"type:varchar(255);index:tags"`
}
