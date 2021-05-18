package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/topic/:topic_id", func(c *gin.Context) {
		c.String(http.StatusOK, "获取到的ID为%s", c.Param("topic_id"))
	})
	router.GET("/topic/top", func(c *gin.Context) {
		c.String(http.StatusOK, "123")
	})

	router.Run()
}
