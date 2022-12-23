Headless PDF Generator From HTML Template

# Prerequisites

- [go1.19](https://go.dev/dl/)
- [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html)

# How to Use

```bash
## build binary
$ go build -o pdfgen .

## run the binary
$ ./pdfgen -t <path_to_html_template> -o <path_to_pdf_output> --data <json_string>
```

# Dynamic Template

Template can be provided as a `.html` file and supports dynamic data injection.

Template engine used is Go's built-in `html/template`, so syntax to use it can be read in their docs.

# Dynamic Data

Data can be passed with the `--data` flag as a JSON string, and can contain anything as long as it's a valid JSON format.
