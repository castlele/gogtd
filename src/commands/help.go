package commands

import (
	"fmt"
	"os"
)

type helpCommand struct {
	message string
}

func newHelpCommand(message string) *helpCommand {
	return &helpCommand{
		message: message,
	}
}

func (cmd *helpCommand) Execute() int {
	fmt.Fprintln(os.Stdout, cmd.message)
	return 0
}
