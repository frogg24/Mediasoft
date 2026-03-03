package database

import (
	"context"
	"mediasoft/lesson9/internal/model"
)

func (db *DB) CreatePerson(ctx context.Context, person model.Person) error {
	const q = `
		insert into persons (name, lastname, birthdate, groupid) values ($1, $2, $3, $4);
	`

	_, err := db.ExecContext(ctx, q, person.Name, person.Lastname, person.Birthdate, person.GroupID)
	return err
}

func (db *DB) ReadPerson(ctx context.Context, id int64) (model.Person, error) {
	const q = `
		select id, name, lastname, birthdate, groupid from persons where id = $1;
	`

	person := model.Person{}
	return person, db.QueryRowContext(ctx, q, id).Scan(
		&person.ID,
		&person.Name,
		&person.Lastname,
		&person.Birthdate,
		&person.GroupID,
	)
}

func (db *DB) UpdatePerson(ctx context.Context, person model.Person) error {
	const q = `
		update persons set name = $1, lastname = $2, birthdate = $3, groupid = $4 where id = $5;
	`

	_, err := db.ExecContext(ctx, q, person.Name, person.Lastname, person.Birthdate, person.GroupID, person.ID)
	return err
}

func (db *DB) DeletePerson(ctx context.Context, id int64) error {
	const q = `
		delete from persons where id = $1;
	`
	_, err := db.ExecContext(ctx, q, id)
	return err
}
