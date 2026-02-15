package commands

import "fmt"

type inboxCommand struct {
}

func NewInboxCommand() inboxCommand {
	return inboxCommand{}
}

func (_ inboxCommand) Execute() int {
	fmt.Println("Tasks")
	return 0
}
