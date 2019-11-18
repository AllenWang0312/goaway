package api_restful

import (
	model "../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetCommits(c *gin.Context) {
	var commits []model.Feedback
	pageNo, err1 := strconv.Atoi(c.Query("pogeNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err1 != nil {
		c.JSON(200, gin.H{"msg": err1.Error()})
	} else if err2 != nil {
		c.JSON(200, gin.H{"msg": err2.Error()})
	} else {
		db.Order("like desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&commits)
		c.JSON(200, gin.H{"data": commits})
	}
}
func Commit(c *gin.Context) {
	//if tokenEnable(c) {
	user_id, err := strconv.Atoi(c.PostForm("user_id"))
	content := c.PostForm("content")
	images := c.PostForm("images")
	println(images)
	if nil == err {
		var feedback = model.Feedback{
			Userid:     user_id,
			Content:    content,
			Images:     images,
			State:      1,
			Createtime: time.Now().Format("2006-01-02 15:04:05"),
		}
		db.Save(&feedback)
		c.JSON(200, gin.H{"toast": "提交成功"})
	}
	//}
}

func LikeCommit(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if nil == err {
		feed := model.Feedback{
			Id: id,
		}
		db.First(&feed)
		feed.Like += 1
		db.Save(feed)
		c.JSON(200, gin.H{"toast": "点赞成功"})
	} else {
		c.JSON(404, gin.H{"toast": "id 不能为空"})
	}
}
