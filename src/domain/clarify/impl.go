package clarify

import (
	"slices"

	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/google/uuid"
)

type clarifyImpl struct {
	tasksRepo      repository.Repo[models.Task, string]
	doneTasksRepo  repository.Repo[models.Task, string]
	inboxItemsRepo repository.Repo[models.InboxItem, string]
}

const (
	InboxIdNotFound = "Inbox Id not found error: %v"
)

func NewClarifyInteractor(
	tasksRepo repository.Repo[models.Task, string],
	doneTasksRepo repository.Repo[models.Task, string],
	inboxItemsRepo repository.Repo[models.InboxItem, string],
) *clarifyImpl {
	return &clarifyImpl{
		tasksRepo:      tasksRepo,
		doneTasksRepo:  doneTasksRepo,
		inboxItemsRepo: inboxItemsRepo,
	}
}

func (this *clarifyImpl) GetAll(
	status []models.TaskStatus,
) []models.Task {
	tasks, err := this.tasksRepo.List()

	if err != nil {
		return make([]models.Task, 0)
	}

	if len(status) == 0 {
		status = []models.TaskStatus{models.TaskStatusPending, models.TaskStatusInProgress}
	}

	doneTasks, err := this.doneTasksRepo.List()

	if err == nil && len(doneTasks) > 0 {
		tasks = append(tasks, doneTasks...)
	}

	return slices.DeleteFunc(tasks, func(task models.Task) bool {
		return !slices.Contains(status, task.Status)
	})
}

func (this *clarifyImpl) ConvertToTask(
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

	return this.AddTask(inboxItem.Message, time, energy, parent)
}

func (this *clarifyImpl) AddTask(
	message string,
	time int64,
	energy models.Energy,
	parent *models.TaskParent,
) (*models.Task, error) {
	var copyParent models.TaskParent

	if parent != nil {
		copyParent = *parent
	} else {
		copyParent = models.NewNextTaskParent()
	}

	task := this.createTask(message, time, energy, copyParent)
	err := this.tasksRepo.Create(task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (this *clarifyImpl) DeleteTask(id string) (*models.Task, error) {
	task, err := this.tasksRepo.Delete(id)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (this *clarifyImpl) ToggleFavourite(id string) (*models.Task, error) {
	task, err := this.tasksRepo.Get(id)

	if err != nil {
		return nil, err
	}

	task.Favourite = !task.Favourite

	err = this.tasksRepo.Update(task)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (this *clarifyImpl) SetStatus(
	id string,
	status models.TaskStatus,
) (*models.Task, error) {
	task, err := this.tasksRepo.Get(id)

	if err != nil {
		return nil, err
	}

	prevStatus := task.Status
	task.Status = status

	if status == models.TaskStatusDone {
		this.migrateTaskToDone(task)

	} else {
		if prevStatus == models.TaskStatusDone {
			this.migrateTaskFromDone(task)
		} else {
			err = this.tasksRepo.Update(task)

			if err != nil {
				return nil, err
			}
		}
	}

	return &task, nil
}

func (this *clarifyImpl) migrateTaskToDone(task models.Task) error {
	_, err := this.tasksRepo.Delete(task.Id)

	if err != nil {
		return err
	}

	err = this.doneTasksRepo.Create(task)

	if err != nil {
		return err
	}

	return nil
}

func (this *clarifyImpl) migrateTaskFromDone(task models.Task) error {
	_, err := this.doneTasksRepo.Delete(task.Id)

	if err != nil {
		return err
	}

	err = this.tasksRepo.Create(task)

	if err != nil {
		return err
	}

	return nil
}

func (_ *clarifyImpl) createTask(
	message string,
	time int64,
	energy models.Energy,
	parent models.TaskParent,
) models.Task {
	return models.Task{
		Id:        uuid.NewString(),
		Message:   message,
		Time:      time,
		Energy:    energy,
		Parent:    parent,
		Status:    models.TaskStatusPending,
		Favourite: false,
	}
}
