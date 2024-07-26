// Package main ...
package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/spf13/afero"
)

func main() {
	rod_pod()
}

func rod_pod() {
	page := rod.New().MustConnect().MustPage("https://www.olhonocarro.com.br/resultado-da-consulta/?queryCode=100&queryId=668c36793d0a5db6151a2296").MustWaitDOMStable() //.MustPDF("sample.pdf")

	res := page.MustEval(`() => document.body.scrollHeight`)

	scrollHeight := res.Num()

	fmt.Println("scrollHeight", scrollHeight)

	marginTop := 0.0
	marginBottom := 0.0

	reportPdf, err := page.PDF(&proto.PagePrintToPDF{
		MarginBottom:    &marginBottom,
		MarginTop:       &marginTop,
		PrintBackground: true,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fs := afero.NewOsFs()

	bin, err := afero.ReadAll(reportPdf)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = afero.WriteFile(fs, "report.pdf", bin, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("wrote sample.pdf")

}
