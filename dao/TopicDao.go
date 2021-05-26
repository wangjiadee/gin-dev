package dao

import (
	"gin-dev/db"
	"gin-dev/service"
	"github.com/gin-gonic/gin"

	"net/http"
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
	// dsn := "root:12345678@tcp(192.168.230.150:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("[ERROR:] Parameter error", err.Error())
	// }
	tid := c.Param("topic_id")
	topics := service.Topic{}
	db.DB.Find(&topics, tid)
	c.JSON(http.StatusOK, topics)
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

func NewTopics(c *gin.Context) {
	topic := service.Topics{}
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, "[ERROR:] Parameter error", err.Error())
	} else {
		c.JSON(http.StatusOK, topic)
	}

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
