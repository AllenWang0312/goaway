package api

import (
	model "../../model/meituri"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetCommits(c *gin.Context) {
	var commits []model.Feedback
	pageNo, err1 := strconv.Atoi(c.Query("pageNo"))
	pageSize, err2 := strconv.Atoi(c.Query("pageSize"))
	if err1 != nil {
		c.JSON(200, gin.H{"msg": err1.Error()})
	} else if err2 != nil {
		c.JSON(200, gin.H{"msg": err2.Error()})
	} else {
		db.Table("feedbacks").Order("likes desc").Limit(pageSize).Offset((pageNo - 1) * pageSize).Find(&commits)
		c.JSON(200, gin.H{"data": commits})
	}
}
func Commit(c *gin.Context) {
	user_id := getUserIdWithToken(c)
	if user_id == -1 {
		return
	} else if user_id > 0 {
		content := c.PostForm("content")
		tel:= c.PostForm("tel")
		//content=mahonia.NewDecoder("utf-8").ConvertString(content)
		println(content)
		//images :=[]string{}
		//_ = json.Unmarshal(, &images)
		//println(images)
		images := []byte(c.PostForm("images"))
		now := time.Now()
		var feedback = model.Feedback{
			Tel:tel,
			Userid:     user_id,
			Content:    content,
			Images:     images,
			State:      1,
			Createtime: now,
		}
		db.Create(&feedback)
		c.JSON(200, gin.H{"toast": "提交成功"})
	}else {
		c.JSON(200, gin.H{"toast": "提交失败"})
	}

}

func LikeCommit(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if nil == err {
		feed := model.Feedback{
			ID: id,
		}
		db.First(&feed)
		feed.Likes += 1
		db.Create(feed)
		c.JSON(200, gin.H{"toast": "点赞成功"})
	} else {
		c.JSON(404, gin.H{"toast": "id 不能为空"})
	}
}
