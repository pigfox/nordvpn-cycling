package main

import "regexp"

var contentMinLength = 2

func nonAlphanumericRegexReplaceAll(str string) string {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z_ ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
