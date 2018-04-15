package bmail

import (
	"gopkg.in/gomail.v2"
	"BMail/config"
)

func send(mail Bmail) {

	message := GenerateGoMail(mail)

	d := gomail.NewDialer(config.Config.SmtpServer, config.Config.SmtpServerPort, config.Config.Username, config.Config.Password)

	// Send the email
	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
}
