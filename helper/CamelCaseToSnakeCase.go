package helper

import (
	"regexp"
	"strings"
)

func ToSnakeCase(input string) string {
	re := regexp.MustCompile("(.)([A-Z][a-z]+)")
	input = re.ReplaceAllString(input, "${1}_${2}")

	re = regexp.MustCompile("([a-z0-9])([A-Z])")
	input = re.ReplaceAllString(input, "${1}_${2}")

	return strings.ToLower(input)
}
