package bmail

import (
	"gopkg.in/gomail.v2"
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"strings"
	"log"
	"bytes"
)

func generatePDFfromHTML(attachment Attachment) {

	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(attachment.Body)))
	pdfg.Dpi.Set(600)

	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	dir, _ := os.Getwd()
	pdfgFromJSON.WriteFile(dir + "/.disk/" + attachment.Name + ".pdf")

}

func GenerateGoMail(mail Bmail) *gomail.Message {

	message := gomail.NewMessage()
	message.SetHeader("From", mail.From)
	message.SetHeader("To", mail.Recipients ...)
	message.SetHeader("Cc", mail.Cc ...)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody(mail.Content.ContentType, mail.Content.Body)

	dir, _ := os.Getwd()
	for _, attachment := range mail.Attachments {
		generatePDFfromHTML(attachment)
		message.Attach(dir + "/.disk/" + attachment.Name + ".pdf")
	}

	return message
}
