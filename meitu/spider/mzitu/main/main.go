package main

import (
	model "../../../model/mzitu"
	"../../../util"
	"../gorm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"io"
	"github.com/axgle/mahonia"
)

var host = "http://www.mzitu.com/"

var client *http.Client
var wg sync.WaitGroup

func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	//http.DefaultClient.Timeout = 20 * time.Second;
	//getColums(0, 80000)
	//getColums(0, 28067) //193336
	getColums( 28068,100000) //193336
	//getColums(1, 100) //193336
	//downloadColum(77923)
	//downloadColum(79058)
	//downloadColum(78993)
	wg.Wait()
}
func getColums(from int, to int) {
	for i := from; i <= to; i++ {
		//exit, err := util.PathExists("../mzitu/" + strconv.Itoa(i))
		//if !(nil == err && exit) {
		//	continue
		//}
		downloadColum(i)
	}
}

var errcount = 0
var saveInfo = false

func downloadColum(columId int) {
	for j := 1; ; j++ {
		var h5_url string
		if (j == 1) {
			h5_url = host + strconv.Itoa(columId)
		} else {
			h5_url = host + strconv.Itoa(columId) + "/" + strconv.Itoa(j)
		}
		println(h5_url)
		req, _ := http.NewRequest("Get", h5_url, nil)
		resp, err := http.DefaultClient.Do(req)
		if nil != err {
			HandleError(err, "http.NewRequest(Get,h5_url,nil)")
			break
		}
		if resp.StatusCode >= 400 {
			println(resp.StatusCode)
			errcount++
			if (j == 1 || errcount >= 3) {
				break
			} else {
				continue
			}
		} else {
			time.Sleep(time.Second)
		}
		dec := mahonia.NewDecoder("UTF-8")
		rd := dec.NewReader(resp.Body)
		doc, err := goquery.NewDocumentFromReader(rd)

		if err != nil {
			HandleError(err, "goquery.NewDocument(h5_url)")
			errcount++
			if (j == 1 || errcount >= 3) {
				break
			} else {
				continue
			}
		} else {
			content := doc.Find("div.content")
			if (j == 1 && saveInfo) {
				title := content.Find("h2.main-title").Text()
				var time string
				content.Find("div.main-meta").Find("span").Each(func(i int, selection *goquery.Selection) {
					if (i == 1) {
						timetag := selection.Text()
						time = string(timetag[10:len(timetag)])
						//fmt.Println(time)
					}
				})

				content.Find("div.main-tags").Find("a").Each(func(i int, selection *goquery.Selection) {
					tag_url, _ := selection.Attr("href")
					url1 := tag_url[0 : len(tag_url)-1]
					//fmt.Println(url1)
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
			println(durl)
			filename := util.GetNameFromUri(durl)
			if (len(filename) == 9|| len(filename)==10) {
				downloadImages(durl[0:len(durl)-6], columId)
				break
			} else {
				filename = strconv.Itoa(j)+".jpg"
			}
			path := "../mzitu/" + strconv.Itoa(columId) + "/"
			e, _ := util.PathExists(path + filename)

			if e && util.GetFileSize(path+filename) > 0 {
				fmt.Println("download file faild" + path + filename + "exist")
				continue
			}
			wg.Add(1)
			go downloadImage(durl, path, filename, columId, j)
			//count, _ := strconv.Atoi(content.Find("div.pagenavi").Find("span").Text())
			//if (j >= count) {
			//	break
			//}
		}

	}

}

func downloadImages(url string, columId int) {
	for i := 1; i < 100; i++ {
		var durl string
		if (i < 10) {
			durl = url + "0" + strconv.Itoa(i) + ".jpg"
		} else {
			durl = url + strconv.Itoa(i) + ".jpg"
		}
		filename := strconv.Itoa(i) + ".jpg"
		path := "../mzitu/" + strconv.Itoa(columId) + "/"
		e, _ := util.PathExists(path + filename)
		if e && util.GetFileSize(path+filename) > 0 {
			fmt.Println("download file faild" + path + filename + "exist")
			continue
		}
		wg.Add(1)
		go downloadImage(durl, path, filename, columId, i)
	}
}

func downloadImage(durl string, path string, fileName string, columId int, j int) int {
	req, err := http.NewRequest("GET", durl, nil)
	req.Header.Add("Referer", "https://www.mzitu.com/"+strconv.Itoa(columId)+"/"+strconv.Itoa(j))
	//req.Header.Add("Host", "img1.mmmw.net:443")
	//req.Header.Add("Proxy-Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	if err != nil {
		//fmt.Println("request create faild: " + durl + err.Error())
		return -1
	}
	http.DefaultClient.Timeout = 20 * time.Second;
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//fmt.Println("request error: " + durl + err.Error())
		return -1
	}
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("response status: " + durl + strconv.Itoa(resp.StatusCode))
		return -1
	}
	//fileName := getNameFromUrl(url)
	defer func() {
		resp.Body.Close()
		if r := recover(); r != nil {
			//fmt.Println(r)
		}
		wg.Done()
	}()
	_ = os.MkdirAll(path, 0777)

	//fBytes, e := ioutil.ReadAll(resp.Body)
	//HandleError(e, "ioutil.ReadAll(resp.Body)")
	//err = ioutil.WriteFile(path+fileName, fBytes, 0644)

	localFile, _ := os.OpenFile(path+fileName, os.O_CREATE|os.O_RDWR, 0777)
	if _, err := io.Copy(localFile, resp.Body); err != nil {
	} else {
		//fmt.Println("download file" + path + "/" + fileName + " success")
	}
	return 0
}
func HandleError(err error, tag string) {
	fmt.Println(tag + ":" + err.Error())
}
