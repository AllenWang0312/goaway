package main

import (
	model "../../../model/mm131"
	"../../../util"
	"github.com/PuerkitoBio/goquery"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
	"strings"

	"../orm"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var h5_host = "https://www.mm131.net"
var host = "https://img1.mmmw.net"
var wg sync.WaitGroup

var client *http.Client
var downloadimg = true

func main() {
	orm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	paqufenlei("chemo", 1, 6000)
	//paqufenlei("chemo", 1, 6000)
	wg.Wait()
}

func paqufenlei(fenlei string, from int, to int) {
	//for i := to; i >= from; i-- {
	for i := from; i <= to; i++ {
		url := h5_host + "/" + fenlei + "/" + strconv.Itoa(i) + ".html"
		fenxi(i, fenlei, url)
		if i != 1 {
			old_url := h5_host + "/" + fenlei + "/1_" + strconv.Itoa(i) + ".html"
			fenxi(i, fenlei, old_url)
		}
	}
}

func fenxi(columId int, fenlei string, url string) int {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		//fmt.Println(resp.StatusCode)
		return -1
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	//simplifiedchinese.GB18030
	dec := mahonia.NewDecoder("gbk")
	content := doc.Find("div.content")
	//fmt.Println(content.Length())
	gbk_title := content.Find("h5").Text()
	title := dec.ConvertString(gbk_title)
	fmt.Println(title)
	gbk_msg := content.Find("div.content-msg").Text()
	msg := dec.ConvertString(gbk_msg)
	time := strings.Split(strings.Split(msg, "间：")[1], " M")[0]
	fmt.Println(msg)
	colum := model.Colums{
		ID:    columId,
		Title: title,
		Time:  time,
	}
	orm.SaveColum(&colum, "colums_"+fenlei)
	//orm.UpdateColumTime(&colum, "colums_"+fenlei)
	if downloadimg {
		downloadColum(columId, fenlei)
	}
	return 0
}

func downloadColum(columId int, fenlei string) {
	for i := 1; true; i++ {
		durl := host + "/pic/" + strconv.Itoa(columId) + "/" + strconv.Itoa(i) + ".jpg"
		//resp, err := url.ParseRequestURI(durl)
		//if err != nil {
		//	//panic("url err")
		//	println(err.Error())
		//	return -2
		//}
		filename := strconv.Itoa(i) + ".jpg"
		path := "../mm131/" + fenlei + "/" + strconv.Itoa(columId) + "/"
		//downloadFile(durl,path,filename)
		e, _ := util.PathExists(path + filename)
		if e {
			//fmt.Println("download file faild" + path + "/" + filename + "exist")
			continue
		}
		//filename := path.Base(uri.Path)
		req, err := http.NewRequest("GET", durl, nil)
		req.Header.Add("Referer", "https://www.mm131.net")
		//req.Header.Add("Host", "img1.mmmw.net:443")
		//req.Header.Add("Proxy-Connection", "keep-alive")
		//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
		if err != nil {
			fmt.Println("request create faild: " + err.Error())
			break
		}
		http.DefaultClient.Timeout = 20 * time.Second
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(durl + " request error: " + err.Error())
			break
		}
		if resp.StatusCode != http.StatusOK {
			fmt.Println(durl + " response status: " + strconv.Itoa(resp.StatusCode))
			break
		}
		wg.Add(1)
		go downloadImage(resp, path, filename)
	}
}

// 下载图片
func downloadImage(resp *http.Response, path string, fileName string) {
	//fileName := getNameFromUrl(url)
	defer func() {
		resp.Body.Close()
		if r := recover(); r != nil {
			//fmt.Println(r)
		}
		wg.Done()
	}()

	_ = os.MkdirAll(path, 0777)
	localFile, _ := os.OpenFile(path+fileName, os.O_CREATE|os.O_RDWR, 0777)
	if _, err := io.Copy(localFile, resp.Body); err != nil {
	} else {
		//fmt.Println("download file" + path + "/" + fileName + " success")
	}
}
