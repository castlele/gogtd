package main

import (
	"fmt"
	"os"

	"github.com/castlele/gogtd/src/commands"
	"github.com/castlele/gogtd/src/config"
	"github.com/castlele/gogtd/src/domain/inbox"
	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/castlele/gogtd/src/parsing"
)

func main() {
	conf, err := config.LoadConfig(config.DefaultConfigFilePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	inboxRepo, err := repository.NewFPRepo(
		conf.GetInboxPath(),
		func(inbox models.InboxItem) string {
			return inbox.Id
		},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	inboxInteractor := inbox.NewInboxInteractor(inboxRepo)

	factory := commands.NewCommandsFactory(
		inboxInteractor,
		os.Stdout,
		os.Stderr,
	)
	cmd := parsing.ParseArguments(os.Args, factory)

	if cmd != nil {
		os.Exit(cmd.Execute())
	}

	fmt.Fprintln(os.Stderr, "invalid command usage")
	os.Exit(-1)
}
