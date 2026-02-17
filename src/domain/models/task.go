package models

type Task struct {
	Id      string     `json:"id"`
	Message string     `json:"message"`
	Time    int64      `json:"time"`
	Parent  TaskParent `json:"parent"`
	Energy  Energy     `json:"energy"`
}

type TaskParent struct {
	Id   string         `json:"id"`
	Type TaskParentType `json:"type"`
}

type TaskParentType int

const (
	BoxParentType TaskParentType = iota
	ProjectParentType
	StepParentType
)
