package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/clarify"
)

type toggleFavouriteCommand struct {
	id                string
	clarifyInteractor clarify.Clarify
	successOut        io.Writer
	errOut            io.Writer
}

func newToggleFavouriteCommand(
	id string,
	clarifyInteractor clarify.Clarify,
	successOut io.Writer,
	errOut io.Writer,
) *toggleFavouriteCommand {
	return &toggleFavouriteCommand{
		id:                id,
		clarifyInteractor: clarifyInteractor,
		successOut:        successOut,
		errOut:            errOut,
	}
}

func (this *toggleFavouriteCommand) Execute() int {
	task, err := this.clarifyInteractor.ToggleFavourite(this.id)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	fmt.Fprintln(this.successOut, task)

	return 0
}
