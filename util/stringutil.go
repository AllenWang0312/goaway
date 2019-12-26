package util

import (
	"github.com/henrylee2cn/pholcus/common/mahonia"
	"regexp"
	"strconv"
	"strings"
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/transform"
	//"github.com/axgle/mahonia"
)

//用户名：  	/^[a-z0-9_-]{3,16}$/
//密码：	    /^[a-z0-9_-]{6,18}$/
//十六进制值：	/^#?([a-f0-9]{6}|[a-f0-9]{3})$/
//电子邮箱	：  /^([a-z0-9_\.-]+)@([\da-z\.-]+)\.([a-z\.]{2,6})$/
///^[a-z\d]+(\.[a-z\d]+)*@([\da-z](-[\da-z])?)+(\.{1,2}[a-z]+)+$/
//URL： 	    /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/
//IP 地址：	/((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)/
///^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
//HTML 标签：	/^<([a-z]+)([^<]+)*(?:>(.*)<\/\1>|\s+\/>)$/
//
//删除代码\\注释：      	(?<!http:|\S)//.*$
//Unicode编码中的汉字范围：	/^[\u2E80-\u9FFF]+$/

func IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
func IsMobile(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 型如 https://meituri.com/x/1/1096/ 获取1096
func GetIdFromUri(url string) int {
	ids := strings.Split(url, "/")
	if len(ids) > 2 {
		id, err := strconv.Atoi(ids[len(ids)-2])
		if nil == err {
			return id
		}
	}
	return 0
}

// 型如 https://meituri.com/x/1/sdfjlsf.jpg 获取sdfjlsf.jpg
func GetPathFromUri(url string) string {
	index := strings.LastIndex(url, "?")
	return url[0:index]
}

// 型如 https://meituri.com/x/1/sdfjlsf.jpg 获取sdfjlsf.jpg
func GetNameFromUri(url string) string {
	paths := strings.Split(url, "/")
	return paths[len(paths)-1]
}
func GetRouteFromUri(url string) string {
	paths := strings.Index(url, ".com")
	index := strings.LastIndex(url, "/")
	return url[paths+4 : index]
}
// 型如 https://meituri.com/x/1/1096_3.html 获取1096
func GetNameIDFromUri(url string) string {
	paths := strings.LastIndex(url, "/")
	index := strings.LastIndex(url, ".")
	fg:=strings.LastIndex(url,"_")
	if fg>0 {
		return url[paths+1 : fg]
	}else{
		return url[paths+1 : index]
	}

}

//const (
//	UTF8    = "UTF-8"
//	GB18030 = "GB18030"
//)

//func ConvertStr2GBK(str string) string {
//	//将utf-8编码的字符串转换为GBK编码
//	 ret, err := simplifiedchinese.GBK.NewEncoder().String(str)
//	 return ret //如果转换失败返回空字符串
//
//	//如果是[]byte格式的字符串，可以使用Bytes方法
//	b, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
//	return string(b)
//}
func G2U(str string) string {
	dec := mahonia.NewDecoder("gbk")
	if ret, ok := dec.ConvertStringOK(str); ok {
		return ret
	}
	return ""
}
