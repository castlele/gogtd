package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
)

type deleteTaskCommand struct {
	id              string
	inboxInteractor clarify.Clarify
	successOut      io.Writer
	errOut          io.Writer
}

func newDeleteTaskCommand(
	id string,
	inboxInteractor clarify.Clarify,
	successOut io.Writer,
	errOut io.Writer,
) *deleteTaskCommand {
	return &deleteTaskCommand{
		id:              id,
		inboxInteractor: inboxInteractor,
		successOut:      successOut,
		errOut:          errOut,
	}
}

func (this *deleteTaskCommand) Execute() int {
	item, err := this.inboxInteractor.DeleteTask(this.id)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	fmt.Fprintln(this.successOut, item)

	return 0
}
