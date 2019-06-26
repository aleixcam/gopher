package parser

import (
	"fmt"
	"regexp"
)

const (
	baseURL    = "https://fogo-parser.dev/"
	hashTagURL = "https://fogo-parser.dev/hash/"
)

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}

	return s[:0]
}

func createUrlAnchor(src []byte) []byte {
	url := string(src)
	return []byte(fmt.Sprintf("<a href=\"%s\">%s</a>", url, url))
}

func createTagAnchor(src []byte) []byte {
	tag := string(src)
	return []byte(fmt.Sprintf("<a href=\"%s%s\">%s</a>", baseURL, trimLeftChar(tag), tag))
}

func createHashAnchor(src []byte) []byte {
	hash := string(src)
	return []byte(fmt.Sprintf("<a href=\"%s%s\">%s</a>", hashTagURL, trimLeftChar(hash), hash))
}

// Parse fomating the text into a valid output text
func Parse(msg string) string {
	urlRegex := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`)
	tagRegex := regexp.MustCompile(`(?:@)([a-zA-Z\d]+)?`)
	hashRegex := regexp.MustCompile(`(?:#)([a-zA-Z\d]+)?`)

	in := []byte(msg)
	out := urlRegex.ReplaceAllFunc(in, createUrlAnchor)
	out = tagRegex.ReplaceAllFunc(out, createTagAnchor)
	out = hashRegex.ReplaceAllFunc(out, createHashAnchor)

	return string(out)
}
