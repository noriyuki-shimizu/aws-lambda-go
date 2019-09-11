package trainDelay

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// AreaDelay is struct
type AreaDelay struct {
	Doc *goquery.Document
}

// Goto return error
func (a *AreaDelay) Goto(areaCode string) error {
	accessURL := os.Getenv("TRAININFO_YAHOO_MAIN_URL")
	accessURL += os.Getenv("TRAININFO_YAHOO_CONTEXT")
	accessURL += os.Getenv("TRAININFO_YAHOO_SUB_AREA")
	accessURL += areaCode + "/"

	doc, err := goquery.NewDocument(accessURL)
	if err != nil {
		return err
	}
	a.Doc = doc
	return nil
}

// GetTitle return string
func (a *AreaDelay) GetTitle() string {
	return a.Doc.Find("h1.title").Text()
}

// GetUpdateDateText return string
func (a *AreaDelay) GetUpdateDateText() string {
	return a.Doc.Find("span.subText").Text()
}

// GetAreaDelayText return string
func (a *AreaDelay) GetAreaDelayText() string {
	const linefeed string = "\n"
	areaDelayText := []string{}

	a.Doc.Find("div.elmTblLstLine.trouble > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		s.Children()
		details := []string{}
		s.Children().Each(func(i int, ss *goquery.Selection) {
			details = append(details, ss.Text())
		})
		areaDelayText = append(areaDelayText, strings.Join(details, ", ")+linefeed)
	})

	return strings.Join(areaDelayText, linefeed)
}
