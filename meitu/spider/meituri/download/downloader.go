package download

import (
	"../../../../conf"
	"../../../../util"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var END = "cn"
var WG sync.WaitGroup

func DownloadAlbumCover(modelId int, columId int) {
	filename := "0.jpg"
	durl := conf.OldHost + "/a/1/" + strconv.Itoa(columId) + "/" + filename
	//resp, err := url.ParseRequestURI(durl)
	//if err != nil {
	//	//panic("url err")
	//	println(err.Error())
	//	return -2
	//}
	path := conf.FSRoot + "/meituri_" + END + "/" + strconv.Itoa(modelId) + "/" + strconv.Itoa(columId) + "/"

	DownloadImage(durl, path, filename)
}

// 下载图片
func DownloadImage(durl string, path string, fileName string) int {
	//downloadFile(durl,path,filename)
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
