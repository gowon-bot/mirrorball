-- +migrate Up
CREATE INDEX play_track_id_idx ON plays ((track_id));

-- +migrate Down
DROP INDEX play_track_id_idx;