package model

import "time"

type Person struct {
	ID        int64
	Name      string
	Lastname  string
	Birthdate time.Time
	GroupID   int64
}
