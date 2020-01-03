package mzitu

type Colums struct {
	ID    int    `gorm:"primary_key;index:id" json:"id"`
	Title string `gorm:"size:100;index:title" json:"title"`
	Time  string `gorm:"size:20;index:time" json:"time"`
}
