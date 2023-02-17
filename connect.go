package main

import (
	"bytes"
	"fmt"
	"github.com/jimlawless/whereami"
	"os/exec"
	"strings"
)

/*
Connect to a given city.
*/
func connect(city string) {
	args := []string{"connect", city}
	cmd := exec.Command("nordvpn", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
	}
	parts := strings.Split(out.String(), "\n")
	fmt.Println(parts[2])
}
