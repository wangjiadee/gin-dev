package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UserRegister struct {
	Username string `form:"username" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	{ // Handle deal with get request
		r.Handle("GET", "/v1", func(c *gin.Context) {

			fmt.Println("获取请求接口:" + c.FullPath())
			name := c.DefaultQuery("name", "")
			fmt.Println(name)

			c.Writer.Write([]byte("Hello," + name))
		})
	}

	{ // Pass values ​​through a single form
		r.Handle("POST", "/login", func(context *gin.Context) {

			fmt.Println(context.FullPath())
			//userName
			username := context.PostForm("username")
			fmt.Println(username)

			//passWord
			password := context.PostForm("pwd")
			fmt.Println(password)

			context.Writer.Write([]byte("User login"))
		})
	}

	{
		r.Handle("GET", "/register", func(context *gin.Context) {

			fmt.Println(context.FullPath())
			var register UserRegister
			if err := context.ShouldBind(&register); err != nil {
				log.Fatal(err.Error())
				return
			}
			fmt.Println(register.Username)
			fmt.Println(register.Phone)
			context.Writer.Write([]byte(register.Username + " Register "))
		})
	}

	r.Run(":8089") // 监听并在 0.0.0.0:8080 上启动服务
}
