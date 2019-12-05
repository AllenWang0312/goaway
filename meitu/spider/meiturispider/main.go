package main

import (
	"../../../conf"
	"../../../util"
	model "../../model/meituri"
	"./download"
	"./gorm"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var client *http.Client



func main() {
	gorm.InitDB()
	client = http.DefaultClient
	client.Timeout = 20 * time.Second

	runtime.GOMAXPROCS(100)
	//WG.Add(1)
	//891 8245 8225
	//downloadModelColums([]int{8245}) //795,1289,954,3175,467,1558,429, 3239, 2008, 893,919
	//models := gorm.GetCNModels()
	//for _, model := range *models {
	//	fmt.Println(model.ID)
	//	changdir(model.ID)
	//}
	if len(os.Args) > 1 {
		if len(os.Args) == 2 { //main cover pageNo pageSize

		} else if len(os.Args) == 3 {
			id1, err := strconv.Atoi(os.Args[1])
			id2, err1 := strconv.Atoi(os.Args[2])
			if err == nil { //main pageNo pageSize
				gorm.CreateHistryForAlbum(id1, id2)
			} else { //main jp 1234
				if err1 == nil {
					download.END = os.Args[1]
					getModelColums(id2)
				}
			}
		} else if len(os.Args) == 4 {
			str1 := os.Args[1]
			id1, err := strconv.Atoi(os.Args[2])
			id2, err1 := strconv.Atoi(os.Args[3])
			if err == nil && err1 == nil {
				if strings.EqualFold(str1, "cover") {//4281
					gorm.DownloadCoverForModel(id1, id2)
				}else if strings.EqualFold(str1, "range") { //main range 1 1000 //step 1
					downloadModelColumsRange(id1, id2) //下载 from to
				} else if strings.EqualFold(str1, "hot") {
					gorm.UpDateHot(id1, id2)
				} else if strings.EqualFold(str1, "分类") {
					gorm.CreateTableForModels(id1, id2) //main 模特分类/step2
				} else { //main cn modelid albunid
					download.END = str1
					downloadSingleColum(id1, id2, nil) //下载 model/album
				}
			}
		} else if len(os.Args) == 5 { //main splash modelid columid 11.jpeg
			modelid, err := strconv.Atoi(os.Args[2])
			albumid, err1 := strconv.Atoi(os.Args[3])
			resName := os.Args[4]
			if err == nil && err1 == nil {
				gorm.CreateSplashForColum(modelid, albumid, resName)
			}
		}
	} else {

	}
	download.WG.Wait()
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
	for i := 1; i < conf.UserColumsPageMaxSize; i++ {
		url := ""
		if i > 1 {
			url = conf.Host + "/t/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(i) + ".html"
		} else {
			url = conf.Host + "/t/" + strconv.Itoa(modelId) + "/"
		}
		err := AnalyzeModelHomePageHtml(client, url, modelId, i)
		if err == conf.AnalysisHtmlFaild {
			continue
		} else if err == conf.UrlInvalid {
			break
		}
	}
	return 0
}

func getModelInfo(modelId int) {
	url := conf.Host + "/t/" + strconv.Itoa(modelId) + "/"
	AnalyzeModelInfoHtml(client, url, modelId)
}

func getCompanysColums(compId int) int {
	for i := 1; i < conf.CompanysColumsPageMaxSize; i++ {
		url := ""
		if i > 1 {
			url = conf.Host + "/x/" + strconv.Itoa(compId) + "/index_" + strconv.Itoa(i-1) + ".html"
		} else {
			url = conf.Host + "/x/" + strconv.Itoa(compId) + "/"
		}
		err := AnalyzeCompanyHomePageHtml(client, url, compId, i)
		if err == conf.AnalysisHtmlFaild {
			continue
		} else if err == conf.UrlInvalid {
			break
		}
	}
	return conf.Success
}


func downloadSingleColum(modelId int, columId int, colum *model.Album) int {
	download.DownloadAlbumCover(modelId, columId)
	doc, err := goquery.NewDocument(conf.Host + "/a/" + strconv.Itoa(columId))
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
			if nil != colum {
				colum.ID = columId
				colum.Modelid = modelId
				colum.Title = title
				colum.Subs = subs
				colum.No = no
				colum.Time = t
				gorm.SaveColumInfo(columId, colum)
			}
			//colum.Html = html
			//c := model.Album{
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
			if conf.DownloadImages {
				_ = os.MkdirAll(conf.FSRoot+"/meituri_"+download.END+"/"+strconv.Itoa(modelId)+"/"+strconv.Itoa(columId), os.ModePerm)
				if conf.AsyTaskDownload {
					for i := 1; true; i++ {
						durl := conf.OldHost + "/a/1/" + strconv.Itoa(columId) + "/" + strconv.Itoa(i) + ".jpg"
						//resp, err := url.ParseRequestURI(durl)
						//if err != nil {
						//	//panic("url err")
						//	println(err.Error())
						//	return -2
						//}
						filename := strconv.Itoa(i) + ".jpg"
						path := conf.FSRoot + "/meituri_" + download.END + "/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(columId) + "/"

						download.WG.Add(1)
						go download.DownloadImage(durl, path, filename)
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
					//	path := conf.FSRoot+"/meituri/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(columId) + "/"
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

			return conf.Success
		} else {
			return conf.AnalysisHtmlFaild
		}
	}

	return 0
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
		return conf.AnalysisHtmlFaild
	}
	if resp.StatusCode > 400 {
		fmt.Println("resp.StatusCode > 400")
		return conf.UrlInvalid
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocument(url)
	if nil == err {
		if i == 1 && conf.SaveUserInfo {
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
	var from = map1["来自"]
	var contry = ""
	model := model.Model{
		Cover:         cover,
		Nicknames:     nicknames,
		More:          more,
		Tags:          tags,
		Birthday:      map1["生日"],
		Constellation: map1["星座"],
		Height:        map1["身高"],
		Weight:        map1["体重"],
		Dimensions:    map1["罩杯"],
		Address:       from,
		Jobs:          map1["职业"],
		Interest:      map1["兴趣"],
	}
	model.ID = modelId
	model.Name = strings.Split(nicknames, "、")[0]
	//if (strings.Contains(from, "日本")) {
	//	contry="jp"
	//}else if strings.Contains(from, "泰国") {
	//	contry="tai"
	//}
	gorm.SaveModelInfo(contry, &model)
}
func AnalyzeCompanyHomePageHtml(client *http.Client, url string, companyId int, i int) int {
	resp, err := client.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode > 400 {
		return conf.UrlInvalid
	}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return conf.AnalysisHtmlFaild
	}
	if i == 1 && conf.SaveCompanyGroupRatation {
		doc.Find("div.fenlei").Find("p").Find("a").Each(func(i int, s *goquery.Selection) {
			name := s.Text()
			homepage, _ := s.Attr("href")
			id := util.GetIdFromUri(homepage)
			group := model.Group{
				Homepage: homepage,
				Belong:   companyId,
			}
			group.ID = id
			group.Name = name
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
		return conf.AnalysisHtmlFaild
	}
	return conf.Success
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
			c := model.Album{
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
		return conf.AnalysisHtmlFaild
	}
	return conf.Success
}
