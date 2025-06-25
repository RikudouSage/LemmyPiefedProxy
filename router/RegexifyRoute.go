package router

import (
	"regexp"
	"strings"
)

func RegexifyRoute(path string) (*regexp.Regexp, error) {
	if path[0] == '/' {
		path = path[1:]
	}

	regexString := "^"
	parts := strings.Split(path, "/")

	for _, part := range parts {
		if part[0] == '{' && part[len(part)-1] == '}' {
			regexString += "/([^/]+)"
		} else {
			regexString += regexp.QuoteMeta("/" + part)
		}
	}
	regexString += "$"

	return regexp.Compile(regexString)
}
