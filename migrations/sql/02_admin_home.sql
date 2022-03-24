-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE admin_home
(
    id                            SERIAL  PRIMARY KEY,
    title                         TEXT   NOT NULL DEFAULT '',
    venue                         TEXT   NOT NULL DEFAULT '',
    address                       TEXT   NOT NULL DEFAULT '',
    country                       TEXT   NOT NULL DEFAULT '',
    email                         TEXT   NOT NULL DEFAULT '',
    phone_number                  TEXT   NOT NULL DEFAULT '',
    short_description             TEXT   NOT NULL DEFAULT '',
    description                   TEXT   NOT NULL DEFAULT '',
    image                         TEXT   NOT NULL DEFAULT '',
    from_date                     TIMESTAMPTZ  NOT NULL,
    to_date                       TIMESTAMPTZ  NOT NULL,
    status                        TEXT   NOT NULL DEFAULT '',
    created TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated TIMESTAMPTZ NOT NULL DEFAULT now()
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE admin_home; -- also drops the trigger