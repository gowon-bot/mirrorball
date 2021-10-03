-- +migrate Up
CREATE INDEX rymsl_rate_your_music_id_idx ON rate_your_music_albums ((rate_your_music_id));
-- +migrate Down
DROP INDEX rymsl_rate_your_music_id_idx;