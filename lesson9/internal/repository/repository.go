package repository

import (
	"context"
	"mediasoft/lesson9/internal/model"
)

type PersonRepository interface {
	CreatePerson(ctx context.Context, person model.Person) error
	ReadPerson(ctx context.Context, personid int64) (model.Person, error)
	UpdatePerson(ctx context.Context, person model.Person) error
	DeletePerson(ctx context.Context, personid int64) error
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, group model.Group) error
	ReadGroup(ctx context.Context, groupid int64) (model.Group, error)
	UpdateGroup(ctx context.Context, group model.Group) error
	DeleteGroup(ctx context.Context, groupid int64) error
	GetAllGroups(ctx context.Context) ([]model.Group, error)

	ListPersonLocal(ctx context.Context, groupid int64) ([]model.Person, error)
	ListPersonAll(ctx context.Context, groupid int64) ([]model.Person, error)

	CountGroupLocal(ctx context.Context, groupid int64) (int64, error)
	CountGroupAll(ctx context.Context, groupid int64) (int64, error)
}
