package main

import (
	"gin-dev/dao"
	"gin-dev/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { //类型断言
		v.RegisterValidation("abc", service.TopicUrl) //注册调用tag和自定义验证器
	}

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
		// v1.GET("/topic/:topic_id", dao.GetTopicDetail)
		// v1.GET("", dao.GetTopicList)
		v1.Use(dao.MustLogin())
		{
			v1.POST("", dao.NewTopic)
			v1.DELETE("/topic/:topic_id", dao.DelTopic)
		}
	}

	v2 := router.Group("/v1/mtopics")
	{
		v2.Use(dao.MustLogin())
		{
			v2.POST("", dao.NewTopics)
		}
	}
	router.Run(":8089")
}
