package bmail

import (
	"gopkg.in/gomail.v2"
	"BMail/config"
)

func generate(mail Bmail) *gomail.Message {

	message := gomail.NewMessage()
	message.SetHeader("From", mail.From)
	message.SetHeader("To", mail.Recipients ...)
	message.SetHeader("Cc", mail.Cc ...)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody(mail.Content.ContentType, mail.Content.Body)

	for _, attachment := range mail.Attachments {
		message.Attach(generatePDFFromHTML(attachment))
	}

	return message
}

func send(mail Bmail) {

	message := generate(mail)

	d := gomail.NewDialer(config.Config.SmtpServer, config.Config.SmtpServerPort, config.Config.Username, config.Config.Password)

	if err := d.DialAndSend(message); err != nil {
		panic(err)
	}
}
