package main

import (
	"fmt"
	"gin-dev/service"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// {
	// 	router := gin.Default()
	// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { //类型断言
	// 		v.RegisterValidation("abc", service.TopicUrl) //注册调用tag和自定义验证器
	// 	}
	// 	// router.GET("/topic/:topic_id", func(c *gin.Context) {
	// 	// 	c.String(http.StatusOK, "获取到的ID为%s", c.Param("topic_id"))
	// 	// })
	// 	// router.GET("/topic/top", func(c *gin.Context) {
	// 	// 	c.String(http.StatusOK, "123")
	// 	// })
	// 	// router.Run()
	// 	//Routing Grouping and use Code block
	// 	v1 := router.Group("/v1/push")
	// 	{
	// 		v1.GET("", func(c *gin.Context) {
	// 			if c.Query("username") == "" {
	// 				c.String(http.StatusOK, "获取经验了")
	// 			} else {
	// 				c.String(http.StatusOK, "hi! %s", c.Query("username"))
	// 			}
	// 		})
	// 		// v1.GET("/topic/:topic_id", dao.GetTopicDetail)
	// 		// v1.GET("", dao.GetTopicList)
	// 		v1.Use(dao.MustLogin())
	// 		{
	// 			v1.POST("", dao.NewTopic)
	// 			v1.DELETE("/topic/:topic_id", dao.DelTopic)
	// 		}
	// 	}
	// 	v2 := router.Group("/v1/mtopics")
	// 	{
	// 		v2.Use(dao.MustLogin())
	// 		{
	// 			v2.POST("", dao.NewTopics)
	// 		}
	// 	}
	// 	router.Run(":8089")
	// }
	dsn := "root:12345678@tcp(192.168.230.150:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[ERROR:] Parameter error", err.Error())
	}
	{
		//Connection data test
		// rows, _ := db.Raw("SELECT topic_id, topic_title FROM topic").Rows()
		// for rows.Next() {
		// 	var t_id int
		// 	var t_title string
		// 	rows.Scan(&t_id, &t_title)
		// 	fmt.Println(t_id, t_title)
		// }
	}

	// {
	// 	//Get a piece of data
	// 	TopicClass := &service.TopicClass{}
	// 	db.Table("topic_class").First(&TopicClass, 2)
	// 	fmt.Println(TopicClass)
	// }

	{
		//Get multiple pieces of data
		var tcs []service.TopicClass
		db.Where("class_name=?", "工具部").Table("topic_class").Find(&tcs)
		fmt.Println(tcs)

	}

	// defer db.Close()
}
