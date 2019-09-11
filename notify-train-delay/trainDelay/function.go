package trainDelay

import (
	"regexp"

	"../template"
)

func isArea(code string) bool {
	rep := regexp.MustCompile(`^[0-9]*$`)
	return rep.MatchString(code)
}

// TrainDelayHandle return string
func TrainDelayHandle(code string) string {
	if isArea(code) {
		return getAreaTrainDelayText(code)
	}
	return getRouteTrainDelayText(code)
}

func getRouteTrainDelayText(routeCode string) string {
	routeDelay := RouteDelay{}

	err := routeDelay.Goto(routeCode)
	if err != nil {
		return template.Information
	}

	routeTrainDelayText := routeDelay.GetRouteInfo()
	if len(routeTrainDelayText) == 0 {
		return template.Information
	}

	return routeTrainDelayText
}

func getAreaTrainDelayText(areaCode string) string {
	areaDelay := AreaDelay{}

	err := areaDelay.Goto(areaCode)
	if err != nil {
		return template.Information
	}

	title := areaDelay.GetTitle()
	updateDateText := areaDelay.GetUpdateDateText()

	if len(title) == 0 || len(updateDateText) == 0 {
		return template.Information
	}

	trainDelayText := areaDelay.GetAreaDelayText()
	if len(trainDelayText) == 0 {
		return Serialize([]string{title, updateDateText, template.NotDelay}...)
	}

	return Serialize([]string{title, updateDateText, trainDelayText}...)
}
