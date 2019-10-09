package main

import (
	model "../../../model/doubanmovie"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"

	"../../../util"
	"../gorm"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//:authority: img3.doubanio.com
//:method: GET
//:path: /view/photo/raw/public/p2568276552.jpg
//:scheme: https
//accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3
//accept-encoding: gzip, deflate, br
//accept-language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7
//cache-control: max-age=0
//if-modified-since: Wed, 21 Jan 2004 19:51:30 GMT
//referer: https://movie.douban.com/photos/photo/2568276552/
//upgrade-insecure-requests: 1
//user-agent: Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36

var cookie = "bid=loH9qUTDm5g; __utmc=30149280; ll=\"118172\"; _vwo_uuid_v2=D47CB14E612F54A91B83B25A889C557D0|abc03ad33baf660166a9cb47973185c0; __utmv=30149280.11883; UM_distinctid=16c4c17da9ce9-07177a257b86cc-5f1d3a17-1fa400-16c4c17da9dfd; __utmz=30149280.1568622700.10.10.utmcsr=pianyuan.la|utmccn=(referral)|utmcmd=referral|utmcct=/r_ZZWkD13v0.html; __utmc=223695111; __utmz=223695111.1568622700.1.1.utmcsr=pianyuan.la|utmccn=(referral)|utmcmd=referral|utmcct=/r_ZZWkD13v0.html; dbcl2=\"118832098:g+RgBllfK9k\"; ck=otKZ; push_noty_num=0; push_doumail_num=0; __utma=30149280.1766030497.1548811165.1568710944.1568767473.13; __utma=223695111.453227035.1568622700.1568710944.1568767473.4; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1568852709%2C%22http%3A%2F%2Fpianyuan.la%2Fr_ZZWkD13v0.html%22%5D; ap_v=0,6.0; __yadk_uid=nhCM3fIzyImHQ60xi0Tyc6PRK8VygLbR; _pk_id.100001.4cf6=cd38a9b488d688a9.1567062657.6.1568853750.1568767473."

var posters_url = "https://movie.douban.com/subject/26709258/photos?type=R&start=0&sortby=like&size=a&subtype=a"
var poster_detail_url = "https://movie.douban.com/photos/photo/2568288336/"
var poster_hd = "https://img3.doubanio.com/view/photo/raw/public/p2568288336.jpg"

var client *http.Client

func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second
	//getMovieInfo(10001,100000)
	offset := 1
	wg.Add(1)
	ticker1 := time.NewTicker(1 * time.Second)
	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			paquMovie(26709806 + offset)
			//paquMovie(26709197 - offset)
			offset++
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)
	//paquMovie(26709258)
	wg.Wait()
}

var wg sync.WaitGroup

func getMovieInfo(from int, to int) {
	for i := from; i <= to; i++ {
		paquMovie(26709258)
	}
}
func paquMovie(movieId int) int {
	var url = "https://movie.douban.com/subject/" + strconv.Itoa(movieId)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Referer", "https://movie.douban.com")
	//req.Header.Add("Host", "img1.mmmw.net:443")
	//req.Header.Add("Proxy-Connection", "keep-alive")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	if err != nil {
		fmt.Println("request create faild: " + url + err.Error())
	}
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
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
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	subjectwrap := doc.Find("div.subjectwrap")
	subject := subjectwrap.Find("div.subject")

	durl, _ := subject.Find("div#mainpic").Find("a.nbgnbg").Find("img").Attr("src")
	fmt.Println(durl)
	filename := util.GetNameFromUri(durl)
	fmt.Println(filename)
	path := "../cover/"
	fmt.Println(path)
	downloadfile(durl, path, filename)

	info := subject.Find("div#info")
	//fmt.Println(info.Text())
	var drcs = make([]model.Celebrity, 0)
	var sws = make([]model.Celebrity, 0)
	var acts = make([]model.Celebrity, 0)
	var link string
	spans := info.Find("span")
	fmt.Println(spans.Size())
	spans.Each(func(i int, selection *goquery.Selection) {
		fmt.Println(strconv.Itoa(i) + " " + selection.Text())
		if i == 0 || i == 3 || i == 6 {
			selection.Find("span.attrs").Find("a").Each(func(j int, selection2 *goquery.Selection) {
				url, _ := selection2.Attr("href")
				id := util.GetIdFromUri(url)
				act := model.Celebrity{
					ID:   id,
					Name: selection2.Text(),
				}
				if i == 0 {
					drcs = append(drcs, act)
				} else if i == 3 {
					sws = append(sws, act)
				} else if i == 6 {
					acts = append(acts, act)
				}
			})
		}
		if strings.EqualFold(selection.Text(), "IMDb链接:") {
			link, _ = selection.Find("a").Attr("href")
			fmt.Println(link)
		}
	})
	infos := strings.Split(info.Text(), "\n")
	//fmt.Println(strings.Contains(info.Text(), "\n"))

	director_json, _ := json.Marshal(&drcs)
	sw_json, _ := json.Marshal(&sws)
	acts_json, _ := json.Marshal(&acts)

	var map1 = make(map[string]string)
	for i := range infos {
		t := strings.Trim(infos[i], " ")
		//fmt.Println(t)
		if strings.Contains(t, ": ") {
			kv := strings.Split(t, ": ")
			map1[kv[0]] = kv[1]
		}
	}
	title := strings.Trim(doc.Find("div#content").Find("h1").Find("span").Text(), " ")
	score, _ := strconv.ParseFloat(doc.Find("strong.ll").Text(), 32)
	movie := model.Subject{
		ID:           movieId,
		Name:         title,
		Cover:        filename,
		Director:     string(director_json),
		Screenwriter: string(sw_json),
		Actor:        string(acts_json),
		Website:      map1["官方网站"],
		Genre:        map1["类型"],
		Space:        map1["制片国家/地区"],
		Language:     map1["语言"],
		Time:         map1["上映日期"],
		Duration:     map1["片长"],
		Nickname:     map1["又名"],
		Imdblink:     map1["IMDb链接"],
		Score:        score,
		Sutime:       time.Now().Format("2006-01-02"),
	}
	gorm.SaveMovie(&movie)
	//movie_json, _ := json.Marshal(&movie)
	//println(string(movie_json))
	return 0
}

func downloadfile(durl string, path string, filename string) {
	//downloadFile(durl,path,filename)
	e, _ := util.PathExists(path + filename)
	if e {
		fmt.Println("download file faild" + path + "/" + filename + "exist")
		return
	}
	//filename := path.Base(uri.Path)
	req, err := http.NewRequest("GET", durl, nil)
	req.Header.Add("Referer", "https://movie.douban.com")
	//req.Header.Add("Host", "img1.mmmw.net:443")
	//req.Header.Add("Proxy-Connection", "keep-alive")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5000) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Mobile Safari/537.36")
	if err != nil {
		fmt.Println("request create faild: " + durl + err.Error())
	}
	http.DefaultClient.Timeout = 20 * time.Second

	resp2, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("request error: " + durl + err.Error())
	}
	if resp2.StatusCode != http.StatusOK {
		fmt.Println("response status: " + durl + strconv.Itoa(resp2.StatusCode))
	}

	wg.Add(1)
	go downloadImage(resp2, path, filename)
}

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
