package service

import (
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/apsdehal/go-logger"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) Sendcode(phone string) bool {
	log, err := logger.New("test", 1, os.Stdout)
	if err != nil {
		panic(err) // Check for error
	}
	strconv.Itoa(6)
	//1.产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	//2.调用阿里云sdk 完成发送
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Error(err.Error())
		return false
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	if err != nil {
		log.Error(err.Error())
		return false
	}
	request.TemplateParam = string(par)

	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		log.Error(err.Error())
		return false
	}

	//3.接收返回结果，并判断发送状态
	//短信验证码发送成功
	if response.Code == "OK" {
		return true
	}
	return false
}
