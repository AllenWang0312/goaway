package main

import (
	"fmt"
	"regexp"
)

const text = "My email is ccmouse@gmail.com"

//const text = `<a title="高清拍摄泳池美女性感泳衣诱惑照片" href="http://www.meituba.com/xinggan/93494.html" target="_blank">
//<img src="http://ppic.meituba.com:83/uploads3/180905/3-1PZ50T914H8.jpg" style="display: inline;">
//</a>`
func main() {
	re := regexp.MustCompilePOSIX(`([a-zA-Z0-9])+@([a-zA-Z0-9.])+(\.[a-zA-Z0-9]+)`)
	//re := regexp.MustCompilePOSIX(`<a title="+(.+)+" href="+(.+)+" target="+(\.[a-zA-Z0-9]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
