package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/inbox"
)

type addInboxCommand struct {
	message         string
	inboxInteractor inbox.Inbox
	successOut      io.Writer
	errOut          io.Writer
}

func newAddInboxCommand(
	message string,
	inboxInteractor inbox.Inbox,
	successOut io.Writer,
	errOut io.Writer,
) *addInboxCommand {
	return &addInboxCommand{
		message:         message,
		inboxInteractor: inboxInteractor,
		successOut:      successOut,
		errOut:          errOut,
	}
}

func (this *addInboxCommand) Execute() int {
	item, err := this.inboxInteractor.AddItem(this.message)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	fmt.Fprintln(this.successOut, item)

	return 0
}
