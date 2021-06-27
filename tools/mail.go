package tools

import "gopkg.in/gomail.v2"

// SendMail 发送邮件
// To是接收方邮件地址, subject是邮件主题, body是正文
func SendMail(to, subject, body string) error {
	const (
		user     = "2191236185@qq.com"
		password = "oucoiboflrgndifh"
		host     = "smtp.qq.com"
		port     = 465
	)

	m := gomail.NewMessage()
	m.SetHeader("From", "landlord.huining.tech"+"<"+user+">")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, user, password)
	err := d.DialAndSend(m)
	return err
}
