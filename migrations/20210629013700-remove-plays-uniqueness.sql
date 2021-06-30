-- +migrate Up
ALTER TABLE plays DROP CONSTRAINT uniqueness;

-- +migrate Down
ALTER TABLE plays ADD CONSTRAINT uniqueness UNIQUE (scrobbled_at, track_id, user_id);
