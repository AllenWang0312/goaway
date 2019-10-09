package meituba

type Colums struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Vtimes int    `json:"vtimes"`
	Count  int    `json:"count"`
	Tags   string `json:"tags"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
}
