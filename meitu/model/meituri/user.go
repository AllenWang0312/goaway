package meituri

import (
	"strconv"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Account  string `gorm:"type:varchar(20);index:account" json:"account"`
	Name     string `gorm:"type:varchar(20);index:name" json:"name"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但string则可以接收nil值
	Portarit string `gorm:"type:varchar(100);index:portarit" json:"portarit"`
	Email    string `gorm:"type:varchar(20);index:email" json:"email"`
	Pwd      string `gorm:"type:varchar(20);index:pwd" json:"pwd"`
	Tel      string `gorm:"type:varchar(20);index:tel" json:"tel"`
	Birthday string `gorm:"type:varchar(20);index:birthday" json:"birthday"`
	Token    string `gorm:"-" json:"token"`
	Type     int    `gorm:"type:Integer(10);index:type" json:"type"`
}

func (u *User) Info() string {
	return "id = " + strconv.Itoa(int(u.ID)) + "account = " + u.Account
}

type BindDevice struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserId   int    `gorm:"index:userid" json:"user_id"`
	Platform string `gorm:"index:platform" json:"platform"`
	DeviceID string `gorm:"index:deviceid" json:"device"`
	Key      string `gorm:"index:key" json:"key"`
}
type FollowTab struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserId   int    `gorm:"index:userid" json:"user_id"`
	ResId    int    `gorm:"index:resid" json:"res_id"`
	Type     int    `gorm:"index:type" json:"type"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikCompany struct {
	ID        int `gorm:"primary_key;index:id" json:"id"`
	Userid    int `gorm:"index:userid" json:"user_id"`
	CompanyId int `gorm:"index:companyid" json:"company_id"`
	//Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:ModelId" json:"model"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikGroup struct {
	ID      int `gorm:"primary_key;index:id" json:"id"`
	UserId  int `gorm:"index:userid" json:"userid"`
	GroupId int `gorm:"index:groupid" json:"groupid"`
	//Model   Model `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:ModelId" json:"model"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikeModel struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	UserId   int    `gorm:"index:user_id" json:"user_id"`
	ModelId  int    `gorm:"index:model_id" json:"model_id"`
	Model    Model  `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:ModelId" json:"model"`
	Relation string `gorm:"index:relation" json:"relation"`
}

type LikeAlbum struct {
	ID       int    `gorm:"primary_key;index:id" json:"id"`
	UserId   int    `gorm:"index:user_id" json:"user_id"`
	ModelId  int    `gorm:"index:model_id" json:"model_id"`
	AlbumId  int    `gorm:"index:album_id" json:"album_id"`
	Album    Album  `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:AlbumId" json:"album"`
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
