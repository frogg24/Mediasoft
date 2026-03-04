package database

import (
	"context"
	"log"
	"mediasoft/lesson9/internal/model"
)

func (db *DB) CreateGroup(ctx context.Context, group model.Group) error {
	const q = `
		insert into groups (title, parentgroup) values ($1, $2);
	`

	_, err := db.ExecContext(ctx, q, group.Title, group.ParentGroup)
	return err
}

func (db *DB) ReadGroup(ctx context.Context, id int64) (model.Group, error) {
	const q = `
		select id, title, parentgroup from groups where id = $1;
	`

	group := model.Group{}
	return group, db.QueryRowContext(ctx, q, id).Scan(
		&group.ID,
		&group.Title,
		&group.ParentGroup,
	)
}

func (db *DB) UpdateGroup(ctx context.Context, group model.Group) error {
	const q = `
		update groups set title = $1, parentgroup = $2, where id = $3;
	`

	_, err := db.ExecContext(ctx, q, group.Title, group.ParentGroup, group.ID)
	return err
}

func (db *DB) DeleteGroup(ctx context.Context, id int64) error {
	const q = `
		delete from groups where id = $1;
	`
	_, err := db.ExecContext(ctx, q, id)
	return err
}

func (db *DB) ListPersonLocal(ctx context.Context, id int64) ([]model.Person, error) {
	const q = `
		select id, name, lastname, birthdate, groupid from persons where groupid = $1;
	`
	rows, err := db.QueryContext(ctx, q, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	persons := make([]model.Person, 0, 16)
	for rows.Next() {
		person := model.Person{}
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
	return persons, nil
}

func (db *DB) ListPersonAll(ctx context.Context, groupID int64) ([]model.Person, error) {
	const q = `
		WITH RECURSIVE subgroups AS (
			SELECT id
			FROM groups
			WHERE id = $1

			UNION ALL

			SELECT g.id
			FROM groups g
			INNER JOIN subgroups sg ON g.parent_id = sg.id
		)
		SELECT p.id, p.name, p.lastname, p.birthdate, p.groupid
		FROM persons p
		WHERE p.groupid IN (SELECT id FROM subgroups);
	`

	rows, err := db.QueryContext(ctx, q, groupID)
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

func (db *DB) CountGroupLocal(ctx context.Context, groupid int64) (int64, error) {
	const q = `
		SELECT COUNT(*) 
		FROM persons 
		WHERE group_id = $1;
	`

	var count int64
	err := db.QueryRowContext(ctx, q, groupid).Scan(&count)
	return count, err
}

func (db *DB) CountGroupAll(ctx context.Context, groupid int64) (int64, error) {
	const q = `
		WITH RECURSIVE subgroups AS (
			SELECT id
			FROM groups
			WHERE id = $1

			UNION ALL

			SELECT g.id
			FROM groups g
			INNER JOIN subgroups sg ON g.parent_id = sg.id
		)
		SELECT COUNT(*)
		FROM persons
		WHERE group_id IN (SELECT id FROM subgroups);
	`

	var count int64
	err := db.QueryRowContext(ctx, q, groupid).Scan(&count)
	return count, err
}

func (db *DB) GetAllGroups(ctx context.Context) ([]model.Group, error) {
	log.Println("GetAllGroups in database")

	const q = `
		SELECT * FROM groups;
	`
	rows, err := db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := make([]model.Group, 0, 16)
	for rows.Next() {
		var group model.Group
		if err := rows.Scan(
			&group.ID,
			&group.Title,
			&group.ParentGroup,
		); err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
