package doubanmovie

type Celebrity struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	Name     string `gorm:"size:20;index:name" json:"name"`
	//Nickname sql.NullString `gorm:"size:20;index:nickname" json:"nickname,omitempty"`
	//Gender   sql.NullString `gorm:"size:4;index:gender" "json:"gender,omitempty"`
}

//type Celebrity struct {
//	ID       int    `gorm:"primary_key;index:id" json:"id"`
//	Name     string `gorm:"size:20;index:name" json:"name"`
//	Nickname sql.NullString `gorm:"size:20;index:nickname" json:"nickname,omitempty"`
//	Gender   sql.NullString `gorm:"size:4;index:gender" "json:"gender,omitempty"`
//}