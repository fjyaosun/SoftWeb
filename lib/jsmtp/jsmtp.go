package jsmtp

import (
	"fmt"
	"net/smtp"
	"strings"
)

func init() {
	ConfigInit()
}
func SendDefaultMail(body string) {
	hp := strings.Split(SmtpConf.Host, ":")
	auth := smtp.PlainAuth("", SmtpConf.User, SmtpConf.Password, hp[0])
	contentType := "Content-Type: text/html; charset=UTF-8"

	msg := []byte(fmt.Sprintf("To: %s\nFrom: %s\nSubject:%s\n%s\n\n%s",
		SmtpConf.Default.To,
		SmtpConf.User,
		SmtpConf.Default.Subject,
		contentType,
		body))
	to := strings.Split(SmtpConf.Default.To, ",")
	smtp.SendMail(SmtpConf.Host, auth, SmtpConf.User, to, msg)
}
