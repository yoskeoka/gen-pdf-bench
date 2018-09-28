package pdf

import wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"

// PrintPDF prints url to pdf.
func PrintPDF() {

	// Create new PDF generator
	pdf, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	// set default values
	pdf.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdf.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdf.PageHeight.Set(297)
	pdf.PageWidth.Set(210)

	///usr/local/bin/wkhtmltopdf --page-size Letter --print-media-type --no-background --encoding UTF-8 --orientation portrait --page-width 100 --page-height 148 --margin-top 0 --margin-left 0 --margin-right 0 --margin-bottom 0 --background
	p := wkhtmltopdf.NewPage("https://google.com")
	p.PrintMediaType.Set(true)
	p.NoBackground.Set(true)
	p.Encoding.Set("UTF-8")

	// wait for JavaScript
	// p.JavascriptDelay.Set(100)

	pdf.AddPage(p)
	// pdf.Dpi.Set(400)
	err = pdf.Create()
	if err != nil {
		panic(err)
	}
	pdf.WriteFile("wk_sample.pdf")
}
