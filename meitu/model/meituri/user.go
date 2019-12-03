package meituri

import (
	"strconv"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `gorm:"type:varchar(50)" json:"account"`
	Name     string `gorm:"type:varchar(20)" json:"name"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但string则可以接收nil值
	Portarit string `gorm:"type:varchar(255)" json:"portarit"`
	Email    string `gorm:"type:varchar(30)" json:"email"`
	Pwd      string `gorm:"type:varchar(20)" json:"pwd"`
	Tel      string `gorm:"type:varchar(20)" json:"tel"`
	Birthday string `gorm:"type:varchar(20)" json:"birthday"`
	InvitateCode string `gorm:"type:varchar(100);index:invitatecode" json:"invitate_code"`
	Token    string `gorm:"-" json:"token"`
	//Type     int    `gorm:"type:Integer(10)" json:"type"`
}

func (u *User) Info() string {
	return "id = " + strconv.Itoa(int(u.ID)) + "account = " + u.Account
}

type Zone struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Type      int `gorm:"type:int(4)" json:"type"`
	Userid    int `gorm:"type:int(11)" json:"user_id"`
	Companyid int `gorm:"type:int(11)" json:"company_id"`
	Groupid   int `gorm:"type:int(11)" json:"group_id"`

	Modelid int   `gorm:"type:int(11)" json:"model_id"`
	Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`
	Albumid int   `gorm:"type:int(11);index:albumid;unique" json:"album_id"`
	Album   Album `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Albumid" json:"album"`

	Content string  `gorm:"type:varchar(255)" json:"content"`
	Time    string  `gorm:"type:varchar(32)" json:"time"`
	Address string  `gorm:"type:varchar(100)" json:"address"`
	Lat     float32 `gorm:"type:float(32)" json:"lat"`
	Long    float32 `gorm:"type:float(32)" json:"long"`
}
type BindDevice struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Userid   int    `json:"user_id"`
	Platform string `json:"platform"`
	Deviceid string `json:"device"`
	Key      string `json:"key"`
}
type FollowTab struct {
	ID     int `gorm:"primary_key" json:"id"`
	Userid int `json:"user_id"`

	Resid    int    `json:"res_id"`
	Type     int    `json:"type"`
	Tab      Tab    `gorm:"-" json:"tab"`
	Alias    string `json:"alias"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikCompany struct {
	ID        int `gorm:"primary_key" json:"id"`
	Userid    int `json:"user_id"`
	Companyid int `json:"company_id"`
	//Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikGroup struct {
	ID      int `gorm:"primary_key;index:id" json:"id"`
	UserId  int `json:"user_id"`
	GroupId int `json:"group_id"`
	//Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikeModel struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	Userid   int    `json:"user_id"`
	Modelid  int    `json:"model_id"`
	Model    Model  `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Modelid" json:"model"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikeAlbum struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Userid   int    `json:"user_id"`
	Modelid  int    `json:"model_id"`
	Albumid  int    `json:"album_id"`
	Album    Album  `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:Albumid" json:"album"`
	Relation string `gorm:"index:relation" json:"relation"`
}

//var UserStatusEnumType = graphql.NewEnum(graphql.EnumConfig{
//	Name:        "UserStatusEnum",
//	Description: "用户状态信息",
//	Values: graphql.EnumValueConfigMap{
//		"EnableUser": &graphql.EnumValueConfig{
//			Value:       EnableStatus,
//			Description: "用户可用",
//		},
//		"DisableUser": &graphql.EnumValueConfig{
//			Value:       DisableStatus,
//			Description: "用户不可用",
//		},
//	},
//})
//var UserInfoType = graphql.NewObject(graphql.ObjectConfig{
//	Name:        "userInfo",
//	Description: "用户信息描述",
//	Fields: graphql.Fields{
//		"userID": &graphql.Field{
//			Description: "用户ID",
//			Type:        graphql.Int,
//		},
//		"name": &graphql.Field{
//			Description: "用户名称",
//			Type:        graphql.String,
//		},
//		"email": &graphql.Field{
//			Description: "用户email",
//			Type:        graphql.String,
//		},
//		"phone": &graphql.Field{
//			Description: "用户手机号",
//			Type:        graphql.Int,
//		},
//		"pwd": &graphql.Field{
//			Description: "用户密码",
//			Type:        graphql.String,
//		},
//		"status": &graphql.Field{
//			Description: "用户状态",
//			Type:        UserStatusEnumType,
//		},
//	},
//})
