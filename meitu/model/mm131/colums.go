package mm131

type Colums struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Time   string `json:"time"`
	Fenlei string `json:"fenlei"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
}
