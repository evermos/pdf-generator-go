Headless PDF Generator From HTML Template

# Prerequisites

- [go1.19](https://go.dev/dl/)
- [wkhtmltopdf](https://wkhtmltopdf.org/downloads.html)

# How to Use

```bash
## build binary
$ go build -o pdfgen .

## run the binary
$ ./pdfgen -t <path_to_html_template_folder> -o <path_to_pdf_output> --data <json_string>
```

# Dynamic Template

Template can be provided as `*.html` files in a folder and supports dynamic data injection.

Template engine used is Go's built-in `html/template`, so syntax to use it can be read in their docs.

Be sure to wrap each `*.html` files with template definition (`{{ define "<name>" }}`) so it can be rendered.

## Advanced Header and Footer

For full customized header and footer, simply provide separate `header.html` and `footer.html` inside the folder that's being passed.

# Dynamic Data

Data can be passed with the `--data` flag as a JSON string, and can contain anything as long as it's a valid JSON format.

Pass the data as an array, so a single template can be used to generate a multiple-page file.

# Debug Mode

Optional `--debug` flag can be passed to print out rendered HTML to stdout.

# Examples

Check out two examples for both basic usage and a more advance one inside the `examples` folder.

Use the template with these parameters to see how they work.

## Basic Example

```bash
$ ./pdfgen -t ./examples/templates/basic -o ./examples/pdf/basic.pdf --data '[{ "title": "PDF Generator from HTML", "isActive": true, "list": ["one", "two"] }]'
```

## Advanced Example

```bash
$ ./pdfgen -t ./examples/templates/advanced -o ./examples/pdf/advanced.pdf --data '[{"seller":"Evermos","invoice_no":"6006-1583374588","purchase_date":"13 Desember 2022","buyer":"Kristika Kadarsih","items":[{"name":"XL Xtra Kuota Combo 40gb","descriptions":["Nomor: 087730815702","Serial Number: 22112017482267"],"qty":1,"unit_price":"Rp98.000","total_price":"Rp98.000"},{"name":"XL Xtra Kuota Combo 10gb","descriptions":["Nomor: 087730815702","Serial Number: 22112017482267"],"qty":2,"unit_price":"Rp18.000","total_price":"Rp36.000"}],"subtotal":"Rp134.000","total_items":"Rp134.000","admin_fee":"Rp5.000","discount":"-Rp20.000","total":"Rp119.000","status":"paid","published_date":"13 Desember 2022 14:22 WIB"},{"seller":"Evermos","invoice_no":"6006-1583374588","purchase_date":"13 Desember 2022","buyer":"Kristika Kadarsih","items":[{"name":"XL Xtra Kuota Combo 10gb","descriptions":["Nomor: 087730815702","Serial Number: 22112017482267"],"qty":1,"unit_price":"Rp98.000","total_price":"Rp98.000"}],"subtotal":"Rp98.000","total_items":"Rp98.000","admin_fee":"Rp5.000","discount":"-Rp20.000","total":"Rp83.000","status":"refund","published_date":"13 Desember 2022 14:22 WIB"}]'
```

# References

Read more about the usage of the syntax used for the HTML files in their respective docs:

- [`wkhtmltopdf`](https://wkhtmltopdf.org/index.html)
- [`html/template`](https://pkg.go.dev/html/template)
- [`text/template`](https://pkg.go.dev/text/template)
