package service

// golang 内置验证器
// github_docs: https://pkg.go.dev/github.com/go-playground/validator/v10
//
type Topic struct {
	TopicID         int    `json:"id"`
	TopicTitle      string `json:"title" binding:"min=4,max=20"`
	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string `json:"ip" binding:"ipv4"`
	TopicScore      int    `json:"score" binding:"omitempty,gt=5"`
}

type TopicQuery struct {
	UserName string `json:"username" form:"username" binding:"required"`
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

// CreateTopic: Temporary creation of entity function
func CreateTopic(id int, title string, stitle string, ip string, score int) Topic {
	return Topic{id, title, stitle, ip, score}
}
