package meituri

import (
	"../../conf"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type Model struct {
	Tab
	Cover string `gorm:"type:varchar(100);index:cover" json:"cover"`
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
	ID      int   `gorm:"primary_key;index:id" json:"id"`
	Modelid int   `gorm:"type:int(11)" json:"model_id"`
	Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`
	Groupid int   `gorm:"type:int(11)" json:"group_id"`

	Title string `gorm:"type:varchar(100);index:title" json:"title"`
	Tags  string `gorm:"type:varchar(100);index:tags" json:"tags"`
	Subs  string `gorm:"type:varchar(500);index:subs" json:"subs"`
	Group string `gorm:"type:varchar(100);index:group" json:"org"`
	No    string `gorm:"type:varchar(20);index:no" json:"no"`

	Nums int    `gorm:"type:int(11)" json:"nums"`
	Time string `gorm:"type:varchar(30);index:time" json:"time"`
	Hot  int    `gorm:"type:int(16)" json:"hot"`

	Images []string `gorm:"-" json:"images"`
	//Cover  string   `gorm:"-" json:"cover"`
	Cover ImageInfo `gorm:"-" json:"cover"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
	Platform string `gorm:"-" json:"platform"`
	Column   string `gorm:"type:varchar(20);index:class" json:"column"`
}
type ImageInfo struct {
	Url    string  `json:"url"`
	Scale  float64 `json:"scale"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
}

func GetAlbumCover(a *Album, net bool) string {
	if (net) {
		return conf.FILE_SERVER + "/muri/" + strconv.Itoa(a.Modelid) + "/" + strconv.Itoa(a.ID) + "/0.jpg"
	} else {
		return conf.FSRoot + conf.Muri + "/" + strconv.Itoa(a.Modelid) + "/" + strconv.Itoa(a.ID) + "/0.jpg"
	}
}

type Tag struct {
	Tab
	Des      string `gorm:"type:varchar(200);index:des" json:"des"`
	FullName string `gorm:"type:varchar(100);index:fullname" json:"fullname"`
	Nums     int    `type:int(11);gorm:"index:nums" json:"nums"`
}
type Company struct {
	Tab
	Count int `gorm:"type:int(11);index:count" json:"count"`
}
type Group struct {
	Tab
	Belong   int    `gorm:"type:int(11);index:belong"`
	Homepage string `gorm:"type:varchar(255);index:homepage"`
}

type Tab struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `gorm:"type:varchar(100);index:name" json:"name"`
	Hot   int    `gorm:"type:int(20);index:hot" json:"hot"`
	Type  int    `gorm:"-" json:"type"`
	Alias string `gorm:"-" json:"alias"`
}

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

type Banner struct {
	ID     int    `gorm:"primary_key;index:id" json:"id"`
	Type   int    `gorm:"index:type" json:"type"`
	Enable bool   `gorm:"enable" json:"enable"`
	Start  Time   `gorm:"index:start" json:"start"`
	End    Time   `gorm:"index:end" json:"end"`
	Src    string `gorm:"src" json:"src"`
	Link   string `gorm:"link" json:"link"`
}

type Role struct {
	ID       string `gorm:"primary_key;index:id" json:"id"`
	Rolename string `gorm:"varchar(20);index:rolename" json:"role_name"`
	des      string `gorm:"varchar(255);index:id" json:"des"`
}
type RoleRecord struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	Userid   int    `gorm:"int(11);index:userid" json:"user_id"`
	Nick     string `gorm:"varchar(20);index:nick" json:"nick"`
	Des      string `gorm:"varchar(255);index:des" json:"des"`
	Type     int    `gorm:"int(11);index:type" json:"type"`
	Roleid   int    `gorm:"int(11);index:roleid" json:"role_id"`
	Relation int    `gorm:"varchar(255);index:relation" json:"relation"`
}
type UserRole struct {
	Role
	Name   string
	Userid int
	Jobid  int
}
type Splash struct {
	Banner

	Authors []UserRole
	//Authorid int `gorm:"type:int(11)" json:"author_id"`
	//Author   User `gorm:"FOREIGNKEY Authorid" json:"author"`

	//Modelid int `gorm:"type:int(11)" json:"model_id"`
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
	Tel        string      `gorm:"type:varchar(30);index:tel" json:"tel"`
	Createtime time.Time   `gorm:"type:varchar(50);index:createtime" json:"createtime"`
}

func (t *StringArray) UnmarshalJSON(data []byte) (err error) {
	_ = json.NewDecoder(strings.NewReader(string(data))).Decode(&t)
	return
}

func (t StringArray) MarshalJSON() ([]byte, error) {
	return t, nil
}
