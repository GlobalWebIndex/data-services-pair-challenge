-- +migrate Up

CREATE TABLE audiences (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    expression JSONB NOT NULL
);
