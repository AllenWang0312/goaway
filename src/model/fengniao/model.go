package fengniao

type Album struct {
	ID      int `gorm:"primary_key;index:id" json:"id"`
	//Modelid int `gorm:"type:int(11)" json:"model_id"`
	//Model Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`
	//Groupid int `gorm:"type:int(11)" json:"group_id"`

	Title string `gorm:"type:varchar(100);index:title" json:"title"`
	//Tags  string `gorm:"type:varchar(100);index:tags" json:"tags"`
	Subs  string `gorm:"type:varchar(500);index:subs" json:"subs"`
	//Group string `gorm:"type:varchar(100);index:group" json:"org"`
	//No    string `gorm:"type:varchar(20);index:no" json:"no"`

	Nums int    `gorm:"type:int(11);index:nums" json:"nums"`
	Time string `gorm:"type:varchar(30);index:time" json:"time"`
	Class string `gorm:"type:varchar(20);index:class" json:"class"`
	//Hot  int    `gorm:"type:int(16)" json:"hot"`

	//Images []string `gorm:"-" json:"images"`
	//Cover  string   `gorm:"-" json:"cover"`
	//Cover ImageInfo `gorm:"-" json:"cover"`
	//Html    string `gorm:"type:varchar(255);index:html" json:"html"`
}
