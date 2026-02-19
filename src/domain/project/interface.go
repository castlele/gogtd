package project

import "github.com/castlele/gogtd/src/domain/models"

type Project interface {
	GetAll() []models.Project
	AddProject(name string) (*models.Project, error)
	DeleteProject(id string) (*models.Project, error)

	// TODO: Implement later
	// AddStep(id string, message string) (*models.ProjectStep, error)
}
