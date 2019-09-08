package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "strings"
)

func main() {
    doc, err := goquery.NewDocument("https://transit.yahoo.co.jp/traininfo/area/4/")
    if err != nil {
        fmt.Print("url scarapping failed")
    }

    var text string

    updateDate := doc.Find("span.subText").Text()
    text += updateDate + "\n"

    trainDelays := []string{}
    doc.Find("div.elmTblLstLine.trouble > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
        details := []string{}
        s.Children().Each(func(i int, ss *goquery.Selection) {
            details = append(details, ss.Text())
        })
        trainDelays = append(trainDelays, strings.Join(details, ", "))
    })
    text += "\n"
    text += strings.Join(trainDelays, "\n")
    fmt.Println(text)
}
