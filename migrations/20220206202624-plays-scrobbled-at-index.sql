-- +migrate Up
CREATE INDEX play_scrobbled_at_idx ON plays ((scrobbled_at));

-- +migrate Down
DROP INDEX play_scrobbled_at_idx;