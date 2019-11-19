package util

import (
	"strconv"
	"strings"
	//"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/text/transform"
	//"github.com/axgle/mahonia"
)

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
func GetNameFromUri(url string) string {
	paths := strings.Split(url, "/")
	return paths[len(paths)-1]
}

// 型如 https://meituri.com/x/1/1096.html 获取1096

func GetNameIDFromUri(url string) string {
	paths := strings.LastIndex(url, "/")
	index := strings.LastIndex(url, ".")
	return url[paths+1 : index]
}

//const (
//	UTF8    = "UTF-8"
//	GB18030 = "GB18030"
//)
