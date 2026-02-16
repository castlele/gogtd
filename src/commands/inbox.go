package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/inbox"
)

type inboxCommand struct {
	inboxInteractor inbox.Inbox
	successOut      io.Writer
}

func newInboxCommand(
	inboxInteractor inbox.Inbox,
	successOut io.Writer,
) *inboxCommand {
	return &inboxCommand{
		inboxInteractor: inboxInteractor,
		successOut:      successOut,
	}
}

func (this *inboxCommand) Execute() int {
	fmt.Fprintln(this.successOut, this.inboxInteractor.GetAll())
	return 0
}
