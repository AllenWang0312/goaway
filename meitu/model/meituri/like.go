package meituri

type Like struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	Userid   int    `gorm:"index:userid" json:"userid"`
	Modelid  int    `gorm:"index:modelid" json:"modelid"`
	Columid  int    `gorm:"index:columid" json:"columid"`
	Relation string `gorm:"index:relation" json:"relation"`
}
