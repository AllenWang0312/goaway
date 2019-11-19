package api_restful

import (
	"../../configs"
	"../encrypt"
	model "../model/meituri"
	"../redis"
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
	_, err := redis.Get(token)
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
	token := c.PostForm("token")
	if(len(token)>0){
		userinfo, err := redis.Get(token)
		if nil == err {
			user := model.Users{}
			if err := json.Unmarshal([]byte(userinfo), &user); err == nil {
				c.JSON(200, gin.H{"data": user})
			}
		}else{
			c.JSON(200, gin.H{"toast": err.Error()})
		}
	}else{
		c.JSON(200, gin.H{"toast":"token 为空" })
	}


}
func Login(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("pwd")
	//base64.StdEncoding.EncodeToString(hashData)

	//pwd:=c.PostForm("pwd")
	user := model.Users{}
	db.Where("account = ?", account).First(&user)
	if user.ID > 0 {
		if strings.EqualFold(user.Pwd, pwd) {
			token := encrypt.RandToken(16)
			b, e := json.Marshal(user)
			if nil == e {
				err := redis.Set(token, string(b), 30*conf.Day)
				if nil == err {
					user.Token = token
					c.JSON(200, gin.H{"toast": "密码正确",
						"data": &user})
				}
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
				user := model.Users{
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
				user := model.Users{
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
	var user_id = c.PostForm("user_id")
	var user = model.Users{}
	db.Where("id = ?", user_id).First(&user)
	if user.ID > 0 {
		c.JSON(200, gin.H{"data": user})
	} else {
		c.JSON(404, gin.H{"msg": "用户不存在"})
	}
}

//func EditUserInfo(c *gin.Context){
//	var
//	var name=c.PostForm("name")
//	var user=model.Users{}
//
//	if user.ID>0 {
//		c.JSON(200, gin.H{"data": user})
//	}else {
//		c.JSON(404, gin.H{"msg": "用户不存在"})
//	}
//}
//todo
func InsertUser(user *model.Users) error {

	return nil
}
func ChangeUserName(userId uint64, userNaem string) error {

	return nil
}
func DeleteUser(userId uint64, statusType model.UserStatusType) error {

	return nil
}
