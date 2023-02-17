package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/jimlawless/whereami"
)

/*
Get a list of all the cities in a given country.
*/
func getCities(country string) []string {
	args := []string{"cities"}
	args = append(args, country)
	cmd := exec.Command("nordvpn", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
	}

	return sanitizeCities(out.String())
}

func sanitizeCities(s string) []string {
	parts := strings.Split(s, "\n")
	parts[1] = nonAlphanumericRegexReplaceAll(parts[1])
	cArr := strings.Split(parts[1], " ")
	var filtered []string
	for _, v := range cArr {
		if contentMinLength < len(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
