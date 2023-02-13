package gothulearn

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func stripHtml(htmlString string) string {
	p := bluemonday.StripTagsPolicy()
	html := p.Sanitize(htmlString)

	// html = strings.ReplaceAll(html, "\r ", "")
	// html = strings.ReplaceAll(html, "\r\n", "")

	// Remove all newline at beggining and end
	startEndNewlines := strings.Split(html, "\n")
	finalString := ""
	for index := 1; index < len(startEndNewlines)-1; index++ {
		finalString += (startEndNewlines[index] + "\n")
	}

	// Trim leading / trailing spaces
	finalString = strings.TrimSpace(finalString)

	// html = strings.ReplaceAll(html, "\n", "")
	return finalString
}
