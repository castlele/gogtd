package commands

import (
	"fmt"
	"io"
)

type errorCommand struct {
	message string
	errOut  io.Writer
}

func newErrorCommand(message string, errOut io.Writer) *errorCommand {
	return &errorCommand{
		message: message,
		errOut:  errOut,
	}
}

func (this *errorCommand) Execute() int {
	fmt.Fprintln(this.errOut, this.message)
	return -1
}
