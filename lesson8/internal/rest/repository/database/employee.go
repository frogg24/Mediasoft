package database

import (
	"context"
	"database/sql"
	"mediasoft/lesson8/internal/rest/model"
)

type DB struct {
	*sql.DB
}

func NewDatabase(db *sql.DB) *DB {
	return &DB{db}
}

func (db *DB) Create(ctx context.Context, employee model.Employee) error {
	const q = `
		insert into employees (name, surname, position) values ($1, $2, $3);
	`

	_, err := db.ExecContext(ctx, q, employee.Name, employee.Surname, employee.Position)
	return err
}

func (db *DB) Read(ctx context.Context, id int64) (model.Employee, error) {
	const q = `
		select id, name, surname, position from employees where id = $1;
	`

	employee := model.Employee{}
	return employee, db.QueryRowContext(ctx, q, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Surname,
		&employee.Position,
	)
}

func (db *DB) Update(ctx context.Context, employee model.Employee) error {
	const q = `
		update employees set name = $1, surname = $2, position = $3 where id = $4;
	`

	_, err := db.ExecContext(ctx, q, employee.Name, employee.Surname, employee.Position, employee.ID)
	return err
}

func (db *DB) Delete(ctx context.Context, id int64) error {
	const q = `
		delete from employees where id = $1;
	`
	_, err := db.ExecContext(ctx, q, id)
	return err
}

func (db *DB) List(ctx context.Context) ([]model.Employee, error) {
	const q = `
		select id, name, surname, position from employees order by id desc;
	`
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	employees := make([]model.Employee, 0, 16)
	for rows.Next() {
		employee := model.Employee{}
		if err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Surname,
			&employee.Position,
		); err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}
	return employees, nil
}
