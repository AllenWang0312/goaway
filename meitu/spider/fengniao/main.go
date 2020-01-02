package main

import (
	"../../../conf"
	"../../../util"
	model "../../model/fengniao"
	"./gorm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var client *http.Client

func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second

	runtime.GOMAXPROCS(100)
	//var url = "http://m.fengniao.com/travel/"
	//AnalyzeColumHomePageHtml(client, url)
	//AnalyzeFromM(5359571)
	//downloadModelColumsRange(10000,20000)
	var url = "http://travel.fengniao.com/slide/535/5359673_1.html"
	var dir = util.GetNameIDFromUri(url)
	AnalyzeFrom(url, dir)
	WG.Wait()
}
func AnalyzeFrom(url string, dir string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("request create faild: " + err.Error())
		return -1
	}
	//addHeaders(req,false)
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		println("request create faild: " + strconv.Itoa(resp.StatusCode))
		return conf.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return conf.AnalysisHtmlFaild
	}

	docs, err := ioutil.ReadAll(resp.Body)
	var titleDoc = doc.Find("div.title")
	var title = titleDoc.Find("h4").Text()
	var totalStr = titleDoc.Find("span.mark-num").Find("span.total-num").Text()
	total, _ := strconv.Atoi(totalStr)
	if err != nil {
		println(err.Error())
	} else {
		docstr := util.G2U(string(docs))
		println(docstr)
		r, _ := regexp.Compile("var picInfoJson = '(.+)'")
		var jsons = r.FindString(docstr)
		fmt.Println(jsons)
		r, _ = regexp.Compile("pic_url\":\"(.+?)\"")
		var urls = r.FindAllStringSubmatch(jsons, -1)
		println(len(urls))



		var desc = doc.Find("p.describe-text").Text()

		id, _ := strconv.Atoi(dir)
		//var docDoc=util.G2U(doc.Text())
		println(util.G2U(title), totalStr)

		var album = model.Album{
			Title: title,
			Subs:  desc,
			Nums:  total,
			ID:    id,
		}

		gorm.SaveAlbum(&album)
		if(conf.DownloadImages){
			for i, kv := range urls {
				downloadImage(strings.Replace(kv[1], "\\", "", -1), dir, strconv.Itoa(i+1)+".jpg")
			}
		}
		next, e := doc.Find("div.next-last").Find("a").Attr("href")
		if (e) {
			var dir = util.GetNameIDFromUri(next)
			AnalyzeFrom(next, dir)
		}
		//fmt.Println(urls)
		return 0
	}

	if (err == nil) {
		var bigImgBox = doc.Find("div.big-img-box")
		var picbox = bigImgBox.Find("div.pic-box")
		var box = picbox.Text()
		println(box)
		var img = picbox.Find("img")
		var orc_url, _ = img.Attr("orc_url")
		//println(img.Text(),orc_url)
		downloadImage(orc_url, dir, "1.jpg")
		for i := 2; i <= total; i++ {
			AnalyzeImage(url+"#p="+strconv.Itoa(i), dir, strconv.Itoa(i)+".jpg")
		}
	}
	next, e := doc.Find("div.next-last").Find("a").Attr("href")
	if (e) {
		var dir = util.GetNameIDFromUri(next)
		AnalyzeFrom(next, dir)
	}
	return 0
}
func AnalyzeImage(url string, dir string, saveName string) int {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("request create faild: " + err.Error())
		return -1
	}
	addHeaders(req, false)
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		println("request create faild: " + strconv.Itoa(resp.StatusCode))
		return conf.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	var img = doc.Find("div.pic-box").Find("img")
	var orc_url, _ = img.Attr("orc_url")
	println(img.Text(), orc_url)
	downloadImage(orc_url, dir, saveName)
	return 0
}

func downloadImage(orc_url string, dir string, filename string) {
	println(orc_url)
	//var route = util.GetRouteFromUri(orc_url)
	//var name = util.GetNameFromUri(orc_url)
	DownloadImage(orc_url, conf.FSRoot+"/fengniao/"+dir+"/", filename)
}

//func AnalyzeFromM(center int) {
//	getAlbumCount(center)
//	//for i := 0;; i++ {
//	//	AnalyzeAlbumHtml(center-i)
//	//	AnalyzeAlbumHtml(center+i)
//	//	//getModelInfo(i)
//	//}
//}
//func downloadModelColumsRange(from int, to int) {
//	for i := from; i <= to; i++ {
//		AnalyzeAlbumHtml(i,dir)
//		//getModelInfo(i)
//	}
//}

var count = 10

func getAlbumCount(albumId int, dir string) int {
	var url = "https://m.fengniao.com/slide/" + strconv.Itoa(albumId) + ".html"
	println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("request create faild: " + err.Error())
		return -1
	}
	addHeaders(req, true)
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	//resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		println("request create faild: " + strconv.Itoa(resp.StatusCode))
		return conf.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return conf.AnalysisHtmlFaild
	}
	//var title = doc.Find("div.cont92").Find("h3").Text()
	var totalStr = doc.Find("div.swiper-pagination").Find("span.swiper-pagination-total").Text()
	count, _ = strconv.Atoi(totalStr)
	for i := 1; i <= count; i++ {
		AnalyzeAlbumHtml(albumId, dir)
	}
	return 0
}

func addHeaders(req *http.Request, m bool) {
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Cookie", "ip_ck=58KI7/v3j7QuMDU5ODEwLjE1NzYwNTA0NDI%3D; Hm_lvt_916ddc034db3aa7261c5d56a3001e7c5=1576050390,1576131082,1576196777,1577084602; fn_forum_id=15; Hm_lvt_29a81ae42ab828b4819bb3d3a871d9ef=1577092807; _csrf-frontend-m=7db9474df01fdade23c6e977503c47319b8e8dd1566073647b41fab81d2ef530a%3A2%3A%7Bi%3A0%3Bs%3A16%3A%22_csrf-frontend-m%22%3Bi%3A1%3Bs%3A32%3A%22FwBwPosGB72-LwzlEp7XcxAmUMOdy4zn%22%3B%7D; lv=1577154407; vn=7; Adshow=5; Hm_lpvt_29a81ae42ab828b4819bb3d3a871d9ef=1577154534; Hm_lpvt_916ddc034db3aa7261c5d56a3001e7c5=1577154686; fn_document_class_doc_id=340-5359624%2C278-5359604%2C278-5359544%2C278-5359571%2C278-5359594%2C278-5359459")
	req.Header.Add("Host", "m.fengniao.com")
	if m {
		req.Header.Add("Referer", "https://m.fengniao.com/slide/5359604.html")
	} else {
		req.Header.Add("Referer", "http://travel.fengniao.com/slide/535/5359628_1.html")
	}
	//

	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
}

func AnalyzeAlbumHtml(albumId int, dir string) int {
	var url = "https://m.fengniao.com/slide/" + strconv.Itoa(albumId) + ".html"
	println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("request create faild: " + err.Error())
		return -1
	}
	addHeaders(req, true)
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	//resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		println("request create faild: " + strconv.Itoa(resp.StatusCode))
		return conf.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return conf.AnalysisHtmlFaild
	}
	var title = doc.Find("div.cont92").Find("h3").Text()
	var totalStr = doc.Find("div.swiper-pagination").Find("span.swiper-pagination-total").Text()
	count, _ = strconv.Atoi(totalStr)
	var img, e = doc.Find("div.swiper-slide").Find("img").Attr("src")
	if (e) {
		println(util.G2U(title))
		img = util.GetPathFromUri(img)
		var name = util.GetNameFromUri(img)
		println(img, dir, name)
		DownloadImage(img, conf.FSRoot+"/fengniao/"+dir+"/", name)
	}
	return 0
}

func AnalyzeColumHomePageHtml(url string) int {
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		return conf.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return conf.AnalysisHtmlFaild
	}
	//var menu=doc.Find("div.menuBox")
	var listBox = doc.Find("div.listBox")

	var divs = listBox.Find("div.li")

	println(divs.Length())
	divs.Each(func(i int, s *goquery.Selection) {
		var a = s.Find("a.pic")
		href, e := a.Attr("href")
		if (e) {
			println(href)
		}
		//title, e := a.Attr("title")
	})
	return 0
}

var WG sync.WaitGroup
// 下载图片
func DownloadImage(durl string, path string, fileName string) int {
	//downloadFile(durl,path,filename)
	//println(durl,path,fileName)
	e, _ := util.PathExists(path + fileName)
	if e {
		fmt.Println("download file faild" + path + "/" + fileName + "exist")
		return -1
	}
	//filename := path.Base(uri.Path)
	req, err := http.NewRequest("GET", durl, nil)
	if err != nil {
		fmt.Println("request create faild: " + err.Error())
		return -2
	}
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("request error: " + err.Error())
		return -2
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("response status: " + strconv.Itoa(resp.StatusCode))
		return -2
	}
	WG.Add(1)
	go DownloadImageFromResp(resp, path, fileName)
	return 0
}

func DownloadImageFromResp(resp *http.Response, path string, fileName string) {
	//fileName := getNameFromUrl(url)
	defer func() {
		resp.Body.Close()
		//if r := recover(); r != nil {
		//fmt.Println(r)
		//}
		WG.Done()
	}()

	_ = os.MkdirAll(path, 0777)
	localFile, _ := os.OpenFile(path+fileName, os.O_CREATE|os.O_RDWR, 0777)
	if _, err := io.Copy(localFile, resp.Body); err != nil {
	} else {
		//fmt.Println("download file" + path + "/" + fileName + " success")
	}
}
