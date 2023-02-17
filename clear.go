package main

import (
	"fmt"
	"github.com/jimlawless/whereami"
	"os"
	"os/exec"
)

/*
Clear history from the console.
We only need to know the current connection, not old ones.
*/
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error() + " @ " + whereami.WhereAmI())
	}
}
