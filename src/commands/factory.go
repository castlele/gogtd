package commands

import (
	"io"

	"github.com/castlele/gogtd/src/domain/inbox"
)

type CommandsFactory interface {
	Inbox() Command
	AddInbox(message string) Command
	DeleteInbox(id string) Command

	Help(message string) Command

	Error(message string) Command
}

type commandsFactoryImpl struct {
	inboxInteractor inbox.Inbox
	successOut      io.Writer
	errOut          io.Writer
}

func NewCommandsFactory(
	inboxInteractor inbox.Inbox,
	successOut io.Writer,
	errOut io.Writer,
) CommandsFactory {
	return &commandsFactoryImpl{
		inboxInteractor: inboxInteractor,
		successOut:      successOut,
		errOut:          errOut,
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

func (this *commandsFactoryImpl) Help(message string) Command {
	return newHelpCommand(message, this.successOut)
}

func (this *commandsFactoryImpl) Error(message string) Command {
	return newErrorCommand(message, this.errOut)
}
