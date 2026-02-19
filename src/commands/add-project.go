package commands

import (
	"fmt"
	"io"

	"github.com/castlele/gogtd/src/domain/project"
)

type addProjectCommand struct {
	projectsInteractor project.Project
	successOut         io.Writer
	errOut             io.Writer
	name               string
}

func newAddProjectCommand(
	name string,
	projectsInteractor project.Project,
	successOut io.Writer,
	errOut io.Writer,
) *addProjectCommand {
	return &addProjectCommand{
		projectsInteractor: projectsInteractor,
		successOut:         successOut,
		errOut:             errOut,
		name:               name,
	}
}

func (this *addProjectCommand) Execute() int {
	proj, err := this.projectsInteractor.AddProject(this.name)

	if err != nil {
		fmt.Fprintln(this.errOut, err)
		return -1
	}

	out, err := prettyPrint(proj)

	if err != nil {
		fmt.Fprintln(this.successOut, proj)
	} else {
		fmt.Fprintln(this.successOut, out)
	}

	return 0
}
