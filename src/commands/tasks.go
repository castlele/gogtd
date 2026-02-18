package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
)

type tasksCommand struct {
	clarifyInteractor clarify.Clarify
	successOut        io.Writer
}

func newTasksCommand(
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,

) *tasksCommand {
	return &tasksCommand{
		clarifyInteractor: clarifyInteractor,
		successOut:        successOut,
	}
}

func (this *tasksCommand) Execute() int {
	fmt.Fprintln(this.successOut, this.clarifyInteractor.GetAll())
	return 0
}
