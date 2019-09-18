package main

import (
	model "../../../model/mzitu"
	"../../../util"
	"../gorm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var host = "http://www.mzitu.com/"

var client *http.Client
var wg sync.WaitGroup

func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	getColums(4001, 10000)//193336
	//downloadColum(192173)
	wg.Wait()
}
func getColums(from int, to int) {
	for i := from; i <= to; i++ {
		downloadColum(i)
	}
}

func downloadColum(columId int) {
	for j := 1; ; j++ {
		var h5_url string
		if (j == 1) {
			h5_url = host + strconv.Itoa(columId)
		} else {
			h5_url = host + strconv.Itoa(columId) + "/" + strconv.Itoa(j)
		}
		resp, err := client.Get(h5_url)
		defer resp.Body.Close()
		if resp.StatusCode > 400 {
			break
		}
		doc, err := goquery.NewDocument(h5_url)
		if err != nil {
			break
		}
		content := doc.Find("div.content")
		if (j == 1) {
			title := content.Find("h2.main-title").Text()
			var time string
			content.Find("div.main-meta").Find("span").Each(func(i int, selection *goquery.Selection) {
				if (i == 1) {
					timetag := selection.Text()
					time = string(timetag[10:len(timetag)])
					fmt.Println(time)
				}
			})
			content.Find("div.main-tags").Find("a").Each(func(i int, selection *goquery.Selection) {
				tag_url, _ := selection.Attr("href")
				url1 := tag_url[0 : len(tag_url)-1]
				fmt.Println(url1)
				ename := string(url1[strings.LastIndex(url1, "/")+1 : len(url1)])
				cname := selection.Text()
				tag := model.Tags{
					Ename: ename,
					Cname: cname,
				}
				id := gorm.SaveTags(ename, &tag)
				relation := model.Columtags{
					Columid: columId,
					Tagid:   id,
					Lock:    strconv.Itoa(columId) + "_" + strconv.Itoa(id),
				}
				gorm.SaveTagRelation(&relation)
			})
			colum := model.Colums{
				ID:    columId,
				Title: title,
				Time:  time,
			}
			gorm.SaveColum(&colum)
		}
		//break//只更新表结构不下载
		durl, _ := content.Find("div.main-image").Find("p").Find("a").Find("img").Attr("src")

		//downloadImage(h5_url, columId)
		//resp, _ := url.ParseRequestURI(durl)
		//if err != nil {
		//	//panic("h5_url err")
		//	println(err.Error())
		//	return -2
		//}
		filename := util.GetNameFromUri(durl)
		path := "../mzitu/" + strconv.Itoa(columId) + "/"
		//downloadFile(durl,path,filename)
		e, _ := util.PathExists(path + filename)
		if e {
			fmt.Println("download file faild" + path + "/" + filename + "exist")
			continue
		}
		//filename := path.Base(uri.Path)
		req, err := http.NewRequest("GET", durl, nil)
		req.Header.Add("Referer", "https://www.mzitu.com")
		//req.Header.Add("Host", "img1.mmmw.net:443")
		//req.Header.Add("Proxy-Connection", "keep-alive")
		//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
		if err != nil {
			fmt.Println("request create faild: " + durl + err.Error())
			break
		}
		http.DefaultClient.Timeout = 20 * time.Second;
		resp2, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("request error: " + durl + err.Error())
			break
		}
		if resp2.StatusCode != http.StatusOK {
			fmt.Println("response status: " + durl + strconv.Itoa(resp2.StatusCode))
			break
		}
		wg.Add(1)
		go downloadImage(resp2, path, filename)
	}

}

//func downloadImage(url string, columId int) {
//	for i := 1; true; i++ {
//		var durl string
//		if (i < 10) {
//			durl = url + "0" + strconv.Itoa(i) + ".jpg"
//		} else {
//			durl = url + strconv.Itoa(i) + ".jpg"
//		}
//
//		//resp, err := url.ParseRequestURI(durl)
//		//if err != nil {
//		//	//panic("url err")
//		//	println(err.Error())
//		//	return -2
//		//}
//		filename := strconv.Itoa(i) + ".jpg"
//		path := "../mzitu/" + strconv.Itoa(columId) + "/"
//		//downloadFile(durl,path,filename)
//		e, _ := util.PathExists(path + filename)
//		if e {
//			fmt.Println("download file faild" + path + "/" + filename + "exist")
//			continue
//		}
//		//filename := path.Base(uri.Path)
//		req, err := http.NewRequest("GET", durl, nil)
//		req.Header.Add("Referer", "https://www.mzitu.com")
//		//req.Header.Add("Host", "img1.mmmw.net:443")
//		//req.Header.Add("Proxy-Connection", "keep-alive")
//		//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
//		if err != nil {
//			fmt.Println("request create faild: " + durl + err.Error())
//			break
//		}
//		http.DefaultClient.Timeout = 20 * time.Second;
//		resp, err := http.DefaultClient.Do(req)
//		if err != nil {
//			fmt.Println("request error: " + durl + err.Error())
//			break
//		}
//		if resp.StatusCode != http.StatusOK {
//			fmt.Println("response status: " + durl + strconv.Itoa(resp.StatusCode))
//			break
//		}
//		wg.Add(1)
//		go downloadImage(resp, path, filename)
//	}
//}

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
