-- +migrate Up
CREATE TABLE tracks (
  id SERIAL PRIMARY KEY,
  name text,
  artist_id integer REFERENCES artists(id),
  album_id integer REFERENCES albums(id)
);
-- +migrate Down
DROP TABLE tracks;
