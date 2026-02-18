package commands

import (
	"errors"
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
	"github.com/castlele/gogtd/src/domain/models"
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

	fmt.Fprintln(this.successOut, task)

	return 0
}

func parseStatus(status string) (models.TaskStatus, error) {
	if status == "" {
		return models.TaskStatusPending, errors.New("You provided empty status")
	}

	switch status {
	case "pending":
		return models.TaskStatusPending, nil
	case "in_progress":
		return models.TaskStatusInProgress, nil
	case "done":
		return models.TaskStatusDone, nil
	default:
		return models.TaskStatusPending, fmt.Errorf(
			"Invalid status provided. Expected one of: pending, in_progress, done. Got: %v",
			status,
		)
	}
}
