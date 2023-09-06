package email

import (
	"blog_server/global"

	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Verification Subject = "Verification Code"
	Notification Subject = "Notification"
	Alarm        Subject = "Alarm"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(receiverEmail, body string) error {
	return send(receiverEmail, string(a.Subject), body)
}

func NewVerification() Api {
	return Api{
		Subject: Verification,
	}
}
func NewNotification() Api {
	return Api{
		Subject: Notification,
	}
}
func NewAlarm() Api {
	return Api{
		Subject: Alarm,
	}
}

func send(receiverEmail, subject, body string) error {
	e := global.Config.Email
	return sendMail(
		e.SenderName,
		e.Password,
		e.Host,
		e.Port,
		receiverEmail,
		e.SenderEmail,
		subject,
		body,
	)
}

func sendMail(senderName, pwd, host string, port int, receiverEmail, senderEmail string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(senderEmail, senderName))
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, senderEmail, pwd)
	err := d.DialAndSend(m)
	return err
}
