package models

type Project struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Tasks []string `json:"tasks"`
}
