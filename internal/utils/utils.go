package utils

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"gopkg.in/gomail.v2"
	"strings"
)

// 发送邮件
func SentEmail(content string, subject string, addressHeader string, config g.Map) {

	var ctx = context.Background()
	// 设置 SMTP 服务器的认证信息
	smtp := gconv.String(config["smtp"])
	smtpPort := 465
	senderEmail := gconv.String(config["smtpEmail"])
	senderPassword := gconv.String(config["smtpPass"])

	body := content
	// 邮件内容
	toEmail := gconv.String(config["remindEmail"])
	toEmails := strings.Split(toEmail, "|")

	m := gomail.NewMessage()
	m.SetHeader("To", toEmails...)
	m.SetHeader("Subject", subject)
	m.SetAddressHeader("From", senderEmail, addressHeader)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtp, smtpPort, senderEmail, senderPassword)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	g.Log().Info(ctx, "邮件发送成功")

}
