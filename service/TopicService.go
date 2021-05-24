package service

type Topic struct {
	TopicID    int    `json:"id"`
	TopicTilte string `json:"title"`
}

type TopicQuery struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

// CreateTopic: Temporary creation of entity function
func CreateTopic(id int, title string) Topic {
	return Topic{id, title}
}
