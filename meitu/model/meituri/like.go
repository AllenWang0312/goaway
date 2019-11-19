package meituri

type LikeModel struct {
	ID     int `gorm:"primary_key;index:id" json:"id"`
	Userid int `gorm:"index:userid" json:"userid"`

	Modelid int   `gorm:"index:modelid" json:"modelid"`
	Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`

	Relation string `gorm:"index:relation" json:"relation"`
}

type LikeColum struct {
	ID      int   `gorm:"primary_key;index:id" json:"id"`
	Userid  int   `gorm:"index:userid" json:"userid"`
	Modelid int   `gorm:"index:modelid" json:"modelid"`
	Columid int   `gorm:"index:columid" json:"columid"`
	Colum   Colum `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Columid" json:"colum"`

	Relation string `gorm:"index:relation" json:"relation"`
}
