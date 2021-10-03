-- +migrate Up
CREATE INDEX rymsl_rate_your_music_id_idx ON rate_your_music_albums ((rate_your_music_id));
CREATE INDEX rymsll_album_id_idx ON rate_your_music_album_albums ((album_id));
-- +migrate Down
DROP INDEX rymsl_rate_your_music_id_idx;
DROP INDEX rymsll_album_id_idx;