package main

import (
	model "../../../model/meituba"
	"../../../util"
	"../orm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var host = "http://www.meituba.com/"

var wg sync.WaitGroup
var client *http.Client

func main() {
	orm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	jiexiPage(host, "xinggan")
	wg.Wait()
}

func jiexiPage(host string, fenlei string) int {
	url := host + fenlei + "/"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", "https://www.meituba.com")
	//req.Header.Add("Host", "img1.mmmw.net:443")
	//req.Header.Add("Proxy-Connection", "keep-alive")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	if err != nil {
		fmt.Println("request create faild: " + url + err.Error())
	}
	//http.DefaultClient.Timeout = 20 * time.Second;
	resp, err := client.Do(req)
	if err != nil {
		println(err.Error())
		return -1
	}
	//resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		fmt.Println(resp.StatusCode)
		return -1
	}
	if resp.StatusCode == 403 {
		panic(resp.StatusCode)
	}
	doc, err := goquery.NewDocument(url)
	doc.Find("div.channel_list").Find("ul").Find("li").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find("div.imgc").Find("a")
		vtimes := selection.Find("div.items_likes").Text()

		re := regexp.MustCompilePOSIX(`浏览次数：+([0-9]+)+次  共+([0-9]*)+张`)
		match := re.FindAllStringSubmatch(vtimes, -1)

		title, _ := a.Attr("title")
		println(title)

		href, _ := a.Attr("href")
		idstr := util.GetNameIDFromUri(href)
		id, _ := strconv.Atoi(idstr)
		println(title, href, idstr)
		img := a.Find("img")
		src, _ := img.Attr("src")
		wg.Add(1)
		go downloadImage(true, fenlei, idstr, src, "0.jpg")

		var times int
		var count int
		for _, m := range match {
			fmt.Println(m[1], m[2])
			times, err = strconv.Atoi(m[1])
			count, err = strconv.Atoi(m[2])
		}

		colum := model.Colums{
			ID:     id,
			Title:  title,
			Vtimes: times,
			Count:  count,
		}
		downloadColum(fenlei, idstr, count, &colum)
	})
	//doc.Find("div.pages").Find("ul").Find("li").Each(func(i int, selection *goquery.Selection) {
	//	a := selection.Find("a")
	//	if (strings.EqualFold(a.Text(), "下一页")) {
	//		href, e := a.Attr("href")
	//		if (e) {
	//			nextUrl := url + href
	//			jiexiPage(nextUrl)
	//		}
	//
	//	}
	//})
	return 0
}

var errtimes = 0

func downloadColum(fenlei string, id string, count int, colum *model.Colums) {
	if count == 0 {
		count = 100
	}
	for i := 1; i <= count; i++ {
		var url string
		if i == 1 {
			url = host + fenlei + "/" + id + ".html"
		} else {
			url = host + fenlei + "/" + id + "_" + strconv.Itoa(i) + ".html"
		}
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Referer", "https://www.meituba.com")
		//req.Header.Add("Host", "img1.mmmw.net:443")
		//req.Header.Add("Proxy-Connection", "keep-alive")
		//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
		if err != nil {
			fmt.Println("request create faild: " + err.Error())
			errtimes++
			if errtimes > 3 {
				break
			}
		}
		http.DefaultClient.Timeout = 20 * time.Second
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode >= 400 {
			fmt.Println(url+" request error: "+err.Error(), resp.StatusCode)
			errtimes++
			if errtimes > 3 {
				break
			}
		}

		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println(err.Error())
			break
		} else {
			fmt.Println(url)
		}
		photo := doc.Find("div.main").Find("div.photo")
		img := photo.Find("a").Find("img")
		src, e := img.Attr("src")
		if e {
			fmt.Println(src)
			errcode := downloadImage(false, fenlei, id, src, strconv.Itoa(i)+".jpg")
			if errcode == -1 {
				continue
			} else if errcode == -2 {
				break
			}
		}
		var end bool
		doc.Find("div.pages").Find("ul").Find("li").Each(func(i int, selection *goquery.Selection) {
			a := selection.Find("a")
			fmt.Println(a.Text())
			if strings.EqualFold(a.Text(), "下一页") {
				href, e := a.Attr("href")
				if e {
					fmt.Println(href, len(href))
					end = len(href) <= 1
				}
			}
		})
		if i == 1 {
			doc.Find("div.photo-fb1").Find("a.blue").Each(func(i int, selection *goquery.Selection) {
				cname := selection.Text()
				var ename string
				href, e := selection.Attr("href")
				if e {
					ename = util.GetNameIDFromUri(href)
				}
				tag := model.Tags{
					Cname: cname,
					Ename: ename,
				}
				orm.SaveTag(&tag)
			})
			orm.SaveColum(colum, "colums_"+fenlei)
		}

		if end {
			break
		}
	}
}

// 下载图片
func downloadImage(asy bool, fenlei, columId, url, filename string) int {

	path := "../meituba/" + fenlei + "/" + columId + "/"
	//downloadFile(durl,path,filename)
	e, err := util.PathExists(path + filename)
	if e {
		fmt.Println("file exist")
		return -1
	}
	//filename := path.Base(uri.Path)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", "https://www.meituba.com")
	//req.Header.Add("Host", "img1.mmmw.net:443")
	//req.Header.Add("Proxy-Connection", "keep-alive")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	if err != nil {
		fmt.Println("request create faild: " + err.Error())
		return -1
	}
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(url + " request error: " + err.Error())
		return -1
	}
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println(url + " response status: " + strconv.Itoa(resp.StatusCode))
	//	return -1
	//}
	//fileName := getNameFromUrl(url)
	defer func() {
		resp.Body.Close()
		if r := recover(); r != nil {
			//fmt.Println(r)
		}
		if asy {
			wg.Done()
		}
	}()
	_ = os.MkdirAll(path, 0777)
	localFile, _ := os.OpenFile(path+filename, os.O_CREATE|os.O_RDWR, 0777)
	if _, err := io.Copy(localFile, resp.Body); err != nil {
	} else {
		//fmt.Println("download file" + path + "/" + fileName + " success")
	}
	return 0
}
