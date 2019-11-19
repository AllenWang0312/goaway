package meituri

import (
	"encoding/json"
	"strings"
	"time"
)

type StringArray []byte

type Feedback struct {
	Id         int         `gorm:"primary_key" json:"id"`
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
