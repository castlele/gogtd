package commands

import (
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
	"github.com/castlele/gogtd/src/domain/inbox"
	"github.com/castlele/gogtd/src/domain/project"
)

type CommandsFactory interface {
	Inbox() Command
	AddInbox(message string) Command
	DeleteInbox(id string) Command

	Tasks(status string) Command
	AddTaskFromInbox(
		id string,
		time int64,
		energy string,
		parent string,
	) Command
	AddTask(
		message string,
		time int64,
		energy string,
		parent string,
	) Command
	DeleteTask(id string) Command
	ToggleFavourite(id string) Command
	SetStatus(id string, status string) Command

	Projects() Command
	AddProject(name string) Command
	DeleteProject(id string) Command

	Help(message string) Command

	Error(message string) Command
}

type commandsFactoryImpl struct {
	inboxInteractor    inbox.Inbox
	clarifyInteractor  clarify.Clarify
	projectsInteractor project.Project
	successOut         io.Writer
	errOut             io.Writer
}

func NewCommandsFactory(
	inboxInteractor inbox.Inbox,
	clarifyInteractor clarify.Clarify,
	projectsInteractor project.Project,
	successOut io.Writer,
	errOut io.Writer,
) CommandsFactory {
	return &commandsFactoryImpl{
		inboxInteractor:    inboxInteractor,
		clarifyInteractor:  clarifyInteractor,
		projectsInteractor: projectsInteractor,
		successOut:         successOut,
		errOut:             errOut,
	}
}

func (this *commandsFactoryImpl) Inbox() Command {
	return newInboxCommand(this.inboxInteractor, this.successOut)
}

func (this *commandsFactoryImpl) AddInbox(message string) Command {
	return newAddInboxCommand(
		message,
		this.inboxInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) DeleteInbox(id string) Command {

	return newDeleteInboxCommand(
		id,
		this.inboxInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) Tasks(status string) Command {
	return newTasksCommand(
		status,
		this.clarifyInteractor,
		this.successOut,
	)
}

func (this *commandsFactoryImpl) AddTaskFromInbox(
	id string,
	time int64,
	energy string,
	parent string,
) Command {
	return newAddFromInboxTaskCommand(
		id,
		time,
		energy,
		parent,
		this.clarifyInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) AddTask(
	message string,
	time int64,
	energy string,
	parent string,
) Command {
	return newCreateTaskCommand(
		message,
		time,
		energy,
		parent,
		this.clarifyInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) DeleteTask(id string) Command {
	return newDeleteTaskCommand(
		id,
		this.clarifyInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) ToggleFavourite(id string) Command {
	return newToggleFavouriteCommand(
		id,
		this.clarifyInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) SetStatus(id string, status string) Command {
	return newSetStatusCommand(
		id,
		status,
		this.clarifyInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) Projects() Command {
	return newProjectsCommand(
		this.projectsInteractor,
		this.successOut,
	)
}

func (this *commandsFactoryImpl) AddProject(name string) Command {
	return newAddProjectCommand(
		name,
		this.projectsInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) DeleteProject(id string) Command {
	return newDeleteProjectCommand(
		id,
		this.projectsInteractor,
		this.successOut,
		this.errOut,
	)
}

func (this *commandsFactoryImpl) Help(message string) Command {
	return newHelpCommand(message, this.successOut)
}

func (this *commandsFactoryImpl) Error(message string) Command {
	return newErrorCommand(message, this.errOut)
}
