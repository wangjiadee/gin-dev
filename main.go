package main

import (
	"gin-dev/dao"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v3 := router.Group("/v3")
	{
		v3.Use(dao.MustLogin())
		{
			v3.GET("/:topic_id", dao.GetTopicDetail)
		}
	}
	router.Run(":8089")
}
