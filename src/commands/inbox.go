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
	items := this.inboxInteractor.GetAll()
	out, err := prettyPrint(items)

	if err != nil {
		fmt.Fprintln(this.successOut, items)
	} else {
		fmt.Fprintln(this.successOut, out)
	}

	return 0
}
