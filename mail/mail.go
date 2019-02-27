package mail

import (
	"crypto/tls"

	"github.com/laughmaker/go-pkg/config"
	"github.com/laughmaker/go-pkg/log"
	"gopkg.in/gomail.v2"
)

/**
 * attach 为文件路径，例如 "readme.md"
 */
func Send(to, subject, body, attach string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.MailConfig.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if attach != "" {
		m.Attach(attach)
	}

	d := gomail.NewDialer(config.MailConfig.Host, config.MailConfig.Port, config.MailConfig.User, config.MailConfig.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Error(err)
	}
}
