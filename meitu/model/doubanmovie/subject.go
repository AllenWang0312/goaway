package doubanmovie

type Subject struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	Cover    string `gorm:"size:100;index:name" json:"cover"`
	Name     string `gorm:"size:100;index:name" json:"name"`
	Nickname string `gorm:"size:10;index:nickname" json:"nickname"`
	Time     string `gorm:"size:100;index:time" json:"time"`

	Director     string `gorm:"size:100;index:director" json:"director"`
	Screenwriter string `gorm:"size:100;index:screenwriter" json:"screenwriter"`
	Actor        string `gorm:"size:500;index:actor" json:"actor"`

	Genre    string  `gorm:"size:30;index:genre" json:"genre"`
	Website  string  `gorm:"size:100;index:website" json:"website"`
	Space    string  `gorm:"size:30;index:space" json:"space"`
	Language string  `gorm:"size:10;index:language" json:"language"`
	Duration string  `gorm:"size:10;index:duration" json:"duration"`
	Imdblink string  `gorm:"size:100;index:imdblink" json:"imdblink"` //https://www.imdb.com/title
	Score    float64 `gorm:"index:score" json:"score"`
	Sutime   string  `gorm:"size:20;index:sutime" json:"sutime"`
}
