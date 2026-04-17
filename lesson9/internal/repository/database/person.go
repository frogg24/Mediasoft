package database

import (
	"context"
	"mediasoft/lesson9/internal/model"
	"mediasoft/lesson9/internal/repository"
)

func (db *DB) CreatePerson(ctx context.Context, person model.Person) error {
	groupExists, err := db.groupExists(ctx, person.GroupID)
	if err != nil {
		return err
	}
	if !groupExists {
		return repository.ErrGroupNotFound
	}

	const q = `
		insert into persons (name, lastname, birthdate, groupid) values ($1, $2, $3, $4);
	`

	_, err = db.ExecContext(ctx, q, person.Name, person.Lastname, person.Birthdate, person.GroupID)
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

func (db *DB) GetAllPersons(ctx context.Context) ([]model.Person, error) {
	const q = `
		select id, name, lastname, birthdate, groupid from persons;
	`
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	persons := make([]model.Person, 0, 16)
	for rows.Next() {
		var person model.Person
		if err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Lastname,
			&person.Birthdate,
			&person.GroupID,
		); err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}

func (db *DB) UpdatePerson(ctx context.Context, person model.Person) error {
	groupExists, err := db.groupExists(ctx, person.GroupID)
	if err != nil {
		return err
	}
	if !groupExists {
		return repository.ErrGroupNotFound
	}

	const q = `
		update persons set name = $1, lastname = $2, birthdate = $3, groupid = $4 where id = $5;
	`

	result, err := db.ExecContext(ctx, q, person.Name, person.Lastname, person.Birthdate, person.GroupID, person.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository.ErrNotFound
	}

	return nil
}

func (db *DB) DeletePerson(ctx context.Context, id int64) error {
	const q = `
		delete from persons where id = $1;
	`
	result, err := db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repository.ErrNotFound
	}

	return nil
}
