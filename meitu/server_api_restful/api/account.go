package api

import (
	"../../../conf"
	"../../../util"
	"../../cache"
	"../../encrypt"
	model "../../model/meituri"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
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
	if user.ID > 0 {
		platform := c.Query("platform")
		device := c.PostForm("device")

		c.JSON(200, gin.H{"data": model.LoginResp{
			User: user,
			Tab:  getFirstTabFollowID(user.ID) > 0,
			Bind: hasThisDeviceBinded(user.ID, platform, device),
		}})
	} else {
		c.JSON(200, gin.H{"toast": "token 登录失败"})
	}
}

func Login(c *gin.Context) {
	platform := c.Query("platform")

	account := c.PostForm("account")
	pwd := c.PostForm("pwd")
	device := c.PostForm("device")

	//base64.StdEncoding.EncodeToString(hashData)
	//pwd:=c.PostForm("pwd")
	user := model.User{}
	db.Where("account = ?", account).First(&user)

	if user.ID > 0 {
		user_id := strconv.Itoa(user.ID)

		if strings.EqualFold(user.Pwd, pwd) {
			token := platform + "_" + device
			b, e := json.Marshal(user)
			if nil == e {
				err2 := cache.Set(token, user_id, 30*conf.Day)
				err1 := cache.SetV(user_id, string(b))

				if nil == err1 {
					if nil == err2 {
						user.Token = token

						c.JSON(200, gin.H{"toast": "密码正确",
							"data": model.LoginResp{
								User: user,
								Tab:  getFirstTabFollowID(user.ID) > 0,
								Bind: hasThisDeviceBinded(user.ID, platform, device),
							}})
					} else {
						c.JSON(200, gin.H{"msg": err2.Error()})
					}
				} else {
					c.JSON(200, gin.H{"msg": err1.Error()})
				}
			} else {
				c.JSON(200, gin.H{"msg": e.Error()})
			}
		} else {
			c.JSON(200, gin.H{"code": -1, "toast": "确认密码不符"})
		}
	} else {
		c.JSON(200, gin.H{"code": -1, "toast": "用户不存在"})
	}
}

func hasThisDeviceBinded(userid int, platform string, device string) bool {
	var bindDevices [] model.BindDevice
	db.Where("userid = ", userid).Find(&bindDevices)

	var bind = false
	for i, k := range bindDevices {
		println(i, k.Key)
		if k.Key == platform+"_"+device {
			bind = true
		}
	}
	return bind
}

func getFirstTabFollowID(userid int) int {
	var frecord = model.FollowTab{
		Userid:userid,
	}
	db.First(&frecord)
	println(frecord.ID)
	return frecord.ID
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

func Regist(c *gin.Context) {
	account := c.PostForm("account")
	pwd := c.PostForm("pwd")
	if len(account) > 0 && len(pwd) > 0 {
		var user = model.User{}
		if util.IsMobile(account) {

			user = model.User{
				Account: account,
				Pwd:     pwd,
				Tel:     account,
			}
		} else {
			c.JSON(200, gin.H{"toast": "暂时只支持手机号注册哦"})
			return
		}
		//else if util.IsEmail(account) {
		//	user = model.User{
		//		Account: account,
		//		Pwd:     pwd,
		//		Email:   account,
		//	}
		//}

		//new := db.NewRecord(&user)
		//if new {
		db.Create(&user)
		println(user.ID)
		var code = encrypt.AesCBCDecrypt(strconv.Itoa(user.ID), encrypt.Key)
		user.InvitateCode = code
		db.Where("id = ", user.ID).Update("invitatecode", code)
		c.JSON(200, gin.H{"toast": "创建成功",
			"data": user})
		//} else {
		//	c.JSON(200, gin.H{"toast": "用户已存在"})
		//}
	}
}
func GetUser(c *gin.Context) {
	var user = getUserWithToken(c)
	c.JSON(200, gin.H{"data": &user})
}
func GetUserInfo(c *gin.Context) {
	var user_id = getUserIdWithToken(c)
	if user_id > 0 {
		var user = model.User{}
		db.Where("id = ?", user_id).First(&user)
		var roles []model.RoleRecord
		db.Table("user_role"+strconv.Itoa(user_id/1000)).Where("userid = ?",user_id).Find(&roles)
		var usercenter=model.UserCenter{
			User:user,
			Roles:roles,
		}
		c.JSON(200, gin.H{"data": usercenter})
	} else {
		c.JSON(200, gin.H{"code": -1, "toast": "用户不存在"})
	}
}

func getUserWithToken(c *gin.Context) model.User {
	//aes.NewCipher([]byte(conf.AESSecretKey))
	var token = c.GetHeader("token")
	userStr, err := cache.GetSecondaryToken(token)
	var user = model.User{}
	//if len(userStr)>0 {
	if nil == err {
		err := json.NewDecoder(strings.NewReader(string(userStr))).Decode(&user)
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
	//}
}
func getUserIdWithToken(c *gin.Context) int {
	//aes.NewCipher([]byte(conf.AESSecretKey))
	var token = c.GetHeader("token")
	//print(token)
	user_id, err := cache.Get(token)
	id, err1 := strconv.Atoi(user_id)
	if err == nil && err1 == nil {
		return id
	} else {
		return 0
	}
	//user_str, err := cache.GetSecondaryToken(token)
	//print(user_str)
	//if nil == err {
	//	var user = model.User{}
	//	err := json.NewDecoder(strings.NewReader(string(user_str))).Decode(&user)
	//	if nil == err {
	//		return user.ID
	//	} else {
	//		c.JSON(200, gin.H{"msg": "token 有效 解析失败:" + err.Error()})
	//		return -1
	//	}
	//} else {
	//	c.JSON(200, gin.H{"msg": "token 获取出错:" + err.Error()})
	//	return -1
	//}
}
func getUserIdStrWithToken(c *gin.Context) string {
	//aes.NewCipher([]byte(conf.AESSecretKey))
	var token = c.GetHeader("token")
	//print(token)
	user_id, err := cache.Get(token)
	if err == nil {
		return user_id
	} else {
		return "0"
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
