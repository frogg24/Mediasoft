package model

type Group struct {
	ID          int64
	Title       string
	ParentGroup *int64
}
