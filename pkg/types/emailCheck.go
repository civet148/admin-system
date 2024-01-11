package types

import (
	"fmt"
	"github.com/civet148/log"
	"gopkg.in/gomail.v2"
	"strconv"
)

// 验证码位数
var Width = 6

// 验证码保存时间 单位：s
var KeepSecond int64 = 5 * 60

type EmailCheckStr struct {
	CheckCode string `json:"check_code"`
	TimeStamp int64  `json:"time_stamp"`
}

type EmailConfig struct {
	EmailServer string `json:"email_server"`
	Port        string `json:"port"`
	EmailName   string `json:"email_name"`
	AuthCode    string `json:"auth_code"`
	SendName    string `json:"send_name"`
}

func SendMail(emailConfig EmailConfig, toMail string, checkCode string) (err error) {
	m := gomail.NewMessage()

	//发送人
	m.SetAddressHeader("From", emailConfig.EmailName, "动环")
	//接收人
	m.SetHeader("To", toMail)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", emailConfig.SendName)
	//内容
	m.SetBody("text/html", fmt.Sprintf("<h1>尊敬的用户您好：<br></br></h1><div>您的验证码为：<br></br></div><h3>%s</h3><div><br></br>出于安全原因，请勿将验证码透露给别人，以免造成不必要的损失</div>", checkCode))
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	var port int
	port, err = strconv.Atoi(emailConfig.Port)
	if err != nil {
		log.Errorf("emailConfig.Port,err: %s", err.Error())
		return err
	}
	d := gomail.NewDialer(emailConfig.EmailServer, port, emailConfig.EmailName, emailConfig.AuthCode)

	// 发送邮件
	if err = d.DialAndSend(m); err != nil {
		log.Errorf("Send mail failed,err: %s", err.Error())
		return err
	}
	log.Info("send to mail [%s] code [%s]\n", toMail, checkCode)
	return nil
}
