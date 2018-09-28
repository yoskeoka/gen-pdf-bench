package pdf_test

import (
	"testing"

	pdf "github.com/gcoka/gen-pdf-bench/wkhtmltopdf"
)

func BenchmarkPrintPDF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pdf.PrintPDF()
	}
}
