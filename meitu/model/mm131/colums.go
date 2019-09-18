package mm131

type Colums struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Time   string `json:"time"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
}
