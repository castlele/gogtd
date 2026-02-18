package commands

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/castlele/gogtd/src/domain/clarify"
	"github.com/castlele/gogtd/src/domain/models"
)

type addTaskCommand struct {
	clarifyInteractor clarify.Clarify
	successOut        io.Writer
	errOut            io.Writer
	time              int64
	energy            string
}

type addFromInboxTaskCommand struct {
	*addTaskCommand

	id string
}

type createTaskCommand struct {
	*addTaskCommand

	message string
}

func newAddFromInboxTaskCommand(
	id string,
	time int64,
	energy string,
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,
	errOut io.Writer,
) *addFromInboxTaskCommand {
	return &addFromInboxTaskCommand{
		addTaskCommand: &addTaskCommand{
			clarifyInteractor: clarifyInteractor,
			successOut:        successOut,
			errOut:            errOut,
			time:              time,
			energy:            energy,
		},
		id: id,
	}
}

func newCreateTaskCommand(
	message string,
	time int64,
	energy string,
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,
	errOut io.Writer,
) *createTaskCommand {
	return &createTaskCommand{
		addTaskCommand: &addTaskCommand{
			clarifyInteractor: clarifyInteractor,
			successOut:        successOut,
			errOut:            errOut,
			time:              time,
			energy:            energy,
		},
		message: message,
	}
}

func (this *addFromInboxTaskCommand) Execute() int {
	energy, err := getEnergyModel(this.energy)

	if err != nil {
		return -1
	}

	task, err := this.clarifyInteractor.ConvertToTask(
		this.id,
		this.time,
		energy,
		nil,
	)

	if err != nil {
		return -1
	}

	fmt.Fprintln(this.successOut, task)

	return 0
}

func (this *createTaskCommand) Execute() int {
	energy, err := getEnergyModel(this.energy)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	task, err := this.clarifyInteractor.AddTask(
		this.message,
		this.time,
		energy,
		nil,
	)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	fmt.Fprintln(this.successOut, task)

	return 0
}

func getEnergyModel(energy string) (models.Energy, error) {
	lenergy := strings.ToLower(energy)

	switch lenergy {
	case "low":
		return models.EnergyLow, nil
	case "mid":
		return models.EnergyMid, nil
	case "high":
		return models.EnergyHigh, nil
	default:
		return models.EnergyLow, errors.New(
			fmt.Sprintf(
				"Invalid energy passed: %v. Expected one of: low, mid, high",
				lenergy,
			),
		)
	}
}
