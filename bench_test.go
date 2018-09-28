package bench_test

import (
	"os"
	"os/exec"
	"testing"

	pdf "github.com/gcoka/gen-pdf-bench/wkhtmltopdf"
)

func BenchmarkPuppeteer(b *testing.B) {
	wd, _ := os.Getwd()
	defer os.Chdir(wd)

	os.Chdir("puppeteer")
	for i := 0; i < b.N; i++ {
		cmd := exec.Command("node", "index.js")
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkWkhtmltopdf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pdf.PrintPDF()
	}
}
