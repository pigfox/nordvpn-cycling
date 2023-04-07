package main

import (
	"regexp"
	"strings"
)

func nonAlphanumericRegexReplaceAll(str string) string {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z_ ]+`)
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func sanitize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, ",", "")
	return s
}
