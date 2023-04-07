package main

import (
	"bytes"
	"fmt"
	"github.com/jimlawless/whereami"
	"os/exec"
	"strings"
)

/*
Get the status of a connection.
*/
func status() bool {
	cmd := exec.Command("nordvpn", "status")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
	}
	if strings.Contains(out.String(), "Status: Connected") {
		fmt.Print(out.String())
		return true
	}
	return false
}
