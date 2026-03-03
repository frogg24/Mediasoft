CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    parent_id INT REFERENCES groups(id)
);