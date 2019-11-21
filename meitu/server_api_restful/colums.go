package api_restful

import (
	"../../conf"
	model "../model/meituri"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

func GetColumDetail(c *gin.Context) {
	modelId := c.Query("model_id")
	columId := c.Query("colum_id")
	colum := model.Colum{}

	db.Where("id = ?", columId).First(&colum)

	if(colum.ID>0){

		p := "/" + modelId + "/" + columId + "/"
		path := "../meituri" + p
		//downloadFile(durl,path,filename)
		rd, err := ioutil.ReadDir(path)
		if err == nil {
			paths := []string{}
			for _, fi := range rd {
				if fi.IsDir() {
					fmt.Printf("[%s]\n", fi.Name())
				} else {
					fmt.Println(fi.Name())
					p := conf.FILE_SERVER + p + fi.Name()
					paths = append(paths, p)
					fmt.Println(len(paths), cap(paths), paths, p)
				}
			}
			colum.Images = paths
			c.JSON(200, gin.H{"data": colum})
			return
		}else {
			c.JSON(200, gin.H{"data": colum})
		}
	}else {
		c.JSON(404, gin.H{"message": "colum not exist"})
	}



}
func GetColumsList(c *gin.Context) {
	tag, err0 := strconv.Atoi(c.Query("tag"))

	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	var colums = []model.Colum{}

	if err0 == nil {
		db.Where("tags LIKE ?", tag).Order("id desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&colums) //.Order("created_at desc")
		//c.String(200,)
		c.JSON(200, gin.H{"data": colums})
	} else {
		if nil == err1 && nil == err2 {
			//if len(search) == 0 {
			//} else {
			//}
			db.Limit(pageSize).Offset((pageNo - 1) * pageSize).Order("id desc").Find(&colums) //.Order("created_at desc")
			//c.String(200,)
			c.JSON(200, gin.H{"data": colums})
		} else {
			c.JSON(404, gin.H{"status": 0, "msg": "缺少参数"})
		}
	}
}
