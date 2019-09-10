package trainDelay

import (
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TrainDelay is struct
type TrainDelay struct {
	Doc *goquery.Document
}

// Goto return error
func (t *TrainDelay) Goto(regionalCode string) error {
	doc, err := goquery.NewDocument(os.Getenv("TRAININFO_URL") + regionalCode + "/")
	if err != nil {
		return err
	}
	t.Doc = doc
	return nil
}

// GetTitle return string
func (t *TrainDelay) GetTitle() string {
	return t.Doc.Find("h1.title").Text()
}

// GetUpdateDateText return string
func (t *TrainDelay) GetUpdateDateText() string {
	return t.Doc.Find("span.subText").Text()
}

// GetTrainDelayText return string
func (t *TrainDelay) GetTrainDelayText() string {
	const linefeed string = "\n"
	trainDelayText := []string{}

	t.Doc.Find("div.elmTblLstLine.trouble > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		details := []string{}
		s.Children().Each(func(i int, ss *goquery.Selection) {
			details = append(details, ss.Text())
		})
		trainDelayText = append(trainDelayText, strings.Join(details, ", ")+linefeed)
	})

	return strings.Join(trainDelayText, linefeed)
}
