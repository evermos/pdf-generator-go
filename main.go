package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
)

func main() {
	var (
		err        error
		data       map[string]any
		htmlBuffer bytes.Buffer
	)

	// read from params
	templatePath := flag.String("t", "./templates/index.html", "Template file path")
	outputPath := flag.String("o", "./outputs/index.pdf", "Output file path")
	jsonStr := flag.String("data", "{}", "Data in JSON format")
	flag.Parse()

	// parse json
	err = json.Unmarshal([]byte(*jsonStr), &data)
	if err != nil {
		log.Fatal(err)
	}

	// read HTML template into buffer
	htmlTemplate, err := template.ParseFiles(*templatePath)
	if err != nil {
		log.Fatal(err)
	}

	// merge template with data
	err = htmlTemplate.Execute(&htmlBuffer, data)
	if err != nil {
		log.Fatal(err)
	}

	// init pdf generator
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// read from buffer
	pdfGenerator.AddPage(wkhtmltopdf.NewPageReader(&htmlBuffer))
	pdfGenerator.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfGenerator.Dpi.Set(300)

	// create pdf from htmlBuffer
	err = pdfGenerator.Create()
	if err != nil {
		log.Fatal(err)
	}

	// write pdf into file
	err = pdfGenerator.WriteFile(*outputPath)
	if err != nil {
		log.Fatal(err)
	}
}
