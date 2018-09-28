package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

func BenchmarkCreatePdf(b *testing.B) {

	for i := 0; i < b.N; i++ {
		pdfBuf, err := CreatePdf("https://google.com", 800.0, 600.0, 5*time.Second)
		if err != nil {
			panic(err)
		}
		file, err := ioutil.TempFile("", "")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Println(file.Name())
		_, err = file.Write(pdfBuf)
		if err != nil {
			panic(err)
		}
	}
}

func CreatePdf(urlRequest string, width, height float64, timeout time.Duration) ([]byte, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)

	// Open a DOMContentEventFired client to buffer this event.
	domContent, err := c.Page.DOMContentEventFired(ctx)
	if err != nil {
		return nil, err
	}
	defer domContent.Close()

	// Enable events on the Page domain, it's often preferrable to create
	// event clients before enabling events so that we don't miss any.
	if err = c.Page.Enable(ctx); err != nil {
		return nil, err
	}

	// Create the Navigate arguments with the optional Referrer field set.
	navArgs := page.NewNavigateArgs("https://www.google.com")
	_, err = c.Page.Navigate(ctx, navArgs)
	if err != nil {
		return nil, err
	}

	// Wait until we have a DOMContentEventFired event.
	if _, err = domContent.Recv(); err != nil {
		return nil, err
	}

	// Print to PDF
	printToPDFArgs := page.NewPrintToPDFArgs().
		SetLandscape(true).
		SetPrintBackground(true).
		SetMarginTop(0).
		SetMarginBottom(0).
		SetMarginLeft(0).
		SetMarginRight(0).
		SetPrintBackground(true).
		SetPaperWidth(width).
		SetPaperHeight(height)
	print, _ := c.Page.PrintToPDF(ctx, printToPDFArgs)
	return print.Data, nil
}
