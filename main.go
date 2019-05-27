package main

import (
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//
	//"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"path"
	"net/http"
	"time"
	"os"
	"net/url"
	"strings"
	"github.com/jinzhu/gorm"
	"./dao"
	//"fmt"
	//"database/sql"
	"log"
	"bufio"
	"io"
)

const (
	SaveUserInfo  = true
	SaveColumInfo = true
)

var db *gorm.DB
var client *http.Client

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:Qunsi003@tcp(rm-wz952p7325m8jbe3x9o.mysql.rds.aliyuncs.com:3306)/meitu?charset=utf8&parseTime=True&loc=Local") //?charset=utf8&parseTime=True&loc=Local
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	client = http.DefaultClient
	client.Timeout = 20 * time.Second

	//client := http.DefaultClient
	//client.Timeout = 5 * time.Second
	//downloadItem(client,"27270",1)
	downloadUserColums([]int{101,288,298,918,})
	//downloadColums(455,[]int{26738,})
	//getUserColums("")
}

func downloadColums(userId int, colums []int) {
	for i := 0; i < len(colums); i++ {
		downloadSingleColum(userId, colums[i])
	}
}
func downloadUserColums(userIds []int) int {
	for i := 0; i < len(userIds); i++ {
		err := getUserColums(userIds[i])
		if err == -1 {
			continue
		} else if err == -2 {
			break
		}
	}
	return 0
}
func getUserColums(userId int) int {

	for i := 1; i < 10; i++ {
		err := analyzeHtml(client, userId, i)
		if (err == -1) {
			continue
		} else if (err == -2) {
			break
		}
	}
	return 0
}
func analyzeHtml(client *http.Client, userId int, i int) int {
	url := ""
	if i > 1 {
		url = "https://www.meituri.com/t/" + strconv.Itoa(userId) + "/" + strconv.Itoa(i) + ".html"
	} else {
		url = "https://www.meituri.com/t/" + strconv.Itoa(userId) + "/"
	}
	resp, err := client.Get(url)
	if (resp.StatusCode > 400) {
		return -2
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return -1
	}
	if (i == 1 && SaveUserInfo) {
		saveUserInfo(userId, doc)
	}
	println(doc.Url.String())
	hezi := doc.Find("div.hezi")
	if hezi != nil {
		hezi.Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
			path, _ := s.Find("a").Attr("href")
			paths := strings.Split(path, "/")
			colum, _ := strconv.Atoi(paths[len(paths)-2])
			println(path + " " + strconv.Itoa(len(path)) + strconv.Itoa(colum))

			err := downloadColum(userId, colum)
			if err < 0 {
				//continue
			}
		})
	} else {
		return -1
	}

	return 0
}
func saveUserInfo(userId int, doc *goquery.Document) {
	cover, _ := doc.Find("div.left").Find("img").Attr("src")
	right := doc.Find("div.right")
	nicknames := right.Find("h1").Text()
	shuoming := doc.Find("div.shuoming")
	more := shuoming.Text()
	tags := shuoming.Find("p").Text()
	model := dao.Models{
		ID:        userId,
		Cover:     cover,
		Name:      strings.Split(nicknames, "、")[0],
		Nicknames: nicknames,
		More:      more,
		Tags:      tags,
	}

	if err := db.Create(model).Error; err != nil {
		//return -3
		println(err.Error())
	}
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func downloadColum(userId int, colum int) int {
	//saveColumRelation(userId, colum)
	downloadSingleColum(userId, colum)
	return 0
}
func downloadSingleColum(userId int, colum int) {
	if (SaveColumInfo) {
		doc, err := goquery.NewDocument("https://www.meituri.com/a/" + strconv.Itoa(colum))
		if err != nil {
			//return -1
			println(err.Error())
		}
		saveColumInfo(colum, userId, doc)
	}
	os.MkdirAll("./data/t/"+strconv.Itoa(userId)+"/"+strconv.Itoa(colum), os.ModePerm)
	for i := 1; i < 100; i++ {

		err := downloadItem(userId, colum, i)
		if (err == -1) {
			continue
		} else if err == -2 {
			break
		}
	}
}
func saveColumInfo(colum int, userId int, doc *goquery.Document) int {
	tuji := doc.Find("div.tuji")
	html, _ := tuji.Html()
	weizhi := tuji.Find("div.weizhi")
	if weizhi != nil {
		title := weizhi.Find("h1").Text()
		subs := doc.Find("div.shuoming").Find("p").Text()
		println(title + subs)
		c := dao.Colums{
			ID:     colum,
			Userid: userId,
			Title:  title,
			Subs:   subs,
			Html:   html,
		}
		db.Create(c)
		//if db.NewRecord(c) {
		//	db.Create(c)
		//	println("colum(" + strconv.Itoa(colum) + ") info insert")
		//} else {
		//	db.Update(c)
		//	println("colum info updata")
		//}
		return 0
	} else {
		return -1
	}

}
func saveColumRelation(userId int, colum int) {
	c := dao.Colums{
		ID:     colum,
		Userid: userId,
	}
	if err := db.Create(c).Error; err != nil {
		//return -3
		println(err.Error())
	}
}
func downloadItem(userId int, colum int, i int) int {
	durl := "https://ii.hywly.com/a/1/" + strconv.Itoa(colum) + "/" + strconv.Itoa(i) + ".jpg"
	uri, err := url.ParseRequestURI(durl)
	if err != nil {
		//panic("url err")
		println(err.Error())
		return -2
	}
	filename := path.Base(uri.Path)
	return downloadFile(durl, "./data/t/"+strconv.Itoa(userId)+"/"+strconv.Itoa(colum)+"/", filename)
}
func downloadFile(durl string, path string, name string) int {
	log.Println("filename " + name)
	e, _ := PathExists(path + name)
	if e {
		return -1
	}
	resp, err := client.Get(durl)
	if err != nil {
		//panic(err)
		println(err.Error())
		return -2
	}
	if resp.StatusCode > 400 {
		return -2
	}
	if resp.ContentLength <= 0 {
		log.Println("content length < 0")
	}

	raw := resp.Body
	defer raw.Close()
	reader := bufio.NewReaderSize(raw, 1024*32)

	file, err := os.Create(path + name)
	if err != nil {
		//panic(err)
		println(err.Error())
	}
	writer := bufio.NewWriter(file)
	buff := make([]byte, 32*1024)
	written := 0
	//go
	func() {
		for {
			nr, er := reader.Read(buff)
			if nr > 0 {
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			//panic(err)
			println(err.Error())
		}
	}()

	spaceTime := time.Second * 1
	ticker := time.NewTicker(spaceTime)
	lastWtn := 0
	stop := false
	for {
		select {
		case <-ticker.C:
			speed := written - lastWtn
			log.Printf("speed %s /%s \n", bytesToSize(speed), spaceTime.String())
			if written-lastWtn == 0 {
				ticker.Stop()
				stop = true
				break
			}
			lastWtn = written
		}
		if stop {
			break
		}
	}
	return 0
}
func bytesToSize(length int) string {
	var k = 1024
	var sizes = []string{"Nytes", "KB", "MB", "GB", "TB"}
	if length == 0 {
		return "0 Bytes"
	}
	i := math.Floor(math.Log(float64(length)) / math.Log(float64(k)))
	r := float64(length) / math.Pow(float64(k), i)
	return strconv.FormatFloat(r, 'f', 3, 64) + sizes[int(i)]
}
