package main

import (
	"../../../../meitu"
	model "../../../model/meituri"
	"../../../util"
	"../gorm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var client *http.Client
var end = "cn"

func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second

	runtime.GOMAXPROCS(100)
	//wg.Add(1)
	//891 8245 8225
	//downloadModelColums([]int{8245}) //795,1289,954,3175,467,1558,429, 3239, 2008, 893,919
	getModelColums(786)
	//models := gorm.GetCNModels()
	//for _, model := range *models {
	//	fmt.Println(model.ID)
	//	getModelColums(model.ID)
	//}
	wg.Wait()
}

//func downloadColums(modelId int, columIds []int) {
//	for i := 0; i < len(columIds); i++ {
//		downloadSingleColum(modelId, columIds[i])
//	}
//}
func downloadModelColumsRange(from int, to int) {
	for i := from; i <= to; i++ {
		getModelColums(i)
		//getModelInfo(i)
	}
}
func downloadModelColums(modelIds []int) int {
	for i := 0; i < len(modelIds); i++ {
		err := getModelColums(modelIds[i])
		if err == -1 {
			continue
		} else if err == -2 {
			break
		}
	}
	return 0
}
func getModelColums(modelId int) int {
	for i := 1; i < meitu.UserColumsPageMaxSize; i++ {
		url := ""
		if i > 1 {
			url = meitu.Host + "/t/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(i) + ".html"
		} else {
			url = meitu.Host + "/t/" + strconv.Itoa(modelId) + "/"
		}
		err := AnalyzeModelHomePageHtml(client, url, modelId, i)
		if err == meitu.AnalysisHtmlFaild {
			continue
		} else if err == meitu.UrlInvalid {
			break
		}
	}
	return 0
}
func getModelInfo(modelId int) {
	url := meitu.Host + "/t/" + strconv.Itoa(modelId) + "/"
	AnalyzeModelInfoHtml(client, url, modelId)
}
func getCompanysColums(compId int) int {
	for i := 1; i < meitu.CompanysColumsPageMaxSize; i++ {
		url := ""
		if i > 1 {
			url = meitu.Host + "/x/" + strconv.Itoa(compId) + "/index_" + strconv.Itoa(i-1) + ".html"
		} else {
			url = meitu.Host + "/x/" + strconv.Itoa(compId) + "/"
		}
		err := AnalyzeCompanyHomePageHtml(client, url, compId, i)
		if err == meitu.AnalysisHtmlFaild {
			continue
		} else if err == meitu.UrlInvalid {
			break
		}
	}
	return meitu.Success
}

var wg sync.WaitGroup

func downloadSingleColum(modelId int, columId int, colum *model.Colums) int {
	downloadColumCover(modelId, columId)

	doc, err := goquery.NewDocument(meitu.Host + "/a/" + strconv.Itoa(columId))
	if err != nil {
		//return -1
		println(err.Error())
	}
	tuji := doc.Find("div.tuji")
	var no string
	var t string
	tuji.Find("p").Each(func(i int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "编号") {
			no = strings.Split(s.Text(), "：")[1]
		}
		if strings.Contains(s.Text(), "日期") {
			t = strings.Split(s.Text(), "：")[1]
		}
		//fmt.Println(s.Text())
	})
	if nil != tuji {
		//html, _ := tuji.Html()
		weizhi := tuji.Find("div.weizhi")
		if weizhi != nil {
			title := weizhi.Find("h1").Text()
			subs := doc.Find("div.shuoming").Find("p").Text()
			//println(title + subs)
			colum.ID = columId
			colum.Modelid = modelId
			colum.Title = title
			colum.Subs = subs
			colum.No = no
			colum.Time = t
			//colum.Html = html

			//c := model.Colums{
			//	ID:      columId,
			//	Modelid: modelId,
			//	Title:   title,
			//	Subs:    subs,
			//	Html:    html,
			//}
			//err := db.Create(c).Error
			//if err != nil {
			//	//return -3
			//	println(err.Error())
			//} else {
			//	println("saveColumInfo,Success" + strconv.Itoa(colum))
			//}
			gorm.SaveColumInfo(columId, colum)

			if meitu.DownloadImages {
				_ = os.MkdirAll("../meituri_"+end+"/"+strconv.Itoa(modelId)+"/"+strconv.Itoa(columId), os.ModePerm)
				if meitu.AsyTaskDownload {
					for i := 1; true; i++ {
						durl := meitu.OldHost + "/a/1/" + strconv.Itoa(columId) + "/" + strconv.Itoa(i) + ".jpg"
						//resp, err := url.ParseRequestURI(durl)
						//if err != nil {
						//	//panic("url err")
						//	println(err.Error())
						//	return -2
						//}
						filename := strconv.Itoa(i) + ".jpg"
						path := "../meituri_" + end + "/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(columId) + "/"
						//downloadFile(durl,path,filename)
						e, _ := util.PathExists(path + filename)
						if e {
							//fmt.Println("download file faild" + path + "/" + filename + "exist")
							continue
						}
						//filename := path.Base(uri.Path)
						req, err := http.NewRequest("GET", durl, nil)
						if err != nil {
							//fmt.Println("request create faild: " + err.Error())
							break
						}
						http.DefaultClient.Timeout = 20 * time.Second
						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							//fmt.Println("request error: " + err.Error())
							break
						}
						if resp.StatusCode != http.StatusOK {
							//fmt.Println("response status: " + strconv.Itoa(resp.StatusCode))
							break
						}
						wg.Add(1)
						go downloadImage(resp, path, filename)
					}
				} else {
					//for i := 1; i < 100; i++ {
					//	durl := meitu.OldHost + "/a/1/" + strconv.Itoa(columId) + "/" + strconv.Itoa(i) + ".jpg"
					//	uri, _ := url.ParseRequestURI(durl)
					//	//if err != nil {
					//	//	//panic("url err")
					//	//	println(err.Error())
					//	//	return -2
					//	//}
					//	filename := path.Base(uri.Path)
					//	path := "../meituri/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(columId) + "/"
					//	e, _ := util.PathExists(path + filename)
					//	if e {
					//		fmt.Println("file has exist" + path + filename)
					//		break
					//	}
					//	err := spider.DownloadFile(*client, durl, path, filename)
					//	if (err == -1) {
					//		continue
					//	} else if err == -2 {
					//		break
					//	}
					//}
				}
			}

			return meitu.Success
		} else {
			return meitu.AnalysisHtmlFaild
		}
	}

	return 0
}

func downloadColumCover(modelId int, columId int) {
	filename := "0.jpg"
	durl := meitu.OldHost + "/a/1/" + strconv.Itoa(columId) + "/" + filename
	//resp, err := url.ParseRequestURI(durl)
	//if err != nil {
	//	//panic("url err")
	//	println(err.Error())
	//	return -2
	//}
	path := "../meituri_" + end + "/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(columId) + "/"
	//downloadFile(durl,path,filename)
	e, _ := util.PathExists(path + filename)
	if e {
		//fmt.Println("download file faild" + path + "/" + filename + "exist")
		return
	}
	//filename := path.Base(uri.Path)
	req, err := http.NewRequest("GET", durl, nil)
	if err != nil {
		//fmt.Println("request create faild: " + err.Error())
		return
	}
	http.DefaultClient.Timeout = 20 * time.Second
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		//fmt.Println("request error: " + err.Error())
		return
	}
	if resp.StatusCode != http.StatusOK {
		//fmt.Println("response status: " + strconv.Itoa(resp.StatusCode))
		return
	}
	wg.Add(1)
	go downloadImage(resp, path, filename)
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

func AnalyzeModelHomePageHtml(client *http.Client, url string, modelId int, i int) int {
	resp, err := client.Get(url)
	fmt.Println(url)

	if err != nil {
		fmt.Println("analysis html faild")
		return meitu.AnalysisHtmlFaild
	}
	if resp.StatusCode > 400 {
		fmt.Println("resp.StatusCode > 400")
		return meitu.UrlInvalid
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocument(url)
	if nil == err {
		if i == 1 && meitu.SaveUserInfo {
			saveUseInfo(modelId, doc)
		}
	}
	return AnalyzeModelColumPage(modelId, doc)
}
func AnalyzeModelInfoHtml(client *http.Client, url string, modelId int) {
	resp, err := client.Get(url)
	fmt.Println(url)

	if err != nil {
		fmt.Println("analysis html faild")
	}
	if resp.StatusCode > 400 {
		fmt.Println("resp.StatusCode > 400")
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocument(url)
	if nil == err {
		saveUseInfo(modelId, doc)
	}
}
func saveUseInfo(modelId int, doc *goquery.Document) {
	cover, _ := doc.Find("div.left").Find("img").Attr("src")
	right := doc.Find("div.right")
	fmt.Println(right.Text())
	var map1 = make(map[string]string)
	right.Find("p").Each(func(i int, s *goquery.Selection) {
		html, _ := s.Html()
		text := strings.Replace(strings.Replace(html, "<span>", ";", -1), "</span>", "", -1)
		texts := strings.Split(text, ";")
		for _, t := range texts {
			len := len(strings.Trim(t, ""))
			//fmt.Println(t + "/" + strconv.Itoa(len))
			if len == 0 {

			} else {
				if strings.Contains(t, "：") {
					kv := strings.Split(t, "：")
					map1[kv[0]] = kv[1]
				} else if strings.Contains(t, ":") {
					kv := strings.Split(t, ":")
					map1[kv[0]] = kv[1]
				}
				//fmt.Println(kv[0])
			}
		}
		//fmt.Print(texts)
		//fmt.Println(strconv.Itoa(len(texts)))
		//texts:=strings.Split(text,"<span>")
	})

	nicknames := right.Find("h1").Text()
	shuoming := doc.Find("div.shuoming")
	more := shuoming.Text()
	tags := shuoming.Find("p").Text()
	fmt.Println(nicknames, modelId)
	model := model.Models{
		ID:            modelId,
		Cover:         cover,
		Name:          strings.Split(nicknames, "、")[0],
		Nicknames:     nicknames,
		More:          more,
		Tags:          tags,
		Birthday:      map1["生日"],
		Constellation: map1["星座"],
		Height:        map1["身高"],
		Weight:        map1["体重"],
		Dimensions:    map1["罩杯"],
		Address:       map1["来自"],
		Jobs:          map1["职业"],
		Interest:      map1["兴趣"],
	}
	gorm.SaveModelInfo(&model)
}

func AnalyzeCompanyHomePageHtml(client *http.Client, url string, companyId int, i int) int {
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		return meitu.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return meitu.AnalysisHtmlFaild
	}
	if i == 1 && meitu.SaveCompanyGroupRatation {
		doc.Find("div.fenlei").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
			name := s.Text()
			homepage, _ := s.Attr("href")
			id := util.GetIdFromUri(homepage)
			group := model.Groups{
				Id:       id,
				Name:     name,
				Homepage: homepage,
				Belong:   companyId,
			}
			gorm.SaveGroupInfo(group)
			//fmt.Print()
		})
	}
	return AnalyzeCompanyPage(doc)
}

func AnalyzeCompanyPage(doc *goquery.Document) int {
	//println(doc.Url.String())
	hezi := doc.Find("div.hezi")
	if hezi != nil {
		hezi.Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
			//path, _ := s.Find("a").Attr("href")
			//columId:=getIdFromUri(path)
			s.Find("p").Each(func(i int, s1 *goquery.Selection) {
				if i == 0 {
					modelhomepage, _ := s1.Find("a").Attr("href")
					//fmt.Println(modelhomepage)
					modelId := util.GetIdFromUri(modelhomepage)
					//downloadColum(modelId,columId)
					getModelColums(modelId)
				}
			})
		})
	} else {
		return meitu.AnalysisHtmlFaild
	}
	return meitu.Success
}

func AnalyzeModelColumPage(modelId int, doc *goquery.Document) int {
	//println(doc.Url.String())
	hezi := doc.Find("div.hezi")
	if hezi != nil {
		hezi.Find("ul").Find("li").Each(func(i int, s *goquery.Selection) {
			a := s.Find("a")
			path, _ := a.Attr("href")
			var groupId int
			var group string
			var tags string
			s.Find("p").Each(func(i int, s1 *goquery.Selection) {
				if strings.Contains(s1.Text(), "机构") {
					p1a := s1.Find("a")
					href, _ := p1a.Attr("href")
					groupId = util.GetIdFromUri(href)
					group = a.Text()
				} else if strings.Contains(s1.Text(), "类型") {

					s1.Find("a").Each(func(i int, s2 *goquery.Selection) {
						href, _ := s2.Attr("href")
						tags += s2.Text() + "(" + strconv.Itoa(util.GetIdFromUri(href)) + ")"
					})
				}
			})
			//cover, _ := a.Find("img").Attr("src")
			nums := s.Find("span").Text()
			num, _ := strconv.Atoi(nums[0 : len(nums)-1])

			paths := strings.Split(path, "/")
			columId, _ := strconv.Atoi(paths[len(paths)-2])
			//println(path + " " + strconv.Itoa(len(path)) + strconv.Itoa(columId))
			//if(!util.PathExists(strconv.Itoa(colum))){
			//
			//}
			c := model.Colums{
				Nums:    num,
				Modelid: modelId,
				Groupid: groupId,
				Group:   group,
				Tags:    tags,
			}
			//gorm.SaveColum(c)
			err := downloadSingleColum(modelId, columId, &c)
			if err < 0 {
				//continue
			}
		})
	} else {
		return meitu.AnalysisHtmlFaild
	}
	return meitu.Success
}
