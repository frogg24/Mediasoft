CREATE TABLE employees (
    id BIGSERIAL CONSTRAINT employees_pkey PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    position TEXT NOT NULL
);