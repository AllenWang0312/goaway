package meituri

type Like struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	Userid   int    `gorm:"index:userid" json:"userid"`
	Modelid  int    `gorm:"index:modelid" json:"modelid"`
	Relation string `gorm:"index:relation" json:"relation"`
}
