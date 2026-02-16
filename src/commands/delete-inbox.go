package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/inbox"
)

type deleteInboxCommand struct {
	id              string
	inboxInteractor inbox.Inbox
	successOut      io.Writer
	errOut          io.Writer
}

func newDeleteInboxCommand(
	id string,
	inboxInteractor inbox.Inbox,
	successOut io.Writer,
	errOut io.Writer,
) *deleteInboxCommand {
	return &deleteInboxCommand{
		id:              id,
		inboxInteractor: inboxInteractor,
		successOut:      successOut,
		errOut:          errOut,
	}
}

func (this *deleteInboxCommand) Execute() int {
	item, err := this.inboxInteractor.DeleteItem(this.id)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	fmt.Fprintln(this.successOut, item)

	return 0
}
