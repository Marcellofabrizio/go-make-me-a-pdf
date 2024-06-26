// Package main ...
package main

import (
	"fmt"

	"github.com/go-rod/rod"

	"context"
	"log"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	rod_pod()
}

func rod_pod() {
	rod.New().MustConnect().MustPage("https://app-staging.gringo.com.vc/report-vehicle/999995756").MustWaitStable().MustPDF("sample.pdf")
	fmt.Println("wrote sample.pdf")
}

func chromedp_poc() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(
		ctx,
		printToPDF(`https://app-staging.gringo.com.vc/report-vehicle/999995756`,
			&buf),
		chromedp.Sleep(5000)); err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("sample.pdf", buf, 0o644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("wrote sample.pdf")
}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
