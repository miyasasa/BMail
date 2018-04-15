package bmail

import (
	"os"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"strings"
	"log"
	"bytes"
)

func generatePDFFromHTML(attachment Attachment) string {

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
	fileName := dir + "/.disk/" + attachment.Name + ".pdf"

	pdfgFromJSON.WriteFile(fileName)

	return fileName
}
