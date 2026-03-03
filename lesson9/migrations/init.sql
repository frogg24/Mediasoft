CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    parentgroup INT REFERENCES groups(id)
);

CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    lastname TEXT NOT NULL,
    birthdate DATE NOT NULL,
    groupid INT REFERENCES groups(id)
);