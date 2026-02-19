package project

import (
	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/google/uuid"
)

type projectImpl struct {
	projectsRepo repository.Repo[models.Project, string]
}

func NewProjectInteractor(
	projectsRepo repository.Repo[models.Project, string],
) *projectImpl {
	return &projectImpl{
		projectsRepo: projectsRepo,
	}
}

func (this *projectImpl) GetAll() []models.Project {
	projects, err := this.projectsRepo.List()

	if err != nil {
		return make([]models.Project, 0)
	}

	return projects
}

func (this *projectImpl) AddProject(name string) (*models.Project, error) {
	proj := this.createProject(name)

	err := this.projectsRepo.Create(proj)

	if err != nil {
		return nil, err
	}

	return &proj, nil
}

func (this *projectImpl) DeleteProject(id string) (*models.Project, error) {
	// TODO: Add cascade deletion of the tasks with force flag
	proj, err := this.projectsRepo.Delete(id)

	if err != nil {
		return nil, err
	}

	return &proj, nil
}

func (_ *projectImpl) createProject(name string) models.Project {
	return models.Project{
		Id:   uuid.NewString(),
		Name: name,
	}
}
