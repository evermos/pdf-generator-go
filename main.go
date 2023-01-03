package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		err          error
		data         []map[string]interface{}
		files        []string
		htmlTemplate *template.Template
		htmlBuffer   bytes.Buffer
		headerPath   string
		footerPath   string
	)

	// read from params
	name := flag.String("name", "htmlIndex", "Template name")
	templateDir := flag.String("t", "examples/templates/simple", "Template file folder")
	outputPath := flag.String("o", "./outputs/index.pdf", "Output file path")
	jsonStr := flag.String("data", "[]", "Data in JSON format")
	debug := flag.Bool("debug", false, "Debug flag")
	flag.Parse()

	// read template directory
	dir, err := os.ReadDir(*templateDir)
	if err != nil {
		log.Fatal(err)
	}

	// get path of all files within templateDir
	for _, file := range dir {
		filename := file.Name()
		if filename == "footer.html" {
			footerPath = filepath.Join(*templateDir, filename)
		} else if filename == "header.html" {
			headerPath = filepath.Join(*templateDir, filename)
		} else {
			filePath := filepath.Join(*templateDir, filename)
			files = append(files, filePath)
		}
	}

	// merge template
	htmlTemplate, err = template.New(*name).ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}

	// parse json into data
	err = json.Unmarshal([]byte(*jsonStr), &data)
	if err != nil {
		log.Fatal(err)
	}

	// merge template with data into buffer
	err = htmlTemplate.ExecuteTemplate(&htmlBuffer, *name, data)
	if err != nil {
		log.Fatal(err)
	}

	// prints rendered HTML if debug flag is passed
	if *debug {
		fmt.Println(&htmlBuffer)
	}

	// init pdf generator
	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// init page reader from buffer
	page := wkhtmltopdf.NewPageReader(&htmlBuffer)

	// optionally inject header and footer
	if headerPath != "" {
		page.HeaderHTML.Set(headerPath)
	}
	if footerPath != "" {
		page.FooterHTML.Set(footerPath)
	}

	// add page into pdf generator (using default options)
	pdfGenerator.AddPage(page)

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
