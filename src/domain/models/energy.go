package models

type Energy int

const (
	EnergyLow  Energy = iota
	EnergyMid  Energy = iota
	EnergyHigh Energy = iota
)
