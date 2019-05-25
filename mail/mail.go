package mail

import (
	"fmt"

	"go-pkg/conf"

	"gopkg.in/gomail.v2"
)

/**
 * attach 为文件路径，例如 "readme.md"
 */
func Send(to, subject, body, attach string) {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Mail.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if attach != "" {
		m.Attach(attach)
	}

	d := gomail.NewDialer(conf.Mail.Host, conf.Mail.Port, conf.Mail.User, conf.Mail.Password)

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}
}
