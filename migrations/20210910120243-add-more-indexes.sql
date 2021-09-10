-- +migrate Up
CREATE INDEX ac_artist_id_idx ON artist_counts ((artist_id));
CREATE INDEX gm_user_id ON guild_members ((user_id));
-- +migrate Down
DROP INDEX ac_artist_id_idx;
DROP INDEX gm_user_id;