package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetTrainDelayText() string {
	doc, err := goquery.NewDocument(os.Getenv("TRAININFO_URL"))
	if err != nil {
		fmt.Print("url scarapping failed")
	}

	var text string
	const linefeed string = "\n"

	updateDate := doc.Find("span.subText").Text()
	text += updateDate + "\n"

	trainDelays := []string{}
	doc.Find("div.elmTblLstLine.trouble > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		details := []string{}
		s.Children().Each(func(i int, ss *goquery.Selection) {
			details = append(details, ss.Text())
		})
		trainDelays = append(trainDelays, strings.Join(details, ", "))
		trainDelays = append(trainDelays, "\n")
	})
	text += "\n"
	text += strings.Join(trainDelays, "\n")
	return text
}
