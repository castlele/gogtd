package config

type Config struct {
	Storage string `json:"storage"`
}

const (
	inbox = "/inbox.json"
)

func (this *Config) GetInboxPath() string {
	return this.Storage + inbox
}
