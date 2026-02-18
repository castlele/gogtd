package clarify

import (
	"errors"
	"testing"

	"github.com/castlele/gogtd/src/domain/models"
	"github.com/castlele/gogtd/src/domain/repository"
	"github.com/castlele/gogtd/src/utils"
	"github.com/google/uuid"
)

const (
	storageFp   = "./gdt"
	inboxFp     = storageFp + "/inbox.json"
	tasksFp     = storageFp + "/tasks.json"
	doneTasksFp = storageFp + "/done_tasks.json"
)

func TestGetAll(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN empty repository WHEN getting all tasks THEN returning empty tasks list",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			tasks := sut.GetAll()

			if len(tasks) > 0 {
				t.Errorf(
					"Invalid amount of tasks in the repo. Expected 0, but got: %v. Tasks: %v",
					len(tasks),
					tasks,
				)
			}
		},
	)

	t.Run(
		"GIVEN non empty repository WHEN getting all tasks THEN returning the same tasks as in repo",
		func(t *testing.T) {
			utils.Delete(storageFp)
			task := models.Task{
				Id:      uuid.NewString(),
				Message: "Hello, World",
				Time:    900_000,
				Energy:  models.EnergyLow,
			}
			sut := createInteractor()
			sut.tasksRepo.Create(task)

			tasks := sut.GetAll()

			if len(tasks) != 1 {
				t.Errorf("Invalid amount of tasks. Expected: 1, got: %v", len(tasks))
				return
			}

			if tasks[0] != task {
				t.Errorf("Invalid task got. Expected: %v, got: %v", tasks[0], task)
				return
			}
		},
	)
}

func TestConvertToTask(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN empty inbox repo WHEN adding new task THEN error is thrown",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			task, err := sut.ConvertToTask("random", 0, models.EnergyLow, nil)

			if task != nil {
				t.Errorf("Unexpected task. Expecting nil, but got: %v", task)
				return
			}

			if err == nil {
				t.Errorf("Error is nil, while expecting the an error")
				return
			}

			if !errors.Is(err, repository.ErrNotFound) {
				t.Errorf("Error of invalid type. Expected ErrNotFound, got: %v", err)
				return
			}
		},
	)

	t.Run(
		"GIVEN none empy inbox repo WHEN adding new task THEN task is created and inbox item is deleted",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()
			id := "random_id"
			msg := "Super puper hello world"
			sut.inboxItemsRepo.Create(models.InboxItem{
				Id:      id,
				Message: msg,
			})

			task, err := sut.ConvertToTask(id, 0, models.EnergyHigh, nil)

			if err != nil {
				panic(err)
			}

			if task.Message != msg {
				t.Errorf("Invalid task message. Expecting: %v, got: %v", msg, task.Message)
				return
			}

			if task.Energy != models.EnergyHigh {
				t.Errorf("Invalid energy level. Expected: %v, got: %v", models.EnergyHigh, task.Energy)
				return
			}
		},
	)

	t.Run(
		"GIVEN none empy inbox repo WHEN adding new task with energy THEN task is created with energy",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()
			id := "random_id"
			msg := "Super puper hello world"
			sut.inboxItemsRepo.Create(models.InboxItem{
				Id:      id,
				Message: msg,
			})

			task, err := sut.ConvertToTask(id, 0, models.EnergyLow, nil)

			if err != nil {
				t.Errorf("Got unexpected error: %v", err)
				return
			}

			if task == nil {
				t.Errorf("Task is nil. Expecting task with message: %v", msg)
				return
			}

			if task.Message != msg {
				t.Errorf("Invalid task message. Expecting: %v, got: %v", msg, task.Message)
				return
			}

			items, err := sut.inboxItemsRepo.List()

			if err != nil {
				panic(err)
			}

			if len(items) != 0 {
				t.Errorf("Item wasn't deleted from repo. Expected 0, but got: %v", len(items))
				return
			}
		},
	)

	t.Run(
		"GIVEN none empy inbox repo WHEN adding new task with time THEN task is created with time",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()
			id := "random_id"
			msg := "Super puper hello world"
			sut.inboxItemsRepo.Create(models.InboxItem{
				Id:      id,
				Message: msg,
			})
			time := int64(60 * 1000)

			task, err := sut.ConvertToTask(id, time, models.EnergyLow, nil)

			if err != nil {
				t.Errorf("Got unexpected error: %v", err)
				return
			}

			if task == nil {
				t.Errorf("Task is nil. Expecting task with message: %v", msg)
				return
			}

			if task.Message != msg {
				t.Errorf("Invalid task message. Expecting: %v, got: %v", msg, task.Message)
				return
			}

			if task.Time != time {
				t.Errorf("Invalid time value. Expected: %v, got: %v", time, task.Time)
				return
			}
		},
	)

	t.Run(
		"GIVEN new task added WHEN no parent specified THEN task is added to the BoxTypeNext",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()
			id := "random_id"
			msg := "Super puper hello world"
			sut.inboxItemsRepo.Create(models.InboxItem{
				Id:      id,
				Message: msg,
			})

			task, err := sut.ConvertToTask(id, 0, models.EnergyLow, nil)

			if err != nil {
				t.Errorf("Got unexpected error: %v", err)
				return
			}

			if task == nil {
				t.Errorf("Task is nil. Expecting task with message: %v", msg)
				return
			}

			if task.Message != msg {
				t.Errorf("Invalid task message. Expecting: %v, got: %v", msg, task.Message)
				return
			}

			if task.Parent.Type != models.BoxParentType {
				t.Errorf(
					"Invalid parent type value. Expected: %v, got: %v",
					models.BoxParentType,
					task.Parent.Type,
				)
				return
			}

			if task.Parent.Id != models.BoxTypeNext.String() {
				t.Errorf(
					"Invalid parent id value: Expected: %v, got: %v",
					models.BoxTypeNext.String(),
					task.Parent.Id,
				)
				return
			}
		},
	)
}

func TestAddTask(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN empty tasks repo WHEN adding new task THEN task is added",
		func(t *testing.T) {
			utils.Delete(storageFp)
			msg := "hello, world"
			sut := createInteractor()

			task, err := sut.AddTask(msg, 0, models.EnergyLow, nil)

			if err != nil {
				panic(err)
			}

			if task.Message != msg {
				t.Errorf("Invalid task created. Expected: %v, got: %v", msg, task)
				return
			}
		},
	)
	t.Run(
		"GIVEN empty tasks repo with none empty inbox "+
			"WHEN adding new task "+
			"THEN task is added without affect on inbox",
		func(t *testing.T) {
			utils.Delete(storageFp)
			msg := "hello, world"
			sut := createInteractor()
			sut.inboxItemsRepo.Create(
				models.InboxItem{Id: "random id", Message: "other message"},
			)

			task, err := sut.AddTask(msg, 0, models.EnergyLow, nil)

			if err != nil {
				panic(err)
			}

			items, _ := sut.inboxItemsRepo.List()

			if len(items) != 1 {
				t.Errorf(
					"AddTask method affected inbox repo, expected len to be: 1, bug got: %v",
					len(items),
				)
				return
			}

			if task.Message != msg {
				t.Errorf("Invalid task created. Expected: %v, got: %v", msg, task)
				return
			}
		},
	)
}

func TestDeleteTask(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN no tasks in repo WHEN try to delete task THEN no task deleted",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			task, err := sut.DeleteTask("random_id")

			if task != nil {
				t.Errorf("Task was deleted, but shouldn't: %v", task)
			}

			if err == nil {
				t.Errorf("Error is nil, but shouldn't: %v", err)
			}

			if !errors.Is(err, repository.ErrNotFound) {
				t.Errorf(
					"Got an error of different type. Expected: %v, got: %v",
					repository.ErrNotFound,
					err,
				)
			}
		},
	)

	t.Run(
		"GIVEN tasks in repo "+
			"WHEN try to delete task that do NOT in the repo "+
			"THEN no task is deleted",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()
			sut.tasksRepo.Create(models.Task{Id: "Hello"})

			task, err := sut.DeleteTask("random_id")

			if task != nil {
				t.Errorf("Task was deleted, but shouldn't: %v", task)
			}

			if err == nil {
				t.Errorf("Error is nil, but shouldn't: %v", err)
			}

			if !errors.Is(err, repository.ErrNotFound) {
				t.Errorf(
					"Got an error of different type. Expected: %v, got: %v",
					repository.ErrNotFound,
					err,
				)
			}
		},
	)

	t.Run(
		"GIVEN tasks in repo "+
			"WHEN try to delete task that is in repo "+
			"THEN this task is deleted",
		func(t *testing.T) {
			utils.Delete(storageFp)
			id := "random_id"
			exp := models.Task{Id: id}
			sut := createInteractor()
			sut.tasksRepo.Create(exp)

			task, err := sut.DeleteTask(id)

			if err != nil {
				panic(err)
			}

			if *task != exp {
				t.Errorf(
					"Invalid task deleted. Expected: %v, got: %v",
					exp,
					*task,
				)
			}
		},
	)
}

func TestToggleFavourite(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN no tasks in repo WHEN toggling favourite THEN nothing changed",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			task, err := sut.ToggleFavourite("random_id")

			if task != nil {
				t.Errorf("Task was toggled favourite, but shouldn't: %v", task)
			}

			if err == nil {
				t.Errorf("Error is nil, but shouldn't: %v", err)
			}

			if !errors.Is(err, repository.ErrNotFound) {
				t.Errorf(
					"Got an error of different type. Expected: %v, got: %v",
					repository.ErrNotFound,
					err,
				)
			}
		},
	)

	t.Run(
		"GIVEN favourite task in repo WHEN toggling favourite THEN favourite set to false",
		func(t *testing.T) {
			utils.Delete(storageFp)
			id := "random"
			initialTask := models.Task{Id: id, Favourite: true}
			sut := createInteractor()
			sut.tasksRepo.Create(initialTask)

			task, err := sut.ToggleFavourite(id)

			if err != nil {
				panic(err)
			}

			if task.Id != initialTask.Id {
				t.Errorf(
					"Invalid task was toggled favourite. Expected: %v, got: %v",
					initialTask.Id,
					task.Id,
				)
			}

			if task.Favourite == initialTask.Favourite {
				t.Errorf(
					"Favourite wasn't toggle. Expected: %v, got: %v",
					initialTask,
					task,
				)
			}
		},
	)

	t.Run(
		"GIVEN unfavourite task in repo WHEN toggling favourite THEN favourite set to true",
		func(t *testing.T) {
			utils.Delete(storageFp)
			id := "random"
			initialTask := models.Task{Id: id, Favourite: false}
			sut := createInteractor()
			sut.tasksRepo.Create(initialTask)

			task, err := sut.ToggleFavourite(id)

			if err != nil {
				panic(err)
			}

			if task.Id != initialTask.Id {
				t.Errorf(
					"Invalid task was toggled favourite. Expected: %v, got: %v",
					initialTask.Id,
					task.Id,
				)
			}

			if task.Favourite == initialTask.Favourite {
				t.Errorf(
					"Favourite wasn't toggle. Expected: %v, got: %v",
					initialTask,
					task,
				)
			}
		},
	)
}

func TestSetStatus(t *testing.T) {
	defer utils.Delete(storageFp)

	t.Run(
		"GIVEN no tasks in repo WHEN setting status THEN nothing changed",
		func(t *testing.T) {
			utils.Delete(storageFp)
			sut := createInteractor()

			task, err := sut.SetStatus("random_id", models.TaskStatusInProgress)

			if task != nil {
				t.Errorf(
					"Task was created duting setting the status, but shouldn't: %v",
					task,
				)
			}

			if err == nil {
				t.Errorf("Error is nil, but shouldn't: %v", err)
			}

			if !errors.Is(err, repository.ErrNotFound) {
				t.Errorf(
					"Got an error of different type. Expected: %v, got: %v",
					repository.ErrNotFound,
					err,
				)
			}
		},
	)

	t.Run(
		"GIVEN task in pending status WHEN setting the same status THEN nothing is changed",
		func(t *testing.T) {
			utils.Delete(storageFp)
			id := "random"
			initialTask := models.Task{Id: id, Status: models.TaskStatusPending}
			sut := createInteractor()
			sut.tasksRepo.Create(initialTask)

			task, err := sut.SetStatus(id, models.TaskStatusPending)

			if err != nil {
				panic(err)
			}

			if task.Id != initialTask.Id {
				t.Errorf(
					"Invalid task changed. Expected: %v, got: %v",
					initialTask.Id,
					task.Id,
				)
			}

			if task.Status != initialTask.Status {
				t.Errorf(
					"Status changed. Expected: %v, got: %v",
					initialTask,
					task,
				)
			}
		},
	)

	t.Run(
		"GIVEN task in pending status WHEN setting other status THEN status is changed",
		func(t *testing.T) {
			utils.Delete(storageFp)
			id := "random"
			initialTask := models.Task{Id: id, Status: models.TaskStatusPending}
			sut := createInteractor()
			sut.tasksRepo.Create(initialTask)

			task, err := sut.SetStatus(id, models.TaskStatusInProgress)

			if err != nil {
				panic(err)
			}

			if task.Id != initialTask.Id {
				t.Errorf(
					"Invalid task was changed. Expected: %v, got: %v",
					initialTask.Id,
					task.Id,
				)
			}

			if task.Status == initialTask.Status {
				t.Errorf(
					"Status wasn't changed. Expected: %v, got: %v",
					initialTask,
					task,
				)
			}
		},
	)

	t.Run(
		"GIVEN task in pending status "+
			"WHEN setting done status "+
			"THEN status is changed and task is moved to the done repo",
		func(t *testing.T) {
			utils.Delete(storageFp)
			id := "random"
			initialTask := models.Task{Id: id, Status: models.TaskStatusPending}
			sut := createInteractor()
			sut.tasksRepo.Create(initialTask)

			task, err := sut.SetStatus(id, models.TaskStatusDone)

			if err != nil {
				panic(err)
			}

			if task.Id != initialTask.Id {
				t.Errorf(
					"Invalid task was changed. Expected: %v, got: %v",
					initialTask.Id,
					task.Id,
				)
			}

			if task.Status == initialTask.Status {
				t.Errorf(
					"Status wasn't changed. Expected: %v, got: %v",
					initialTask,
					task,
				)
			}

			tasks, err := sut.tasksRepo.List()

			if err != nil {
				panic(err)
			}

			if len(tasks) != 0 {
				t.Errorf(
					"Invalid amount of tasks in the original repo. Expected: 0, got: %v",
					len(tasks),
				)
			}

			tasks, err = sut.doneTasksRepo.List()

			if err != nil {
				panic(err)
			}

			if len(tasks) != 1 {
				t.Errorf(
					"Invalid amount of tasks in the done repo. Expected: 1, got: %v",
					len(tasks),
				)
			}
		},
	)
}

func createInteractor() *clarifyImpl {
	return NewClarifyInteractor(
		createTasksRepo(),
		createDoneTasksRepo(),
		createInboxItemsRepo(),
	)
}

func createTasksRepo() repository.Repo[models.Task, string] {
	repo, err := repository.NewFPRepo(tasksFp, func(task models.Task) string { return task.Id })

	if err != nil {
		panic(err)
	}

	return repo
}

func createDoneTasksRepo() repository.Repo[models.Task, string] {
	repo, err := repository.NewFPRepo(doneTasksFp, func(task models.Task) string { return task.Id })

	if err != nil {
		panic(err)
	}

	return repo
}

func createInboxItemsRepo() repository.Repo[models.InboxItem, string] {
	repo, err := repository.NewFPRepo(inboxFp, func(item models.InboxItem) string { return item.Id })

	if err != nil {
		panic(err)
	}

	return repo
}
