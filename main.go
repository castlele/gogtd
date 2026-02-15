package main

import (
	"fmt"
	"os"

	"github.com/castlele/gogtd/src/commands"
	"github.com/castlele/gogtd/src/parsing"
)

func main() {
	factory := commands.NewCommandsFactory()
	cmd := parsing.ParseArguments(os.Args, factory)

	if cmd != nil {
		os.Exit(cmd.Execute())
	}

	fmt.Fprintln(os.Stderr, "invalid command usage")
	os.Exit(-1)
}
