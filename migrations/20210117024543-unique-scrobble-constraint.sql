-- +migrate Up
ALTER TABLE scrobbles
ADD CONSTRAINT uniqueness UNIQUE (timestamp, track_id);
-- +migrate Down
ALTER TABLE scrobbles DROP CONSTRAINT uniqueness;