package controller

import (
	"CloudRestaurant/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberController struct {
}

//Router Enging
func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
}

// http://localhost:8090/api/sendcode?phone=13167582436
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	//实例化service
	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 200,
			"msg":  "发送成功",
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"code": 300,
		"msg":  "发送失败",
	})
}
