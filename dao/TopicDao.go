package dao

import (
	"gin-dev/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MustLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "缺少token")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

//GetTopicDetail: Past corresponding posts by id
func GetTopicDetail(c *gin.Context) {
	// c.String(http.StatusOK, "获取到的ID为%s", c.Param("topic_id"))
	c.JSON(http.StatusOK, service.CreateTopic(2019, "帖子标题"))
}

func NewTopic(c *gin.Context) {
	topic := service.Topic{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, "[ERROR:] Parameter error", err.Error())
	} else {
		c.JSON(http.StatusOK, topic)
	}
	// c.String(http.StatusOK, "新增帖子成功!")
}
func DelTopic(c *gin.Context) {
	c.String(http.StatusOK, "删除帖子成功!")
}

func GetTopicList(c *gin.Context) {
	query := service.TopicQuery{}
	err := c.BindQuery(&query)
	if err != nil {
		c.String(400, "[ERROR:] Parameter error", err.Error())
	} else {
		c.JSON(http.StatusOK, query)
	}
}
