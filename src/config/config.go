package config

type Config struct {
	Storage string `json:"storage"`
}

const (
	inbox     = "/inbox.json"
	tasks     = "/tasks.json"
	doneTasks = "/done_tasks.json"
)

func (this *Config) GetInboxPath() string {
	return this.Storage + inbox
}

func (this *Config) GetTasksPath() string {
	return this.Storage + tasks
}

func (this *Config) GetDoneTasksPath() string {
	return this.Storage + doneTasks
}
