package main

import (
	"bytes"
	"fmt"
	"github.com/jimlawless/whereami"
	"os/exec"
	"strings"
)

/*
Get a list of all the cities in a given country.
*/
func getCities() {
	for _, country := range availableData.Countries {
		args := []string{"cities"}
		args = append(args, country)
		cmd := exec.Command("nordvpn", args...)

		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
		}
		cities := sanitize(out.String())

		findCities(cities, country)
	}
}

func findCities(cities string, country string) {
	var tmpCities []string
	for _, city := range potentialData.Cities[country] {
		if strings.Contains(cities, city) {
			tmpCities = append(tmpCities, city)
		}
	}
	availableData.Cities[country] = tmpCities
}
