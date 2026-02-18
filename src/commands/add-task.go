package commands

import (
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
	parent            string
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
	parent string,
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
			parent:            parent,
		},
		id: id,
	}
}

func newCreateTaskCommand(
	message string,
	time int64,
	energy string,
	parent string,
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
			parent:            parent,
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

	parent, err := getParent(this.parent)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	task, err := this.clarifyInteractor.AddTask(
		this.message,
		this.time,
		energy,
		parent,
	)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	fmt.Fprintln(this.successOut, task)

	return 0
}

func getEnergyModel(energy string) (models.Energy, error) {
	lowerEnergy := strings.ToLower(energy)

	switch lowerEnergy {
	case "low":
		return models.EnergyLow, nil
	case "mid":
		return models.EnergyMid, nil
	case "high":
		return models.EnergyHigh, nil
	default:
		return models.EnergyLow, fmt.Errorf(
			"Invalid energy passed: %v. Expected one of: low, mid, high",
			lowerEnergy,
		)
	}
}

func getParent(parent string) (*models.TaskParent, error) {
	if parent == "" {
		return nil, nil
	}

	comps := strings.Split(parent, "::")

	if len(comps) != 2 {
		return nil, fmt.Errorf(
			"Invalid format of the parent param. Has to be id::type, but got: %v",
			parent,
		)
	}

	var taskType models.TaskParentType
	var id string

	switch comps[1] {
	case "box":
		taskType = models.BoxParentType
		id = parseBoxType(comps[0]).String()
	case "project":
		taskType = models.ProjectParentType
		panic("Can't get id for project type")
	case "step":
		taskType = models.StepParentType
		panic("Can't get id for step type")
	default:
		return nil, fmt.Errorf(
			"Invalid task type received. Expected one of: box, project, step. But got: %v",
			comps[1],
		)
	}

	taskParent := models.TaskParent{
		Id:   id,
		Type: taskType,
	}

	return &taskParent, nil
}

func parseBoxType(id string) models.BoxType {
	switch id {
	case "next":
		return models.BoxTypeNext
	case "waiting":
		return models.BoxTypeWaiting
	case "someday_maybe":
		return models.BoxTypeSomedayMaybe
	default:
		return models.BoxTypeNext
	}
}
