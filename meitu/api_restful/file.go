package api_restful

import "github.com/gin-gonic/gin"

func UploadFile(c *gin.Context) {
	header, err := c.FormFile("file")
	if err != nil {
		println(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	if err := c.SaveUploadedFile(header, header.Filename); err != nil {
		println(err.Error())
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
}
