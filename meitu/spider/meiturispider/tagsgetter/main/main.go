package main

import (
	"../../../../../conf"
	model "../../../../model/meituri"
	"../../gorm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"time"
)

func main() {
	for i := 100; i < 500; i++ {
		url := conf.Host + "/s/" + strconv.Itoa(i)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("request create faild: " + err.Error())
			continue
		}
		http.DefaultClient.Timeout = 20 * time.Second;
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("request error: " + err.Error())
			continue
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println("response status: " + strconv.Itoa(resp.StatusCode))
			continue
		}
		doc, err := goquery.NewDocument(url)
		fenlei := doc.Find("div.fenlei")
		shuliang := doc.Find("div.shoulushuliang").Find("span").Text()
		num, _ := strconv.Atoi(shuliang)
		name := fenlei.Find("h1").Text()
		des := fenlei.Find("p").Text()
		fmt.Println(name + des + shuliang)
		tag := model.Tag{
			Des:  des,
			Nums: num,
		}
		tag.ID = i
		tag.Name = name
		gorm.SaveTag(tag)
	}
}
