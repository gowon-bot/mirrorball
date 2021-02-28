-- +migrate Up
CREATE TABLE artists (id SERIAL PRIMARY KEY, name text);
-- +migrate Down
DROP TABLE artists;
