package commands

import (
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
	"github.com/castlele/gogtd/src/domain/inbox"
)

type CommandsFactory interface {
	Inbox() Command
	AddInbox(message string) Command
	DeleteInbox(id string) Command

	Tasks() Command
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

	Help(message string) Command

	Error(message string) Command
}

type commandsFactoryImpl struct {
	inboxInteractor   inbox.Inbox
	clarifyInteractor clarify.Clarify
	successOut        io.Writer
	errOut            io.Writer
}

func NewCommandsFactory(
	inboxInteractor inbox.Inbox,
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,
	errOut io.Writer,
) CommandsFactory {
	return &commandsFactoryImpl{
		inboxInteractor:   inboxInteractor,
		clarifyInteractor: clarifyInteractor,
		successOut:        successOut,
		errOut:            errOut,
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

func (this *commandsFactoryImpl) Tasks() Command {
	return newTasksCommand(
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

func (this *commandsFactoryImpl) Help(message string) Command {
	return newHelpCommand(message, this.successOut)
}

func (this *commandsFactoryImpl) Error(message string) Command {
	return newErrorCommand(message, this.errOut)
}
