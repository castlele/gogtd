package commands

import (
	"fmt"
	"io"
)

type helpCommand struct {
	message string
	output  io.Writer
}

func newHelpCommand(message string, output io.Writer) *helpCommand {
	return &helpCommand{
		message: message,
		output:  output,
	}
}

func (cmd *helpCommand) Execute() int {
	fmt.Fprintln(cmd.output, cmd.message)
	return 0
}
