package models

type Task struct {
	Id        string     `json:"id"`
	Message   string     `json:"message"`
	Time      int64      `json:"time"`
	Parent    TaskParent `json:"parent"`
	Energy    Energy     `json:"energy"`
	Status    TaskStatus `json:"status"`
	Favourite bool       `json:"favourite"`
}

type TaskParent struct {
	Id   string         `json:"id"`
	Type TaskParentType `json:"type"`
}

func NewNextTaskParent() TaskParent {
	return TaskParent{
		Id:   BoxTypeNext.String(),
		Type: BoxParentType,
	}
}

type TaskParentType int

const (
	BoxParentType TaskParentType = iota
	ProjectParentType
	StepParentType
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "TaskStatusPending"
	TaskStatusInProgress TaskStatus = "TaskStatusInProgress"
	TaskStatusDone       TaskStatus = "TaskStatusDone"
)
