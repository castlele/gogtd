package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
)

type setStatusCommand struct {
	id                string
	status            string
	clarifyInteractor clarify.Clarify
	successOut        io.Writer
	errOut            io.Writer
}

func newSetStatusCommand(
	id string,
	status string,
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,
	errOut io.Writer,
) *setStatusCommand {
	return &setStatusCommand{
		id:                id,
		status:            status,
		clarifyInteractor: clarifyInteractor,
		successOut:        successOut,
		errOut:            errOut,
	}
}

func (this *setStatusCommand) Execute() int {
	status, err := parseStatus(this.status)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	task, err := this.clarifyInteractor.SetStatus(this.id, status)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	out, err := prettyPrint(task)

	if err != nil {
		fmt.Fprintln(this.successOut, task)
	} else {
		fmt.Fprintln(this.successOut, out)
	}

	return 0
}
