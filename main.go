package main

import (
	"fmt"
	"os"

	"github.com/castlele/gogtd/src/parsing"
)

func main() {
	cmd := parsing.ParseArguments(os.Args)

	if cmd != nil {
		os.Exit(cmd.Execute())
	}

	fmt.Fprintln(os.Stderr, "invalid command usage")
	os.Exit(-1)
}
