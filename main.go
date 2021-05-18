package main

import (
	"gin-dev/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.GET("/topic/:topic_id", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "获取到的ID为%s", c.Param("topic_id"))
	// })
	// router.GET("/topic/top", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "123")
	// })
	// router.Run()

	//Routing Grouping and use Code block
	v1 := router.Group("/v1/push")
	{
		v1.GET("", func(c *gin.Context) {
			if c.Query("username") == "" {
				c.String(http.StatusOK, "获取经验了")
			} else {
				c.String(http.StatusOK, "hi! %s", c.Query("username"))
			}
		})
		v1.GET("/topic/:topic_id", dao.GetTopicDetail)
		v1.Use(dao.MustLogin())
		{
			v1.POST("", dao.NewTopic)
			v1.DELETE("/topic/:topic_id", dao.DelTopic)
		}

		router.Run()
	}
}
