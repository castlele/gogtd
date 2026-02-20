package commands

import (
	"fmt"
	"io"
	"strings"

	"github.com/castlele/gogtd/src/domain/clarify"
	"github.com/castlele/gogtd/src/domain/models"
)

type tasksCommand struct {
	status            string
	clarifyInteractor clarify.Clarify
	successOut        io.Writer
}

func newTasksCommand(
	status string,
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,
) *tasksCommand {
	return &tasksCommand{
		status:            status,
		clarifyInteractor: clarifyInteractor,
		successOut:        successOut,
	}
}

func (this *tasksCommand) Execute() int {
	statuses := this.parseStatuses()
	tasks := this.clarifyInteractor.GetAll(statuses)

	out, err := prettyPrint(tasks)

	if err != nil {
		fmt.Fprintln(this.successOut, tasks)
	} else {
		fmt.Fprintln(this.successOut, out)
	}

	return 0
}

func (this *tasksCommand) parseStatuses() []models.TaskStatus {
	if this.status == "" {
		return nil
	}

	comps := strings.Split(this.status, ",")
	var statuses []models.TaskStatus

	for _, comp := range comps {
		status, err := parseStatus(comp)

		if err == nil {
			statuses = append(statuses, status)
		}
	}

	return statuses
}
