package object

import (
	"../../model"
	"database/sql"
	"fmt"
	gql "github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"strings"
)

func GenerateID() string {
	id, _ := uuid.NewV4()
	return strings.Split(id.String(), "-")[0]
}

var MutationType = gql.NewObject(gql.ObjectConfig{
	Name: "Mutation",
	Fields: gql.Fields{
		"createUser": &gql.Field{
			Type:        gql.Boolean,
			Description: "[用户管理] 创建用户",
			Args: gql.FieldConfigArgument{
				"userName": &gql.ArgumentConfig{
					Description: "用户名称",
					Type:        gql.NewNonNull(gql.String),
				},
				"email": &gql.ArgumentConfig{
					Description: "用户邮箱",
					Type:        gql.NewNonNull(gql.String),
				},
				"pwd": &gql.ArgumentConfig{
					Description: "用户密码",
					Type:        gql.NewNonNull(gql.String),
				},
				"phone": &gql.ArgumentConfig{
					Description: "用户联系方式",
					Type:        gql.Int,
				},
			},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				userId, _ := strconv.Atoi(GenerateID())
				user := &model.Users{
					Name: p.Args["userName"].(string),
					Email: sql.NullString{
						String: p.Args["email"].(string),
						Valid:  true,
					},
					Pwd:    p.Args["pwd"].(string),
					Phone:  int64(p.Args["phone"].(int)),
					UserID: uint64(userId),
					Status: int64(model.EnableStatus),
				}
				if err := model.InsertUser(user); err != nil {
					fmt.Println("[mutaition.createUser] invoke InserUser() failed")
					return false, err
				}
				return true, nil
			},
		},
		"changeUserName": &gql.Field{
			Type:        gql.Boolean,
			Description: "[用户管理] 修改用户名称",
			Args: gql.FieldConfigArgument{
				"userId": &gql.ArgumentConfig{
					Description: "用户ID",
					Type:        gql.NewNonNull(gql.Int),
				},
				"userName": &gql.ArgumentConfig{
					Description: "用户名称",
					Type:        gql.NewNonNull(gql.String),
				},
			},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				userId := p.Args["userId"].(uint64)
				name := p.Args["userName"].(string)
				if err := model.ChangeUserName(userId, name); err != nil {
					fmt.Println("[mutaition.changeUserName] invoke InserUser() failed")
					return false, err
				}
				return true, nil
			},
		},

		"deleteUser": &gql.Field{
			Type:        gql.Boolean,
			Description: "[用户管理] 删除用户",
			Args: gql.FieldConfigArgument{
				"userId": &gql.ArgumentConfig{
					Description: "用户ID",
					Type:        gql.NewNonNull(gql.Int),
				},
			},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				userId := p.Args["userId"].(uint64)
				if err := model.DeleteUser(userId, model.DisableStatus); err != nil {
					fmt.Println("[mutaition.deleteUser] invoke InserUser() failed")
					return false, err
				}
				return true, nil

			},
		},
	},
})
var QueryType = gql.NewObject(gql.ObjectConfig{
	Name: "Query",
	Fields: gql.Fields{
		"UserInfo": &gql.Field{
			Description: "[用户管理] 获取指定用户的信息",
			Type:        model.UserInfoType,
			Args: gql.FieldConfigArgument{
				"userId": &gql.ArgumentConfig{
					Description: "用户ID",
					Type:        gql.NewNonNull(gql.Int),
				},
			},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				userId := p.Args["userId"].(uint64)
				user, err := model.GetUser(userId)
				if err != nil {
					fmt.Println("[query.UserInfo] invoke InserUser() failed")
					return false, err
				}
				return model.Users{
					Name:   user.Name,
					UserID: user.UserID,
					Email:  user.Email,
					Phone:  user.Phone,
					Pwd:    user.Pwd,
					Status: user.Status,
				}, nil

			},
		},
		"UserListInfo": &gql.Field{
			Description: "[用户管理] 获取指定用户的信息",
			Type:        gql.NewNonNull(gql.NewList(model.UserInfoType)),
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				users, err := model.GetUsers()
				if err != nil {
					fmt.Println("[query.UserInfo] invoke InserUser() failed")
					return false, err
				}
				return users, nil
			},
		},
	},
})
