package api_restful

import (
	"../../conf"
	"../cache"
	"../encrypt"
	model "../model/meituri"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
)

func tokenEnable(c *gin.Context) bool {
	token := c.GetHeader("token")
	if !checkTokenEnable(token) {
		c.JSON(401, gin.H{"status": -1, "msg": "token已失效"})
		return false
	} else {
		return true
	}
}

//func logedUser(c *gin.Context) bool {
//
//}
func checkTokenEnable(token string) bool {
	//aes.NewCipher([]byte(conf.AESSecretKey))
	_, err := cache.Get(token)
	if nil == err {
		return true
	} else {
		return false
	}

}

func resetPass(c *gin.Context) {

}

//
func TokenLogin(c *gin.Context) {
	//token := c.GetHeader("token")
	//if len(token) > 0 {
	//	v ,err:= cache.Get(token)
	//	if nil==err {
	//		user := model.User{}
	//		if err := json.Unmarshal([]byte(v), &user); err == nil {
	//			c.JSON(200, gin.H{"data": user})
	//		}
	//	} else {
	//		c.JSON(200, gin.H{"toast": "获取token失败:"+err.Error()})
	//	}
	//} else {
	//	c.JSON(200, gin.H{"toast": "token 为空"})
	//}

	user := getUserWithToken(c)
	if (user.ID > 0) {
		c.JSON(200, gin.H{"data": user})
	} else {
		c.JSON(200, gin.H{"toast": "token 登录失败"})
	}
}
func Login(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("pwd")
	//base64.StdEncoding.EncodeToString(hashData)
	//pwd:=c.PostForm("pwd")
	user := model.User{}
	db.Where("account = ?", account).First(&user)
	if user.ID > 0 {
		if strings.EqualFold(user.Pwd, pwd) {
			token := encrypt.RandToken(16)
			b, e := json.Marshal(user)
			if nil == e {
				err := cache.Set(token, string(b), 30*conf.Day)
				if nil == err {
					user.Token = token
					c.JSON(200, gin.H{"toast": "密码正确",
						"data": &user})
				} else {
					c.JSON(200, gin.H{"toast": err.Error()})
				}
			} else {
				c.JSON(200, gin.H{"msg": e.Error()})
			}
		} else {
			c.JSON(200, gin.H{"status": -1, "msg": "确认密码不符"})
		}
	} else {
		c.JSON(200, gin.H{"msg": "用户不存在"})
	}
}

func RegistAccount(c *gin.Context) {
	//c.Header("tel","")
	tel := c.PostForm("tel")
	pwd := c.PostForm("pwd")
	repwd := c.PostForm("repwd")
	//fmt.Println(pwd + repwd)
	if strings.EqualFold(pwd, repwd) {
		email := c.PostForm("email")
		if len(tel) > 0 {
			if len(pwd) > 6 {
				user := model.User{
					Account: tel,
					Tel:     tel,
					Pwd:     pwd,
				}
				createSuccess := db.NewRecord(&user)
				if createSuccess {
					if err := db.Create(&user).Error; err != nil {
						//return -3
						println(err.Error())
					}
				} else {
					c.JSON(200, gin.H{"status": -1, "msg": "创建失败"})
				}

			} else {
				c.JSON(200, gin.H{"status": -1, "msg": "密码过短"})
			}
		} else if len(email) > 4 {
			if len(pwd) > 6 {
				user := model.User{
					Account: email,
					Email:   email,
					Pwd:     pwd,
				}
				createSuccess := db.NewRecord(&user)
				if createSuccess {
					if err := db.Create(&user).Error; err != nil {
						//return -3
						println(err.Error())
					}
				} else {
					c.JSON(200, gin.H{"status": -1, "msg": "创建失败"})
				}
			} else {
				c.JSON(200, gin.H{"status": -1, "msg": "密码过短"})
			}
		}
	} else {
		c.JSON(200, gin.H{"status": -1, "msg": "确认密码不符"})
	}
	c.JSON(200, gin.H{"status": 1, "msg": "创建成功"})
}

func GetUser(c *gin.Context) {
	var user_id = getUserIdWithToken(c)
	if user_id == -1 {
		return
	} else if user_id > 0 {
		var user = model.User{}
		db.Where("id = ?", user_id).First(&user)
		c.JSON(200, gin.H{"data": user})
	} else {
		c.JSON(404, gin.H{"msg": "用户不存在"})
	}
}

func getUserWithToken(c *gin.Context) (model.User) {
	//aes.NewCipher([]byte(conf.AESSecretKey))
	var token = c.GetHeader("token")
	user_str, err := cache.Get(token)
	var user = model.User{}
	if nil == err {
		err := json.NewDecoder(strings.NewReader(string(user_str))).Decode(&user)
		if nil == err {
		} else {
			c.JSON(200, gin.H{"msg": err.Error()})
		}
		return user
		//json.NewDecoder().Decode(user_str,&user)
	} else {
		c.JSON(200, gin.H{"msg": "token 获取失败:" + err.Error()})
		return user
	}
}
func getUserIdWithToken(c *gin.Context) int {
	//aes.NewCipher([]byte(conf.AESSecretKey))
	var token = c.GetHeader("token")
	print(token)
	user_str, err := cache.Get(token)
	print(user_str)
	if nil == err {
		var user = model.User{}
		err := json.NewDecoder(strings.NewReader(string(user_str))).Decode(&user)
		if nil == err {
			return user.ID
		} else {
			c.JSON(200, gin.H{"msg": "token 有效 解析失败:" + err.Error()})
			return -1
		}
	} else {
		c.JSON(200, gin.H{"msg": "token 获取出错:" + err.Error()})
		return -1
	}
}

//func EditUserInfo(c *gin.Context){
//	var
//	var name=c.PostForm("name")
//	var user=model.User{}
//
//	if user.ID>0 {
//		c.JSON(200, gin.H{"data": user})
//	}else {
//		c.JSON(404, gin.H{"msg": "用户不存在"})
//	}
//}
//todo
func InsertUser(user *model.User) error {

	return nil
}
func ChangeUserName(userId uint64, userNaem string) error {

	return nil
}
func DeleteUser(userId uint64, statusType model.UserStatusType) error {

	return nil
}
