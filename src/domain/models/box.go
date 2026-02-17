package models

type BoxType int

const (
	Next BoxType = iota
	Waiting
	SomedayMaybe
)

type Box struct {
	Id   string  `json:"id"`
	Type BoxType `json:"type"`
}
