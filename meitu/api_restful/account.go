package api_restful

import (
	model "../model/meituri"
	"crypto/aes"
	"fmt"
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
func checkTokenEnable(token string) bool {
	aes.NewCipher(conf.SecretKey)
	return strings.EqualFold(token, "token")
}

func resetPass(c *gin.Context) {

}

//todo
func tokenLogin(c *gin.Context) {

}
func Login(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("pwd")
	//base64.StdEncoding.EncodeToString(hashData)

	//pwd:=c.PostForm("pwd")
	user := model.Users{}
	db.First(&user, "account = ?", account)
	fmt.Println(user.Info())

	if user.ID > 0 {
		if strings.EqualFold(user.Pwd, pwd) {

			user.Token = "token"
			c.JSON(200, gin.H{"msg": "密码正确",
				"data": &user})
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
func GetUser(userId uint64) (model.Users, error) {
	return model.Users{}, nil
}
func GetUsers() ([]model.Users, error) {
	return []model.Users{}, nil
}
