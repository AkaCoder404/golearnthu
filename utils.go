package gothulearn

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func stripHtml(htmlString string) string {
	p := bluemonday.StripTagsPolicy()
	html := p.Sanitize(htmlString)

	// Trim leading / trailing spaces
	finalString := strings.TrimSpace(html)
	return finalString
}
