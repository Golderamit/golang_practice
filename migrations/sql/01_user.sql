-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE  users
(
    id         SERIAL PRIMARY KEY,
    first_name TEXT        NOT NULL,
    last_name  TEXT        NOT NULL,
    username   TEXT unique NOT NULL,
    email      TEXT unique NOT NULL,
    password   TEXT       NOT NULL,
    is_active  boolean     default true,
    is_admin   boolean     default false,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()

);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE  users;
