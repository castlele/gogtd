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
	tasks := this.clarifyInteractor.GetAll()
	out, err := prettyPrint(tasks)

	if err != nil {
		fmt.Fprintln(this.successOut, tasks)
	} else {
		fmt.Fprintln(this.successOut, out)
	}

	return 0
}
