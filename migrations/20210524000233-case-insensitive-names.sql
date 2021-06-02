-- +migrate Up
CREATE EXTENSION citext;

ALTER TABLE artists
ALTER COLUMN name TYPE citext;

ALTER TABLE albums
ALTER COLUMN name TYPE citext;

ALTER TABLE tracks
ALTER COLUMN name TYPE citext;

-- +migrate Down
ALTER TABLE artists
ALTER COLUMN name TYPE text;

ALTER TABLE albums
ALTER COLUMN name TYPE text;

ALTER TABLE tracks
ALTER COLUMN name TYPE text;

DROP EXTENSION citext;

