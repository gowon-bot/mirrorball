-- +migrate Up
CREATE INDEX tc_track_id_idx ON track_counts ((track_id));
CREATE INDEX al_artist_id_idx ON albums ((artist_id));
CREATE INDEX tr_artist_id_idx ON tracks ((artist_id));
-- +migrate Down
DROP INDEX tc_track_id_idx;
DROP INDEX al_artist_id_idx;
DROP INDEX tr_artist_id_idx;