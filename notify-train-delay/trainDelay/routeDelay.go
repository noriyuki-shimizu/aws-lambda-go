package trainDelay

import (
	"os"
	"strings"

	"../template"
	"github.com/PuerkitoBio/goquery"
)

// RouteDelay is struct
type RouteDelay struct {
	Doc *goquery.Document
}

// Goto return error
func (r *RouteDelay) Goto(routeCode string) error {
	accessURL := os.Getenv("TRAININFO_YAHOO_MAIN_URL")
	accessURL += os.Getenv("TRAININFO_YAHOO_CONTEXT")
	accessURL += os.Getenv("TRAININFO_YAHOO_SUB_ROUTE_QUERY")
	accessURL += routeCode

	doc, err := goquery.NewDocument(accessURL)
	if err != nil {
		return err
	}
	r.Doc = doc
	return nil
}

// GetRouteInfo return []string
func (r *RouteDelay) GetRouteInfo() string {
	routeLinks := []string{}
	r.Doc.Find("ul.elmSearchItem.double > li").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Children().Attr("href")
		routeLinks = append(routeLinks, url)
	})

	routeInfo := []string{}
	for _, routeLink := range routeLinks {
		doc, err := goquery.NewDocument(os.Getenv("TRAININFO_YAHOO_MAIN_URL") + routeLink)
		if err != nil {
			return template.Information
		}

		title := doc.Find("div.labelLarge > h1.title").Text()

		serviceStatus := []string{}
		doc.Find("div#mdServiceStatus > dl").Children().Each(func(_ int, s *goquery.Selection) {
			serviceStatus = append(serviceStatus, ConvNewline(s.Text(), ""))
		})
		routeInfo = append(routeInfo, Serialize([]string{title, strings.Join(serviceStatus, "\n") + "\n"}...))
	}
	return strings.Join(routeInfo, "\n")
}
