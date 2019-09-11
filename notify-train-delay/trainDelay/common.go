package trainDelay

import "strings"

// Serialize return string
func Serialize(texts ...string) string {
	result := []string{}
	for _, text := range texts {
		result = append(result, text)
	}
	return strings.Join(result, "\n")
}

// ConvNewline return string
func ConvNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}
