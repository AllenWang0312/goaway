package api

import (
	"../../../conf"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strconv"
)

func UploadFile(c *gin.Context) {
	name := c.PostForm("name")
	path := c.PostForm("path")

	header, _ := c.FormFile("file")

	//if err != nil {
	//	println(err.Error())
	//	c.JSON(400, gin.H{"msg": err.Error()})
	//	return
	//}
	if len(path) > 0 && len(name) > 0 {
		if err := c.SaveUploadedFile(header, name); err != nil {
			//println(err.Error())
			c.JSON(400, gin.H{"msg": err.Error()})
		} else {
			abspath := conf.FSRoot + "/file" + path
			err0 := os.MkdirAll(abspath, 0)
			if nil != err0 {

			} else {
				err = os.Rename(conf.ProjectRoot+"/"+name, abspath+"/"+name)
				if nil != err {
					c.JSON(404, gin.H{"msg": "rename file faild:" + err.Error()})
				} else {
					c.JSON(200, gin.H{"url": conf.FILE_SERVER + "/file/" + path + "/" + name})
				}
			}

		}
	} else {
		c.JSON(400, gin.H{"msg": "path 不能为空"})
	}
}
func uploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err == nil {
		filename := header.Filename
		out, err := os.Create("assets/res/uploadFile/excel/" + filename)
		if err == nil {
			defer out.Close()
			_, err = io.Copy(out, file)
			if err == nil {
				log.Println("上传表格成功")
				res := map[string]interface{}{
					"filePath": "/res/uploadFile/excel/" + filename,
					"fileName": filename,
				}
				c.JSON(200, gin.H{"toast": "上传表格成功", "data": res})
			} else {
				c.JSON(200, gin.H{"toast": "复制文件出错"})
			}
		} else {
			c.JSON(200, gin.H{"toast": "创建文件出错"})
		}
	} else {
		c.JSON(200, gin.H{"toast": "接收表格出错"})
	}
}
func UploadFiles(c *gin.Context) {
	var id = getUserIdWithToken(c)
	if (id > 0) {
		form, _ := c.MultipartForm()
		files := form.File["files[]"]
		var result []string
		for _, file := range files {
			//log.Println(file.Filename)
			var path="/file/" + strconv.Itoa(id) + "/feedback/"
			_ = os.MkdirAll(conf.FSRoot + path, 0777)
			out, err := os.Create(conf.FSRoot + path + file.Filename)
			filereader, err1 := file.Open()
			if err == nil && err1 == nil {
				defer out.Close()
				_, err = io.Copy(out, filereader)
				if err == nil {
					log.Println("上传表格成功")
					var url = conf.FILE_SERVER+path + file.Filename
					res := map[string]interface{}{
						"filePath": url,
						"fileName": file.Filename,
					}
					result = append(result, url)
					println(res)
				} else {
					c.JSON(200, gin.H{"toast": "复制文件出错"})
				}
			} else {
				c.JSON(200, gin.H{"toast": "创建文件出错"})
			}
		}
		c.JSON(200, gin.H{"toast": "上传成功", "data": result})
	} else {
		c.JSON(200, gin.H{"toast": "没有权限,请先登录"})
	}

}
