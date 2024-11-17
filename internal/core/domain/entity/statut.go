package entity

type Status int

const (
	_ Status = iota
	StatusQueued
	StatusSent
	StatusHit
	StatusFailed
)

func (s Status) Int() int {
	return int(s)
}
