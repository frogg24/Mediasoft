package repository

import "errors"

var (
	ErrNotFound         = errors.New("record not found")
	ErrGroupNotFound    = errors.New("group not found")
	ErrGroupHasChildren = errors.New("group has child groups")
	ErrGroupHasPersons  = errors.New("group has persons")
	ErrGroupCycle       = errors.New("group hierarchy cycle is not allowed")
)
