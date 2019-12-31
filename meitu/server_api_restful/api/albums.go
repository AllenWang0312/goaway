package api

import (
	"../../../conf"
	model "../../model/meituri"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func GetAlbumDetail(c *gin.Context) {
	useridstr := getUserIdStrWithToken(c)
	userid, err := strconv.Atoi(useridstr)
	if err == nil && userid > 0 {
		modelIdStr := c.Query("model_id")
		albumIdStr := c.Query("album_id")
		albumId, _ := strconv.Atoi(albumIdStr)
		album := model.Album{}
		db.Where("id = ?", albumIdStr).Preload("Model").First(&album)
		if album.ID > 0 {
			//downloadFile(durl,path,filename)
			var now = time.Now().Format("2006-01-02")
			var record = model.VisitHistroy{
				Albumid:  albumId,
				Userid:   userid,
				Date:     now,
				Relation: albumIdStr + "_" + useridstr + "_" + now,
			}
			var tableName = conf.VisitHistroy + strconv.Itoa(userid/1000)
			if !db.HasTable(tableName) {
				db.Table(tableName).Create(model.VisitHistroy{})
			}
			db.Table(tableName).Create(&record)

			appendImagesForAlbum(modelIdStr, albumIdStr, &album)

			c.JSON(200, gin.H{"data": &album})
		} else {
			c.JSON(404, gin.H{"message": "album not exist"})
		}
	}

}

func appendImagesForAlbum(modelIdStr string, albumIdStr string, album *model.Album) {
	p := "/" + modelIdStr + "/" + albumIdStr + "/"
	path := conf.FSMuri + p
	rd, err := ioutil.ReadDir(path)
	if err == nil {
		var paths []string
		for _, fi := range rd {
			if fi.IsDir() {
				fmt.Printf("[%s]\n", fi.Name())
			} else {
				fmt.Println(fi.Name())
				p := conf.FILE_SERVER + conf.Muri + p + fi.Name()
				paths = append(paths, p)
				fmt.Println(len(paths), cap(paths), paths, p)
			}
		}
		album.Images = paths
	} else {
		println(err.Error())
	}

}
func GetAlbumsList(c *gin.Context) {
	plat := c.Query("platform")
	tag := c.Query("tag")
	search := c.Query("search")
	println(tag, search)
	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	var albums []model.Album
	if (strings.EqualFold(plat, "fengniao")) {
		db.Table("fengniao_album").Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&albums) //.Order("created_at desc")
		c.JSON(200, gin.H{"data": &albums})
	}else{
		if nil == err1 && nil == err2 {
			if len(tag) > 0 {
				db.Where("tags like ?", tag).Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&albums) //.Order("created_at desc")
				//c.String(200,)
				//for _, a := range albums {
				//	if reader, err := os.Open(model.GetAlbumCover(&a, false)); err != nil {
				//		defer reader.Close()
				//		im, _, err := image.DecodeConfig(reader)
				//		if err != nil {
				//			fmt.Fprintf(os.Stderr, "%v\n", err)
				//			continue
				//		}
				//		var cover = model.ImageInfo{
				//			Url:    model.GetAlbumCover(&a, true),
				//			Width:  im.Width,
				//			Height: im.Height,
				//			Scale:  float64(im.Width) / float64(im.Height),
				//		}
				//		a.Cover=cover
				//		fmt.Printf("%d %d\n", im.Width, im.Height)
				//	} else {
				//		fmt.Println("Impossible to open the file")
				//	}
				//}
				c.JSON(200, gin.H{"data": &albums})
			} else if len(search) > 0 {
				db.Where("title like ?", search).Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&albums)
				c.JSON(200, gin.H{"data": &albums})
			} else {
				db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Order("id desc").Find(&albums) //.Order("created_at desc")
				c.JSON(200, gin.H{"data": &albums})
			}

		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}
}
