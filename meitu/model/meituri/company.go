package meituri

type Company struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `gorm:"index:name" json:"name"`
	Count int    `gorm:"index:count" json:"count"`
	Hot   int    `gorm:"index:hot" json:"hot"`
}
