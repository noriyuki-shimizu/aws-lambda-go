package trainDelay

import (
	"strings"

	"../template"
)

// GetTrainDelayText return string
func GetTrainDelayText(regionalCode string) string {
	trainDelay := TrainDelay{}

	err := trainDelay.Goto(regionalCode)
	if err != nil {
		return template.Information
	}

	updateDateText := trainDelay.GetUpdateDateText()
	if len(updateDateText) == 0 {
		return template.Information
	}

	trainDelayText := trainDelay.GetTrainDelayText()
	if len(trainDelayText) == 0 {
		return serialize([]string{updateDateText, template.NotDelay}...)
	}

	return serialize([]string{updateDateText, trainDelayText}...)
}

func serialize(texts ...string) string {
	result := []string{}
	for _, text := range texts {
		result = append(result, text)
	}
	return strings.Join(result, "\n")
}
