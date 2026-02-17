package clarify

import (
	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/google/uuid"
)

type clarifyImpl struct {
	tasksRepo      repository.Repo[models.Task, string]
	inboxItemsRepo repository.Repo[models.InboxItem, string]
}

const (
	InboxIdNotFound = "Inbox Id not found error: %v"
)

func NewClarifyInteractor(
	tasksRepo repository.Repo[models.Task, string],
	inboxItemsRepo repository.Repo[models.InboxItem, string],
) *clarifyImpl {
	return &clarifyImpl{
		tasksRepo:      tasksRepo,
		inboxItemsRepo: inboxItemsRepo,
	}
}

func (this *clarifyImpl) GetAll() []models.Task {
	tasks, err := this.tasksRepo.List()

	if err != nil {
		return make([]models.Task, 0)
	}

	return tasks
}

func (this *clarifyImpl) AddTask(
	inboxItemId string,
	time int64,
	energy models.Energy,
	parent *models.TaskParent,
) (*models.Task, error) {
	_, err := this.inboxItemsRepo.Get(inboxItemId)

	if err != nil {
		return nil, err
	}

	inboxItem, err := this.inboxItemsRepo.Delete(inboxItemId)

	if err != nil {
		return nil, err
	}

	var copyParent models.TaskParent

	if parent != nil {
		copyParent = *parent
	} else {
		copyParent = models.NewNextTaskParent()
	}

	task := this.createTask(&inboxItem, time, energy, copyParent)

	return &task, nil
}

func (c *clarifyImpl) DeleteTask(id string) (*models.Task, error) {
	panic("unimplemented")
}

func (c *clarifyImpl) createTask(
	item *models.InboxItem,
	time int64,
	energy models.Energy,
	parent models.TaskParent,
) models.Task {
	return models.Task{
		Id:      uuid.NewString(),
		Message: item.Message,
		Time:    time,
		Energy:  energy,
		Parent:  parent,
	}
}
