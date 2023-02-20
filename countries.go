package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/jimlawless/whereami"
)

/*
Get a list of all countries.
*/
func getCountries() []string {
	cmd := exec.Command("nordvpn", "countries")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
	}

	return sanitizeCountry(out.String())
}

func sanitizeCountry(s string) []string {
	parts := strings.Split(s, "\n")
	parts[1] = nonAlphanumericRegexReplaceAll(parts[1])
	cArr := strings.Split(parts[1], " ")
	var filtered []string
	for _, v := range cArr {
		if config.ContentMinLength < len(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
