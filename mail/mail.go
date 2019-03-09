package mail

import (
	"crypto/tls"

	"github.com/laughmaker/go-pkg/conf"
	"github.com/laughmaker/go-pkg/log"
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
	d.TLSConf = &tls.Conf{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Error(err)
	}
}
