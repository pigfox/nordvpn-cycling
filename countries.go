package main

import (
	"bytes"
	"fmt"
	"github.com/jimlawless/whereami"
	"os/exec"
	"strings"
)

/*
Get a list of all countries.
*/
func getCountries() {
	cmd := exec.Command("nordvpn", "countries")
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
	}

	findCountries(sanitize(out.String()))
}

func findCountries(input string) {
	for _, country := range potentialData.Countries {
		if strings.Contains(input, country) {
			availableData.Countries = append(availableData.Countries, country)
		}
	}
}
