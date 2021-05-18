package dao

import (
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
	c.String(http.StatusOK, "获取到的ID为%s", c.Param("topic_id"))
}

func NewTopic(c *gin.Context) {
	c.String(http.StatusOK, "新增帖子成功!")
}
func DelTopic(c *gin.Context) {
	c.String(http.StatusOK, "删除帖子成功!")
}
