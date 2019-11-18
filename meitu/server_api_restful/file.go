package api_restful

import (
	"../../configs"
	"github.com/gin-gonic/gin"
	"os"
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
