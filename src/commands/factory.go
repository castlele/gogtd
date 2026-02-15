package commands

type CommandsFactory interface {
	Inbox() Command
	Help(message string) Command
}

type commandsFactoryImpl struct {
}

func NewCommandsFactory() CommandsFactory {
	return &commandsFactoryImpl{}
}

func (this *commandsFactoryImpl) Inbox() Command {
	return newInboxCommand()
}

func (this *commandsFactoryImpl) Help(message string) Command {
	return newHelpCommand(message)
}
