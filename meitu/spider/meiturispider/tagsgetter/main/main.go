package main

import (
	"../../../.."
	"../../../../model"
	"../../gorm"
	"strconv"
	"net/http"
	"fmt"
	"time"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	for i := 100; i<500; i++ {
		url := meitu.Host + "/s/" + strconv.Itoa(i)
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
		fmt.Println(name+des+shuliang)
		tag := model.Tags{
			Id:i,
			Name: name,
			Des:  des,
			Nums:  num,
		}
		gorm.SaveTagInfo(tag)
	}
}