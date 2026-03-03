CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    birth_year INT NOT NULL,
    group_id INT REFERENCES groups(id)
);