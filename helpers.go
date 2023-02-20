package main

import "regexp"

func nonAlphanumericRegexReplaceAll(str string) string {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z_ ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}
