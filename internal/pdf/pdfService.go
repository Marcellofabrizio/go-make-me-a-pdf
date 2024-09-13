package pdf

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/spf13/afero"
)

func GeneratePdf(orderId string, emissionDate string, licensePlate string) {

	webUrl := fmt.Sprintf("https://app.gringo.com.vc/report-vehicle/%s?emissionDate=%s&licensePlate=%s", orderId, emissionDate, licensePlate)

	fmt.Printf("%s\n", webUrl)

	page := rod.New().MustConnect().MustPage(webUrl).MustWaitDOMStable()

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
