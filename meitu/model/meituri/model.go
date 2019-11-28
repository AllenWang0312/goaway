package meituri

import (
	"encoding/json"
	"strings"
	"time"
)

type Model struct {
	Tab
	Cover string `gorm:"type:varchar(100);index:cover" json:"cover"`
	//由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但string则可以接收nil值
	Nicknames     string `gorm:"type:varchar(100);index:nicknames" json:"nickname"`
	Birthday      string `gorm:"type:varchar(20);index:birthday" json:"birthday"`
	Constellation string `gorm:"type:varchar(10);index:constellation" json:"constellation"`
	Height        string `gorm:"type:varchar(10);index:height" json:"height"`
	Weight        string `gorm:"type:varchar(10);index:weight" json:"weight"`
	Dimensions    string `gorm:"type:varchar(30);index:dimensions" json:"dimensions"`
	Address       string `gorm:"type:varchar(100);index:address" json:"address"`
	Jobs          string `gorm:"type:varchar(100);index:jobs" json:"jobs"`
	Interest      string `gorm:"type:varchar(200);index:interest" json:"interest"`
	More          string `gorm:"type:varchar(255);index:more" json:"more"`
	Tags          string `gorm:"type:varchar(255);index:tags" json:"tags"`
	Get           bool   `gorm:"-" json:"get"`
}
type Album struct {
	ID      int `gorm:"primary_key;index:id" json:"id"`
	Modelid int `gorm:"type:int(11)" json:"model_id"`
	Groupid int `gorm:"type:int(11)" json:"group_id"`

	Title string `gorm:"type:varchar(100);index:title" json:"title"`
	Tags  string `gorm:"type:varchar(100);index:tags" json:"tags"`
	Subs  string `gorm:"type:varchar(500);index:subs" json:"subs"`
	Group string `gorm:"type:varchar(100);index:group" json:"org"`
	No    string `gorm:"type:varchar(20);index:no" json:"no"`

	Nums int    `gorm:"type:int(11)" json:"nums"`
	Time string `gorm:"type:varchar(30);index:time" json:"time"`
	Hot  int    `gorm:"type:int(16)" json:"hot"`

	Images []string `gorm:"-" json:"images"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
}

type Tag struct {
	Tab
	Des      string `gorm:"type:varchar(200);index:des" json:"des"`
	FullName string `gorm:"index:fullname" json:"fullname"`
	Nums     int    `gorm:"index:nums" json:"nums"`
}
type Company struct {
	Tab
	Count int `gorm:"index:count" json:"count"`
}
type Group struct {
	Tab
	Belong   int    `gorm:"type:int(20);index:belong"`
	Homepage string `gorm:"type:varchar(255);index:homepage"`
}

type Tab struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `gorm:"index:name" json:"name"`
	Hot   int    `gorm:"index:hot" json:"hot"`
	Type  int    `gorm:"-" json:"type"`
	Alias string `gorm:"-" json:"alias"`
}

type Banner struct {
	ID    int    `gorm:"primary_key;index:id" json:"id"`
	State bool   `gorm:"state" json:"state"`
	Res   string `gorm:"res" json:"res"`
	Link  string `gorm:"link" json:"link"`
}

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

type Splash struct {
	ID     int    `gorm:"primary_key" json:"id"`
	Enable bool   `gorm:"index:enable" json:"enable"`
	Start  Time   `gorm:"index:start" json:"start"`
	End    Time   `gorm:"index:end" json:"end"`
	ArtUrl string `gorm:"index:art_url" json:"art_url"`

	Authorid int `gorm:"type:int(11)" json:"author_id"`
	//Author   User `gorm:"FOREIGNKEY Authorid" json:"author"`

	Modelid int `gorm:"type:int(11)" json:"model_id"`
	//Model Model `gorm:"FOREIGNKEY Modelid" json:"model"`
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

//feedback
type StringArray []byte
type Feedback struct {
	ID         int         `gorm:"primary_key" json:"id"`
	Userid     int         `gorm:"type:int(20);index:userid" json:"user_id"`
	Content    string      `gorm:"type:varchar(255);index:name" json:"content"`
	Images     StringArray `gorm:"type:varchar(255);index:images" json:"images"`
	Likes      int         `gorm:"type:int(20);index:likes" json:"likes"`
	State      int         `gorm:"type:int(5);index:state" json:"state"`
	Createtime time.Time   `gorm:"type:varchar(50);index:createtime" json:"createtime"`
}

func (t *StringArray) UnmarshalJSON(data []byte) (err error) {
	_ = json.NewDecoder(strings.NewReader(string(data))).Decode(&t)
	return
}

func (t StringArray) MarshalJSON() ([]byte, error) {
	return t, nil
}
