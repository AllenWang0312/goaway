package meituri

type Models struct {
	ID            int    `gorm:"primary_key;index:id" json:"id"`
	Cover         string `gorm:"type:varchar(100);index:cover" json:"cover"`
	Name          string `gorm:"type:varchar(20);index:name" json:"name"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但string则可以接收nil值
	Nicknames     string `gorm:"type:varchar(100);index:nicknames" json:"nickname"`
	Birthday      string `gorm:"type:varchar(20);index:birthday" json:"birthday"`
	Constellation string `gorm:"type:varchar(10);index:constellation" json:"constellation"`
	Height        string `gorm:"type:varchar(10);index:height" json:"height"`
	Weight        string `gorm:"type:varchar(10);index:weight" json:"weight"`
	Dimensions    string `gorm:"type:varchar(30);index:dimensions" json:"dimensions"`
	Cup           string `gorm:"type:varchar(5);index:cup" json:"cup"`
	Address       string `gorm:"type:varchar(100);index:address" json:"address"`
	Jobs          string `gorm:"type:varchar(100);index:jobs" json:"jobs"`
	Interest      string `gorm:"type:varchar(200);index:interest" json:"interest"`
	More          string `gorm:"type:varchar(255);index:more" json:"more"`
	Tags          string `gorm:"type:varchar(255);index:tags" json:"tags"`
}
