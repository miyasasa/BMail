package bmail

import (
	"gopkg.in/gomail.v2"
	"BMail/config"
)

func GenerateGoMail(mail Bmail) *gomail.Message {

	message := gomail.NewMessage()
	message.SetHeader("From", mail.From)
	message.SetHeader("To", mail.Recipients ...)
	message.SetHeader("Cc", mail.Cc ...)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody(mail.Content.ContentType, mail.Content.Body)

	//message.Attach("/home/Alex/lolcat.jpg")

	return message
}

func send(mail Bmail) {

	message := GenerateGoMail(mail)

	d := gomail.NewDialer(config.Config.SmtpServer, config.Config.SmtpServerPort, config.Config.Username, config.Config.Password)

	// Send the email
	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
}
