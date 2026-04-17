INSERT INTO groups (id, title, parentgroup) VALUES
    (1, 'Company', NULL),
    (2, 'Backend', 1),
    (3, 'Frontend', 1),
    (4, 'QA', 1),
    (5, 'Go Developers', 2);

SELECT setval(pg_get_serial_sequence('groups', 'id'), (SELECT MAX(id) FROM groups));

INSERT INTO persons (id, name, lastname, birthdate, groupid) VALUES
    (1, 'Ivan', 'Ivanov', '1998-01-10', 1),
    (2, 'Petr', 'Petrov', '1999-02-15', 2),
    (3, 'Anna', 'Annova', '2000-03-20', 2),
    (4, 'Maria', 'Mariova', '2001-04-25', 3),
    (5, 'Oleg', 'Olegov', '1997-05-30', 5);

SELECT setval(pg_get_serial_sequence('persons', 'id'), (SELECT MAX(id) FROM persons));