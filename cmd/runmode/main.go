package main

import (
	"fmt"
	"os"
	"github.com/harkaitz/go-runmode"
)

var help string = `
Usage: runmode

Returns RELEASE or DEBUG based on the environment variable RELEASE_MODE.
`

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(help)
			os.Exit(0)
		} else {
			fmt.Fprintf(os.Stderr, "error: Invalid argument: %s\n", os.Args[1])
			os.Exit(1)
		}
	}
	if runmode.IsReleaseMode() {
		fmt.Println("RELEASE")
	} else {
		fmt.Println("DEBUG")
	}
}
