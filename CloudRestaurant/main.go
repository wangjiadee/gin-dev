package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	//读取配置文件
	cfg, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err.Error())
	}
	Router := gin.Default()
	registerRouter(Router)
	{
		// Same as this:
		// r.GET("/hello", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "Fuck world",
		// 	})
		// })
	}
	Router.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
