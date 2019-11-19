package meituri

import (
	"time"
)

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

	AuthorId int `gorm:"index:authorid" json:"author_id"`
	//Author   User `gorm:"FOREIGNKEY AuthorId" json:"author"`

	ModelId int `gorm:"index:modelid" json:"model_id"`
	//Model Model `gorm:"FOREIGNKEY ModelId" json:"model"`
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
