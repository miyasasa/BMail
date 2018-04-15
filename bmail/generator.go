package bmail

import (
	"gopkg.in/gomail.v2"
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"strings"
	"log"
	"bytes"
)

func ExampleNewPDFGenerator() {

	const html = `<!doctype html><html><head><title>WKHTMLTOPDF TEST</title></head><body>HELLO PDF</body></html>`

	// Client code
	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	// The html string is also saved as base64 string in the JSON file
	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	// The JSON can be saved, uploaded, etc.

	// Server code, create a new PDF generator from JSON, also looks for the wkhtmltopdf executable
	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	// Create the PDF
	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Use the PDF
	dir, _ := os.Getwd()
	pdfgFromJSON.WriteFile(dir + "/.disk/test.pdf")
}

func GenerateGoMail(mail Bmail) *gomail.Message {

	ExampleNewPDFGenerator()

	message := gomail.NewMessage()
	message.SetHeader("From", mail.From)
	message.SetHeader("To", mail.Recipients ...)
	message.SetHeader("Cc", mail.Cc ...)
	message.SetHeader("Subject", mail.Subject)
	message.SetBody(mail.Content.ContentType, mail.Content.Body)

	dir, _ := os.Getwd()
	message.Attach(dir + "/.disk/test.pdf")

	return message
}
