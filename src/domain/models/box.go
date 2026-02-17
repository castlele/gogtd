package models

type BoxType string

const (
	BoxTypeNext         BoxType = "BoxTypeNext"
	BoxTypeWaiting      BoxType = "BoxTypeWaiting"
	BoxTypeSomedayMaybe BoxType = "BoxTypeSomedayMaybe"
)

func (bt BoxType) String() string {
	return string(bt)
}
