package repository

import (
	"context"
	"mediasoft/lesson8/internal/rest/model"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee model.Employee) error
	Read(ctx context.Context, id int64) (model.Employee, error)
	Update(ctx context.Context, employee model.Employee) error
	Delete(ctx context.Context, id int64) error

	List(ctx context.Context) ([]model.Employee, error)
}
