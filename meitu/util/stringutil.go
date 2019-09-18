package util

import (
	"strconv"
	"strings"
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

// 型如 https://meituri.com/x/1/sdfjlsf.jpg 获取1096
func GetNameFromUri(url string) string {
	paths := strings.Split(url, "/")
	return paths[len(paths)-1]
}
